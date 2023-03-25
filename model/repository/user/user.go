package userRepo

import (
	"log"

	"github.com/Olapat/go-architecture/model/entity"
	userEntity "github.com/Olapat/go-architecture/model/entity/user"
	"github.com/Olapat/go-architecture/model/repository"
	strUtils "github.com/Olapat/go-architecture/utils/string"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var collectionName = "user"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func Insert(c *fiber.Ctx, rawBody userEntity.User) (interface{}, error) {
	_password := ""
	if rawBody.Password == "" {
		_password = strUtils.StringRandom(8)
	} else {
		_password = rawBody.Password
	}
	password, _ := HashPassword(_password)
	rawBody.Password = password

	result, err := repository.SuperInsertOne(c, collectionName, rawBody)
	return result, err
}

func BuildRows(rows []userEntity.User) []userEntity.User {
	var record = make([]userEntity.User, 0)
	for _, v := range rows {
		record = append(record, userEntity.User{
			ID:            v.ID,
			CompanyID:     v.CompanyID,
			Role:          v.Role,
			Username:      v.Username,
			Email:         v.Email,
			Password:      v.Password,
			Status:        v.Status,
			CreatedAt:     v.CreatedAt,
			CreatedBy:     v.CreatedBy,
			UpdatedAt:     v.UpdatedAt,
			UpdatedBy:     v.UpdatedBy,
			FcmToken:      v.FcmToken,
			EmployeeID:    v.EmployeeID,
			UserProfileID: v.UserProfileID,
		})
	}

	return record
}

func Find(c *fiber.Ctx, filter primitive.M, pagination *entity.PaginationRequests, sort *primitive.M) ([]userEntity.User, error) {
	var opts = options.Find()

	if pagination != nil {
		var page int64 = 1
		var perPage int64 = 10
		if pagination.Page != 0 {
			page = pagination.Page
		}

		if pagination.PerPage != 0 {
			perPage = pagination.PerPage
		}

		var skip = (page - 1) * perPage

		opts = options.Find().SetSkip(int64(skip)).SetLimit(int64(perPage))
	}

	if sort == nil {
		opts.SetSort(bson.M{"_id": 1})
	} else {
		opts.SetSort(&sort)
	}

	var results []userEntity.User = make([]userEntity.User, 0)

	filter["deleted_by"] = nil

	cursor, err := repository.SuperFind(c, collectionName, filter, opts)
	if err != nil {
		log.Println(err)

		return results, err
	}

	cursor.All(c.Context(), &results)

	return results, err
}

func FindOne(c *fiber.Ctx, filter primitive.M) (userEntity.User, error) {
	var entity userEntity.User
	errFind := repository.SuperFindOne(c, collectionName, filter, &entity)
	return entity, errFind
}

func Update(c *fiber.Ctx, filter primitive.M, body primitive.M, upsert bool) (interface{}, error) {
	result, errFind := repository.SuperUpdate(c, collectionName, filter, body, upsert)
	return result, errFind
}

func SoftDelete(c *fiber.Ctx, filter primitive.M, by *primitive.ObjectID) (interface{}, error) {
	result, errFind := repository.SuperSoftDelete(c, collectionName, filter, by)
	return result, errFind
}

func Count(c *fiber.Ctx, filter primitive.M) (int64, error) {
	count, err := repository.SuperCount(c, collectionName, filter, nil)
	return count, err
}

func UpdateByID(c *fiber.Ctx, id primitive.ObjectID, body primitive.M, upsert bool) (interface{}, error) {
	filter := bson.M{
		"_id":        id,
		"deleted_at": nil,
	}
	result, errFind := repository.SuperUpdate(c, collectionName, filter, body, upsert)
	return result, errFind
}

var Status = userEntity.Status{
	ACTIVE:          "ACTIVE",
	TO_SET_PASSWORD: "TO_SET_PASSWORD",
	CANCEL:          "CANCEL",
}
