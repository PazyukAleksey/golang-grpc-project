package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbHandler interface {
	GetUser(Id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) (*mongo.InsertOneResult, error)
	UpdateUser(id string, user User) (*mongo.UpdateResult, error)
	DeleteUser(id string) (*mongo.DeleteResult, error)
	GetAllUsers() ([]*User, error)
}

type DB struct {
	client *mongo.Client
}

func NewDBHandler() dbHandler {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoUrl()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return &DB{client: client}
}

func colHelper(db *DB) *mongo.Collection {
	return db.client.Database("user_service_data").Collection("users")
}

func (db *DB) CreateUser(user User) (*mongo.InsertOneResult, error) {
	col := colHelper(db)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newUser := User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := col.InsertOne(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("create user error: %s", err)
	}

	return res, nil
}

func (db *DB) GetUser(id string) (*User, error) {
	col := colHelper(db)
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	err := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)

	if err != nil {
		return nil, fmt.Errorf("get user error: %s", err)
	}

	return &user, nil
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	col := colHelper(db)
	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := col.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, fmt.Errorf("get user by email error: %s", err)
	}

	return &user, nil
}

func (db *DB) UpdateUser(id string, user User) (*mongo.UpdateResult, error) {
	col := colHelper(db)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{"name": user.Name, "email": user.Email, "password": user.Password}
	result, err := col.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})

	if err != nil {
		return nil, fmt.Errorf("update user error: %s", err)
	}

	return result, err
}

func (db *DB) DeleteUser(id string) (*mongo.DeleteResult, error) {
	col := colHelper(db)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := col.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return nil, fmt.Errorf("delete user error: %s", err)
	}

	return result, err
}

func (db *DB) GetAllUsers() ([]*User, error) {
	col := colHelper(db)
	var users []*User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := col.Find(ctx, bson.M{})

	if err != nil {
		return nil, fmt.Errorf("get all user error: %s", err)
	}

	for result.Next(ctx) {
		var singleUser *User
		if err = result.Decode(&singleUser); err != nil {
			return nil, fmt.Errorf("decode error: %s", err)
		}
		users = append(users, singleUser)
	}
	return users, err
}
