package repository

import (
	"log"
	"time"

	"github.com/Olapat/go-architecture/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SuperFind(c *fiber.Ctx, collectionName string, filter primitive.M, opts *options.FindOptions) (*mongo.Cursor, error) {
	coll := db.MG.Database.Collection(collectionName)

	cursor, err := coll.Find(c.Context(), filter, opts)
	if err != nil {
		log.Println(err)

		return cursor, err
	}

	return cursor, nil
}

func SuperFindOne(c *fiber.Ctx, collectionName string, filter primitive.M, result interface{}) error {
	coll := db.MG.Database.Collection(collectionName)

	err := coll.FindOne(c.Context(), filter).Decode(result)
	if err != nil {
		log.Println(err)

		return err
	}

	return err
}

func SuperFindOneByID(c *fiber.Ctx, collectionName string, id string, result interface{}) error {
	coll := db.MG.Database.Collection(collectionName)
	ID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}

	filter := bson.M{
		"_id": ID,
	}

	err := coll.FindOne(c.Context(), filter).Decode(result)
	if err != nil {
		log.Println(err)

		return err
	}

	return err
}

func SuperInsertOne(c *fiber.Ctx, collectionName string, rawBody interface{}) (interface{}, error) {
	optsSession := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := db.MG.Client.StartSession(optsSession)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.EndSession(c.Context())

	txnOpts := options.Transaction().SetReadPreference(readpref.PrimaryPreferred())
	result, err := sess.WithTransaction(c.Context(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		var coll = db.MG.Database.Collection(collectionName)
		opts := options.InsertOne()
		res, err := coll.InsertOne(c.Context(), rawBody, opts)
		if err != nil {
			log.Println(err)
		}

		return res.InsertedID, err
	}, txnOpts)

	return result, err
}

func SuperInsertMany(c *fiber.Ctx, collectionName string, rawBody []interface{}) (interface{}, error) {
	optsSession := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := db.MG.Client.StartSession(optsSession)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.EndSession(c.Context())

	txnOpts := options.Transaction().SetReadPreference(readpref.PrimaryPreferred())
	result, err := sess.WithTransaction(c.Context(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		var coll = db.MG.Database.Collection(collectionName)
		opts := options.InsertMany().SetOrdered(false)
		res, err := coll.InsertMany(c.Context(), rawBody, opts)
		if err != nil {
			log.Println(err)
		}

		return res.InsertedIDs, err
	}, txnOpts)

	return result, err
}

func SuperDeleteMany(c *fiber.Ctx, collectionName string, filter primitive.M) (interface{}, error) {
	optsSession := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := db.MG.Client.StartSession(optsSession)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.EndSession(c.Context())

	txnOpts := options.Transaction().SetReadPreference(readpref.PrimaryPreferred())
	result, err := sess.WithTransaction(c.Context(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		var coll = db.MG.Database.Collection(collectionName)
		res, err := coll.DeleteMany(c.Context(), filter)
		if err != nil {
			log.Println(err)
		}

		return res.DeletedCount, err
	}, txnOpts)

	return result, err
}

func SuperCount(c *fiber.Ctx, collectionName string, filter primitive.M, opts *options.CountOptions) (int64, error) {
	coll := db.MG.Database.Collection(collectionName)

	cursor, err := coll.CountDocuments(c.Context(), filter, opts)
	if err != nil {
		log.Println(err)

		return cursor, err
	}

	return cursor, nil
}

func SuperAggregate(c *fiber.Ctx, collectionName string, pipeline mongo.Pipeline) (*mongo.Cursor, error) {
	var coll = db.MG.Database.Collection(collectionName)
	cursor, err := coll.Aggregate(c.Context(), pipeline)
	return cursor, err
}

func SuperUpdate(c *fiber.Ctx, collectionName string, filter primitive.M, update primitive.M, upsert bool) (interface{}, error) {
	optsSession := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := db.MG.Client.StartSession(optsSession)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.EndSession(c.Context())

	txnOpts := options.Transaction().SetReadPreference(readpref.PrimaryPreferred())
	result, err := sess.WithTransaction(c.Context(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		var coll = db.MG.Database.Collection(collectionName)
		opts := options.Update().SetUpsert(upsert)
		res, err := coll.UpdateMany(c.Context(), filter, update, opts)
		if err != nil {
			log.Println(err)
		}

		return res.MatchedCount, err
	}, txnOpts)

	return result, err
}

func SuperSoftDelete(c *fiber.Ctx, collectionName string, filter primitive.M, by *primitive.ObjectID) (interface{}, error) {
	optsSession := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := db.MG.Client.StartSession(optsSession)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.EndSession(c.Context())

	txnOpts := options.Transaction().SetReadPreference(readpref.PrimaryPreferred())
	result, err := sess.WithTransaction(c.Context(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		var coll = db.MG.Database.Collection(collectionName)
		opts := options.Update().SetUpsert(false)
		update := bson.M{
			"deleted_at": time.Now(),
			"deleted_by": by,
		}
		res, err := coll.UpdateMany(c.Context(), filter, update, opts)
		if err != nil {
			log.Println(err)
		}

		return res.MatchedCount, err
	}, txnOpts)

	return result, err
}
