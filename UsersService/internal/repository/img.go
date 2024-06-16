package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImgRepo struct {
	db *mongo.Database
}

type ImgStruct struct {
	Img string `bson:"img"`
}

func NewImgRepo(db *mongo.Database) *ImgRepo {
	return &ImgRepo{db: db}
}

func (r *ImgRepo) Create(img string) (string, error) {
	coll := r.db.Collection(IMG_COLLECTION)
	resOfSearch := coll.FindOne(context.Background(), bson.M{"img": img})
	if resOfSearch.Err() == mongo.ErrNoDocuments {
		res, err := coll.InsertOne(context.Background(), bson.M{"img": img})
		if err != nil {
			return "", err
		}
		id, ok := res.InsertedID.(primitive.ObjectID)
		if !ok {
			return "", errors.New("error represent insert id as primitive.ObjectID")
		}
		return id.Hex(), nil
	}
	return "", errors.New("img already exist")
}

func (r *ImgRepo) GetById(id string) (string, error) {
	coll := r.db.Collection(IMG_COLLECTION)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	filter := bson.M{"_id": objId}
	res := coll.FindOne(context.Background(), filter)
	if res.Err() != nil {
		return "", res.Err()
	}
	var img ImgStruct
	err = res.Decode(&img)
	if err != nil {
		return "", err
	}
	return img.Img, nil
}

func (r *ImgRepo) GetByIDArray(idArr []string) ([]string, error) {
	coll := r.db.Collection(IMG_COLLECTION)
	objIdArr := make([]primitive.ObjectID, len(idArr))
	for i, id := range idArr {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objIdArr[i] = objId
	}
	filter := bson.M{"_id": bson.M{"$in": objIdArr}}
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var imgs []string
	for cursor.Next(context.Background()) {
		var img ImgStruct
		err = cursor.Decode(&img)
		if err != nil {
			return nil, err
		}
		imgs = append(imgs, img.Img)
	}
	return imgs, nil
}
func (r *ImgRepo) Update(id, img string) error {
	coll := r.db.Collection(IMG_COLLECTION)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"img": img}}
	_, err = coll.UpdateOne(context.Background(), filter, update)
	return err
}
func (r *ImgRepo) Delete(id string) error {
	coll := r.db.Collection(IMG_COLLECTION)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objId}
	_, err = coll.DeleteOne(context.Background(), filter)
	return err
}
