package tokenResetPasswordRepo

import (
	"log"

	"github.com/Olapat/go-architecture/model/entity"
	tokenResetPasswordEntity "github.com/Olapat/go-architecture/model/entity/token_reset_password"
	"github.com/Olapat/go-architecture/model/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName = "token_reset_password"

func Insert(c *fiber.Ctx, rawBody interface{}) (interface{}, error) {
	result, err := repository.SuperInsertOne(c, collectionName, rawBody)
	return result, err
}

func BuildRows(rows []tokenResetPasswordEntity.TokenResetPassword) []tokenResetPasswordEntity.TokenResetPassword {
	var record = make([]tokenResetPasswordEntity.TokenResetPassword, 0)
	for _, v := range rows {
		record = append(record, tokenResetPasswordEntity.TokenResetPassword{
			ID:        v.ID,
			Email:     v.Email,
			UserID:    v.UserID,
			CreatedAt: v.CreatedAt,
			ExpiresAt: v.ExpiresAt,
			Token:     v.Token,
			IsConfirm: v.IsConfirm,
		})
	}

	return record
}

func Find(c *fiber.Ctx, filter primitive.M, pagination *entity.PaginationRequests, sort *primitive.M) ([]tokenResetPasswordEntity.TokenResetPassword, error) {
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

	var results []tokenResetPasswordEntity.TokenResetPassword = make([]tokenResetPasswordEntity.TokenResetPassword, 0)

	filter["deleted_by"] = nil

	cursor, err := repository.SuperFind(c, collectionName, filter, opts)
	if err != nil {
		log.Println(err)

		return results, err
	}

	cursor.All(c.Context(), &results)

	return results, err
}

func FindOne(c *fiber.Ctx, filter primitive.M) (tokenResetPasswordEntity.TokenResetPassword, error) {
	var entity tokenResetPasswordEntity.TokenResetPassword
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

func FindByToken(c *fiber.Ctx, token string) (tokenResetPasswordEntity.TokenResetPassword, error) {
	var entity tokenResetPasswordEntity.TokenResetPassword
	filter := bson.M{
		"token": token,
	}
	errFind := repository.SuperFindOne(c, collectionName, filter, &entity)
	return entity, errFind
}
