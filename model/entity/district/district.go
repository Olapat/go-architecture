package districtEntity

import "go.mongodb.org/mongo-driver/bson/primitive"

type District struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	DID        int64              `json:"id" bson:"id"`
	Code       string             `json:"code" bson:"code"`
	GeoID      int64              `json:"geo_id" bson:"geo_id"`
	ProvinceID int64              `json:"province_id" bson:"province_id"`
	Name       string             `json:"name" bson:"name"`
	NameTH     string             `json:"name_th" bson:"name_th"`
}
