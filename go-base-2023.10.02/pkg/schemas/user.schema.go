package schemas

import (
	"gobase/pkg/helpers"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	CreatedBy string             `bson:"created_by"`
	UpdatedBy string             `bson:"updated_by"`

	UserName string `bson:"user_name,omitempty" unique:"true"`
	Password string `bson:"password,omitempty"`
	Email    string `bson:"email,omitempty"`
	Name     string `bson:"name,omitempty"`
}

func GetUserCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(helpers.GetENV().MONGO_INITDB_DATABASE).Collection("users")
}
