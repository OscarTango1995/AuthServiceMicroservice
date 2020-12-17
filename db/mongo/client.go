package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/OscarTango1995/AuthServiceMicroservice/config"
	"github.com/OscarTango1995/AuthServiceMicroservice/db"
	"github.com/OscarTango1995/AuthServiceMicroservice/models"
)

const (
	sessionCollection = "session"
	userCollection    = "user"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mysql database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}
func (c client) CreateUser(user *models.User) (string, error) {
	if user.ID != "" {
		return "", errors.New("id is not empty")
	}

	user.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
	if _, err := collection.InsertOne(context.TODO(), user); err != nil {
		return "", errors.Wrap(err, "failed to add user")
	}

	return user.ID, nil
}

func (c client) ListUser(id string) (*models.User, error) {
	var user *models.User
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c client) DeleteUser(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

func (c client) UpdateUser(user *models.User) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{"$set": user}); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (c client) CreateSession(session *models.Session) (string, error) {
	if session.ID != "" {
		return "", errors.New("id is not empty")
	}

	session.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(sessionCollection)
	if _, err := collection.InsertOne(context.TODO(), session); err != nil {
		return "", errors.Wrap(err, "failed to add user")
	}

	return session.ID, nil
}

func (c client) ListSession(id string) (*models.Session, error) {
	var session *models.Session
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(sessionCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&session); err != nil {
		return nil, err
	}
	return session, nil
}

func (c client) DeleteSession(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(sessionCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

func (c client) UpdateSession(session *models.Session) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": session.ID}, bson.M{"$set": session}); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}
