package validate

import (
	"github.com/Olapat/go-architecture/model/entity"

	strUtil "github.com/Olapat/go-architecture/utils/string"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(body interface{}) []entity.ErrorResponse {
	var errors []entity.ErrorResponse
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element entity.ErrorResponse
			field := strUtil.New{V: err.Field()}.CamelToSnake()
			element.FailedField = field
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}
	return errors
}
