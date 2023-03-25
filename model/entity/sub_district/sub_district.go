package subDistrictEntity

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubDistrict struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	SID        int64              `json:"id" bson:"id"`
	ProvinceID int64              `json:"province_id" bson:"province_id"`
	AmphureID  int64              `json:"amphure_id" bson:"amphure_id"`
	Name       string             `json:"name" bson:"name"`
	NameTH     string             `json:"name_th" bson:"name_th"`
	PostCode   string             `json:"post_code" bson:"post_code"`
}
