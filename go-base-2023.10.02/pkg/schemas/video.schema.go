package schemas

import (
	"gobase/pkg/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TVideo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	CreatedBy string             `bson:"created_by"`
	UpdatedBy string             `bson:"updated_by"`

	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	URL         string             `bson:"url,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
	Views       int                `bson:"views,omitempty"`
	Owner       primitive.ObjectID `bson:"owner,omitempty" json:"-"`
}

func GetVideoCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(helpers.GetENV().MONGO_INITDB_DATABASE).Collection("videos")
}
