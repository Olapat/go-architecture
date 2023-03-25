package authEntity

import "github.com/dgrijalva/jwt-go/v4"

type TokenLogin struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CompanyID string `json:"company_id"`
	LoginWith string `json:"login_with"`
}

type TokenParser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Token     string `json:"token"`
	CompanyID string `json:"company_id"`
	LoginWith string `json:"login_with"`
	jwt.StandardClaims
}

type BodySignIn struct {
	Username  string `json:"username" bson:"username" validate:"required" extensions:"x-order=0"`
	Password  string `json:"password" bson:"password" validate:"required" extensions:"x-order=1"`
	LoginWith string `json:"loginWith" bson:"loginWith" extensions:"x-order=2"`
	FcmToken  string `json:"fcmToken" bson:"fcmToken" extensions:"x-order=3"`
	Role      string `json:"role" bson:"role" validate:"required" enums:"HR_MANAGER,EMPLOYEE"`
}

type ResponseSignIn struct {
	ID        string  `json:"id,omitempty" bson:"_id,omitempty" example:"623d38efc41b8a687bedfe25"`
	Username  string  `json:"username" bson:"username" example:"Username"`
	Role      string  `json:"role"  bson:"role" example:"EMPLOYEE"`
	Token     string  `json:"token"  bson:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDg4MDM5MzMsImlkIjoiNjIzZDM4ZWZjNDliOGE2ODdiZWRmZTI1Iiwicm9sZV9pZCI6IjEiLCJ1c2VybmFtZSI6ImFkbWluNCJ9.ZSHeB6BsWBwelIz5uT_4zMFODv8A8nf19pSIavAU6iI"`
	CompanyID string  `json:"company_id" bson:"company_id" example:"623d38efc41b8a687bedfe21"`
	FullName  *string `json:"full_name" bson:"full_name" example:"full name"`
	// Menu          []roleMenuEntity.Menu `json:"menu" bson:"menu"`
	CompanyStatus string `json:"company_status" bson:"company_status"`
	Status        string `json:"status" bson:"status" example:"TO_SET_PASSWORD"`
}

type BodySetupPassword struct {
	Password  string `json:"password" bson:"password" validate:"required"`
	LoginWith string `json:"login_with"`
}

type BodyRequestResetPassword struct {
	Email string `json:"email"`
}

type BodyResetPassword struct {
	Password  string `json:"password" bson:"password" validate:"required"`
	Token     string `json:"token" bson:"token" validate:"required"`
	LoginWith string `json:"login_with"`
}
