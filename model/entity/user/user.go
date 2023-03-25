package userEntity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	CompanyID     primitive.ObjectID  `json:"company_id" bson:"company_id"`
	Role          string              `json:"role" bson:"role"`
	Username      string              `json:"username" bson:"username"`
	Email         string              `json:"email" bson:"email"`
	Password      string              `json:"password" bson:"password"`
	Status        string              `json:"status" bson:"status"`
	CreatedAt     time.Time           `json:"created_at" bson:"created_at"`
	CreatedBy     primitive.ObjectID  `json:"created_by" bson:"created_by"`
	UpdatedAt     *time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy     *primitive.ObjectID `json:"updated_by" bson:"updated_by"`
	FcmToken      *string             `json:"fcm_token" bson:"fcm_token"`
	EmployeeID    *primitive.ObjectID `json:"employee_id" bson:"employee_id"`
	UserProfileID primitive.ObjectID  `json:"user_profile_id" bson:"user_profile_id"`
}

type BodyCreate struct {
	CompanyID     primitive.ObjectID  `json:"company_id" bson:"company_id"`
	Role          string              `json:"role" bson:"role" enums:"HR_MANAGER,EMPLOYEE"`
	Username      string              `json:"username" bson:"username"`
	Email         string              `json:"email" bson:"email"`
	Password      string              `json:"password" bson:"password"`
	EmployeeID    *primitive.ObjectID `json:"employee_id" bson:"employee_id"`
	Status        string              `json:"status" bson:"status"`
	FirstNameEng  string              `json:"first_name_eng" bson:"first_name_eng"`
	UserProfileID primitive.ObjectID  `json:"user_profile_id" bson:"user_profile_id"`
}

type Status struct {
	ACTIVE          string `json:"ACTIVE" bson:"ACTIVE"`
	TO_SET_PASSWORD string `json:"TO_SET_PASSWORD" bson:"TO_SET_PASSWORD"`
	CANCEL          string `json:"CANCEL" bson:"CANCEL"`
}
