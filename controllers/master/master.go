package masterCtl

import (
	"strconv"

	"github.com/Olapat/go-architecture/controllers/response"
	_ "github.com/Olapat/go-architecture/model/entity/district"
	_ "github.com/Olapat/go-architecture/model/entity/province"
	_ "github.com/Olapat/go-architecture/model/entity/sub_district"
	districtRepository "github.com/Olapat/go-architecture/model/repository/district"
	provinceRepository "github.com/Olapat/go-architecture/model/repository/province"
	subDistrictRepository "github.com/Olapat/go-architecture/model/repository/sub_district"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Master godoc
// @Tags        Master
// @Summary     Province
// @Description Master Province
// @Produce     json
// @Success     200 {object} entity.Response200{data=[]provinceEntity.Province}
// @Failure     500 {object} entity.Response500
// @Router      /master/province [get]
func Province(c *fiber.Ctx) error {
	filter := bson.M{}
	records, err := provinceRepository.Find(c, filter, nil, nil)
	if err != nil {
		return response.ResponseError(c, fiber.StatusInternalServerError, err.Error(), "")
	}
	return response.ResponseOK(c, fiber.StatusOK, records, "")
}

// Master godoc
// @Tags        Master
// @Summary     District
// @Description Master District
// @Produce     json
// @Param       province_id path     string true "Province Id"
// @Success     200         {object} entity.Response200{data=[]districtEntity.District}
// @Failure     500         {object} entity.Response500
// @Router      /master/district/{province_id} [get]
func District(c *fiber.Ctx) error {
	province_id := c.Params("province_id")
	intVar, _ := strconv.Atoi(province_id)
	filter := bson.M{
		"province_id": intVar,
	}

	records, err := districtRepository.Find(c, filter, nil, nil)
	if err != nil {
		return response.ResponseError(c, fiber.StatusInternalServerError, err.Error(), "")
	}
	return response.ResponseOK(c, fiber.StatusOK, records, "")
}

// Master godoc
// @Tags        Master
// @Summary     SubDistrict
// @Description Master SubDistrict
// @Produce     json
// @Param       province_id     path     string true "Province Id"
// @Param       district_id path     string true "District Id"
// @Success     200             {object} entity.Response200{data=[]subDistrictEntity.SubDistrict}
// @Failure     500             {object} entity.Response500
// @Router      /master/sub_district/{province_id}/{district_id} [get]
func SubDistrict(c *fiber.Ctx) error {
	province_id := c.Params("province_id")
	district_id := c.Params("district_id")
	intProvince, _ := strconv.Atoi(province_id)
	intSubDistrict, _ := strconv.Atoi(district_id)
	filter := bson.M{
		"province_id": intProvince,
		"amphure_id":  intSubDistrict,
	}

	records, err := subDistrictRepository.Find(c, filter, nil, nil)
	if err != nil {
		return response.ResponseError(c, fiber.StatusInternalServerError, err.Error(), "")
	}
	return response.ResponseOK(c, fiber.StatusOK, records, "")
}
