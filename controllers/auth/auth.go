package authCtl

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Olapat/go-architecture/controllers/response"
	"github.com/Olapat/go-architecture/controllers/validate"
	authEntity "github.com/Olapat/go-architecture/model/entity/auth"
	tokenResetPasswordEntity "github.com/Olapat/go-architecture/model/entity/token_reset_password"
	userEntity "github.com/Olapat/go-architecture/model/entity/user"
	tokenResetPasswordRepo "github.com/Olapat/go-architecture/model/repository/token_reset_password"

	"strings"
	"time"

	notiUtils "github.com/Olapat/go-architecture/utils/notification"
	strUtils "github.com/Olapat/go-architecture/utils/string"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	userRepo "github.com/Olapat/go-architecture/model/repository/user"
)

func createToken(user authEntity.TokenLogin, loginWith string) (string, error) {
	var err error
	secret := os.Getenv("SECRET")
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = user.ID
	atClaims["username"] = user.Username
	atClaims["role"] = user.Role
	atClaims["company_id"] = user.CompanyID
	if loginWith == "mobile" {
		atClaims["exp"] = time.Now().Add(time.Minute * ((60 * 24) * 30)).Unix()
	} else {
		atClaims["exp"] = time.Now().Add(time.Minute * (60 * 24)).Unix()
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func parseToken(tokenString string) (*authEntity.TokenParser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authEntity.TokenParser{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if token == nil && err != nil {
		fmt.Println("err")
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*authEntity.TokenParser); ok {
		// fmt.Println(token.Valid)
		// fmt.Printf("%v %v %v", claims.LoginName, claims.Email, claims.StandardClaims.ExpiresAt.Unix())
		if claims != nil {
			// fmt.Println(claims.StandardClaims.ExpiresAt)
			if claims.StandardClaims.ExpiresAt != nil && claims.StandardClaims.ExpiresAt.Before(time.Now()) {
				// return nil, &jwt.TokenExpiredError{}
				return nil, errors.New("Session is expired")
			}
		}
		claims.Token = tokenString
		return claims, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func saveFCMToken(c *fiber.Ctx, id string, token string) {
	rawBody := bson.M{
		"$set": bson.M{
			"fcm_token": token,
		},
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	_, e := userRepo.UpdateByID(c, objectID, rawBody, false)
	log.Println("error", e)
}

func UserAuth(c *fiber.Ctx, userInfo userEntity.User, loginWith *string) error {
	var _loginWith = "web"
	if loginWith != nil {
		_loginWith = *loginWith
	}
	token, errToken := createToken(authEntity.TokenLogin{
		ID:        userInfo.ID.Hex(),
		Username:  userInfo.Username,
		Role:      userInfo.Role,
		CompanyID: userInfo.CompanyID.Hex(),
		LoginWith: _loginWith,
	}, _loginWith)

	if errToken != nil {
		return response.ResponseError(c, fiber.StatusInternalServerError, "", nil)
	}

	// roleMenu, errRole := configRoleMenuRepo.FindOne(c, nil)
	// if errRole != nil {
	// 	return response.ResponseError(c, fiber.StatusInternalServerError, "", errRole.Error())
	// }

	// menu := make([]configRoleMenuEntity.Menu, 0)

	// company, _ := companyRepository.FindOneByID(c, userInfo.CompanyID.Hex())

	return response.ResponseOK(c, fiber.StatusOK, authEntity.ResponseSignIn{
		ID:        userInfo.ID.Hex(),
		Username:  userInfo.Username,
		Role:      userInfo.Role,
		Token:     token,
		CompanyID: userInfo.CompanyID.Hex(),
		// FullName:      "userInfo.FullName",
		// Menu:          menu,
		CompanyStatus: "ACTIVE",
		Status:        userInfo.Status,
	}, "")
}

func tokenToUser(c *fiber.Ctx) (userEntity.User, error, string) {
	tokenString := c.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	if tokenString == "" {
		return userEntity.User{}, fiber.ErrUnauthorized, ""
	}

	userFromToken, errParseToken := parseToken(tokenString)

	if errParseToken != nil {
		return userEntity.User{}, errParseToken, ""
	}
	if userFromToken == nil {
		return userEntity.User{}, fiber.ErrUnauthorized, ""
	}

	userObjId, _ := primitive.ObjectIDFromHex(userFromToken.ID)
	filter := bson.M{
		"deleted_at": nil,
		"_id":        userObjId,
	}
	userInfo, errFind := userRepo.FindOne(c, filter)
	if errFind != nil {
		return userEntity.User{}, fiber.ErrUnauthorized, ""
	}

	return userInfo, nil, userFromToken.LoginWith
}

func Authorization(c *fiber.Ctx) error {
	userInfo, errTokenToUser, LoginWith := tokenToUser(c)
	if errTokenToUser != nil {
		return response.ResponseError(c, fiber.StatusUnauthorized, errTokenToUser.Error(), nil)
	}

	if userInfo.ID.Hex() == "" {
		return response.ResponseError(c, fiber.StatusUnauthorized, "", nil)
	}

	c.Locals("user.username", userInfo.Username)
	c.Locals("user.id", userInfo.ID.Hex())
	c.Locals("user.company_id", userInfo.CompanyID.Hex())
	c.Locals("user.role", userInfo.Role)
	c.Locals("user.login_with", LoginWith)

	return c.Next()
}

func GetUser(c *fiber.Ctx) (primitive.ObjectID, primitive.ObjectID, string) {
	userID := fmt.Sprintf("%s", c.Locals("user.id"))
	userObjectID, _ := primitive.ObjectIDFromHex(userID)

	companyID := fmt.Sprintf("%s", c.Locals("user.company_id"))
	companyObjectID, _ := primitive.ObjectIDFromHex(companyID)

	role := fmt.Sprintf("%s", c.Locals("user.role"))

	return userObjectID, companyObjectID, role
}

// SignIn godoc
// @Tags     Auth
// @Summary  sign in [mobile, web]
// @Produce  json
// @Param    Body body     authEntity.BodySignIn true "BodySignIn"
// @Success  200  {object} entity.Response200{data=authEntity.ResponseSignIn}
// @Failure  400  {object} entity.Response400
// @Router   /auth/sign_in [post]
func SignIn(c *fiber.Ctx) error {
	body := new(authEntity.BodySignIn)
	// Parse body into struct
	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}

	errorValid := validate.ValidateStruct(*body)
	if errorValid != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", errorValid)
	}

	role := make([]string, 0)
	if body.Role == "EMPLOYEE" {
		role = append(role, "EMPLOYEE", "HR_ADMIN")
	} else {
		role = append(role, body.Role)
	}

	filter := bson.M{
		"role": bson.M{
			"$in": role,
		},
		"username":   body.Username,
		"deleted_at": nil,
	}

	userInfo, errFind := userRepo.FindOne(c, filter)

	if errFind != nil {
		if errFind == mongo.ErrNoDocuments {
			return response.ResponseError(c, fiber.StatusUnauthorized, "", nil)
		}
		return response.ResponseError(c, fiber.StatusInternalServerError, "", nil)
	}

	match := checkPasswordHash(body.Password, userInfo.Password)
	if !match {
		return response.ResponseError(c, fiber.StatusUnauthorized, "", nil)
	}

	if body.FcmToken != "" {
		go saveFCMToken(c, userInfo.ID.Hex(), body.FcmToken)
	}

	return UserAuth(c, userInfo, &body.LoginWith)
}

// RequestResetPassword godoc
// @Tags     Auth
// @Summary  request reset password [mobile, web]
// @Produce  json
// @Param    Body body     authEntity.BodyRequestResetPassword true "BodyRequestResetPassword"
// @Success  200  {object} entity.Response200{data=string}
// @Failure  400  {object} entity.Response400
// @Router   /auth/reset_password/request [post]
func RequestResetPassword(c *fiber.Ctx) error {
	body := new(authEntity.BodyRequestResetPassword)
	// Parse body into struct
	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}

	errorValid := validate.ValidateStruct(*body)
	if errorValid != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", errorValid)
	}

	filterUser := bson.M{
		"email":      body.Email,
		"deleted_at": nil,
	}

	user, errU := userRepo.FindOne(c, filterUser)
	if errU != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "Invalid email", nil)
	}

	now := time.Now()
	token := strUtils.StringRandom(64)
	rawBody := tokenResetPasswordEntity.TokenResetPassword{
		Email:     body.Email,
		UserID:    user.ID,
		CreatedAt: now,
		ExpiresAt: now.Add(time.Minute * 5),
		Token:     token,
		IsConfirm: false,
	}
	_, errT := tokenResetPasswordRepo.Insert(c, rawBody)
	if errT != nil {
		return response.ResponseError(c, fiber.StatusInsufficientStorage, "", nil)
	}

	go func() {
		hostWeb := os.Getenv("HOST_WEB")

		params := bson.M{
			"url": hostWeb + "/forgot-password/reset-password/" + token,
		}
		smpt := notiUtils.SMTP{
			To:     []string{body.Email},
			Params: params,
		}
		smpt.Send("reset_password")
	}()

	return response.ResponseOK(c, fiber.StatusOK, token, "success")
}

// VerifyResetPassword godoc
// @Tags     Auth
// @Summary  Verify setup password [mobile, web]
// @Produce  json
// @Param       token  path     string true "token"
// @Success  200  {object} entity.Response200{data=string}
// @Failure  400  {object} entity.Response400
// @Router   /auth/reset_password/verify/{token} [get]
func VerifyResetPassword(c *fiber.Ctx) error {
	token := c.Params("token")
	now := time.Now()

	filter := bson.M{
		"token": token,
		"expires_at": bson.M{
			"$gt": now,
		},
	}

	tr, err := tokenResetPasswordRepo.FindOne(c, filter)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "not_found", "")
	} else if tr.ID.IsZero() {
		return response.ResponseError(c, fiber.StatusBadRequest, "not_found2", "")
	}

	defer func() {
		filterU := bson.M{"email": tr.Email, "is_confirm": false}
		rawBodyCT := bson.M{
			"$set": bson.M{
				"is_confirm": true,
			},
		}
		tokenResetPasswordRepo.Update(c, filterU, rawBodyCT, false)
	}()

	return response.ResponseOK(c, fiber.StatusOK, tr.Email, "")
}

// ResetPassword godoc
// @Tags     Auth
// @Summary  reset password [mobile, web]
// @Produce  json
// @Param    Body body     authEntity.BodyResetPassword true "BodyResetPassword"
// @Success  200  {object} entity.Response200{data=authEntity.ResponseSignIn}
// @Failure  400  {object} entity.Response400
// @Router   /auth/reset_password/save [patch]
func ResetPassword(c *fiber.Ctx) error {
	body := new(authEntity.BodyResetPassword)
	// Parse body into struct
	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}

	errorValid := validate.ValidateStruct(*body)
	if errorValid != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", errorValid)
	}

	filterToken := bson.M{
		"is_confirm": true,
		"token":      body.Token,
	}
	tr, err := tokenResetPasswordRepo.FindOne(c, filterToken)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "token_not_found", "")
	} else if tr.ID.IsZero() {
		return response.ResponseError(c, fiber.StatusBadRequest, "token_not_found2", "")
	}

	filter := bson.M{
		"_id":        tr.UserID,
		"deleted_at": nil,
	}

	userInfo, errFind := userRepo.FindOne(c, filter)

	if errFind != nil {
		if errFind == mongo.ErrNoDocuments {
			return response.ResponseError(c, fiber.StatusUnauthorized, "", nil)
		}
		return response.ResponseError(c, fiber.StatusInternalServerError, "", nil)
	}

	password, _ := userRepo.HashPassword(body.Password)

	bodyUpdateUser := bson.M{
		"$set": bson.M{
			"password": password,
			"status":   userRepo.Status.ACTIVE,
		},
	}
	_, errUp := userRepo.UpdateByID(c, userInfo.ID, bodyUpdateUser, false)
	if errUp != nil {
		return response.ResponseError(c, fiber.StatusInternalServerError, "", nil)
	}

	return UserAuth(c, userInfo, &body.LoginWith)
}
