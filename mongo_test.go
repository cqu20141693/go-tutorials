package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cqu20141693/go-tutorials/constant"
	"github.com/cqu20141693/go-tutorials/db/mongo"
	"github.com/cqu20141693/go-tutorials/domain"
	config "github.com/cqu20141693/go-tutorials/resource"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func init() {
	//config.Init()
	mongo.Init()
}
func TestCRUD(t *testing.T) {
	tp := &domain.Topology{Name: "test", Desc: "test", Data: struct{}{}, UserID: "1", Username: "cc", Tags: []string{"test"}, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	marshal, err := json.Marshal(tp)
	if err != nil {
		fmt.Println("json Marshal failed")
		return
	}
	var data = bson.M{}
	json.Unmarshal(marshal, &data)
	err, id := TopologyAdd(data)
	if err != nil {
		fmt.Println("topology add failed")
		return
	}
	// Rest of the code will go here
	StarAdd(&domain.Star{TopologyID: id.String()}, "1")

}

// StarAdd 点赞
func StarAdd(data *domain.Star, uid string) (err error) {
	if data.TopologyID == "" {
		err = errors.New(constant.ErrorID)
		return
	}

	coll1 := mongo.Client.Database(config.App.Mongo.Database).Collection(constant.Stars)

	data.ID = primitive.NewObjectID()
	data.CreatedAt = time.Now().UTC()
	data.UserID = uid
	_, err = coll1.InsertOne(context.TODO(), data)
	if err != nil {
		log.Error().Caller().Err(err).Str("func", "topology.StarAdd.Insert").Msgf("Fail to write mongo(Stars): data=%v", data)
		return
	}

	coll2 := mongo.Client.Database(config.App.Mongo.Database).Collection(constant.Topologies)
	_id, _ := primitive.ObjectIDFromHex(data.TopologyID)
	_, err = coll2.UpdateByID(context.TODO(), _id, bson.M{"$inc": bson.M{"star": 1}})

	if err != nil {
		log.Error().Caller().Err(err).Str("func", "topology.StarAdd.Inc").Msgf("Fail to write mongo(Topologies): data=%v", data)
	}

	return
}

func TopologyAdd(data bson.M) (err error, id primitive.ObjectID) {
	coll := mongo.Client.Database(config.App.Mongo.Database).Collection(constant.Topologies)
	data["updatedAt"] = time.Now().UTC()
	data["_id"] = primitive.NewObjectID()
	data["createdAt"] = data["updatedAt"]
	data["star"] = 0
	data["shared"] = false

	ret, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		return err, primitive.ObjectID{}
	}
	return nil, ret.InsertedID.(primitive.ObjectID)
}

func TestAPI(t *testing.T) {
	id := primitive.NewObjectID()
	fmt.Println(id.Hex(), id.String())

	ret := make(map[string]interface{})
	ret["image"] = "/image/test.png"
	fmt.Println(ret["image"].(string))
}
