package service_test

import (
	"context"
	"testing"

	"grpcTest4/configs"
	pb "grpcTest4/proto"
	"grpcTest4/service"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserServiceServer(t *testing.T) {

	objectID, err := primitive.ObjectIDFromHex("65b27360501f41e42a0f5a7a")
	assert.NoError(t, err)

	testUser := configs.User{
		Id:       objectID,
		Name:     "Oleksii",
		Email:    "paziukal@gmail.com",
		Password: "Root1234",
	}

	userService := &service.UserServiceServer{}

	req := &pb.GetUserByEmailRequest{Email: testUser.Email}
	resp, err := userService.GetUserByEmail(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, testUser.Name, resp.Name)
	assert.Equal(t, testUser.Email, resp.Email)
	assert.Equal(t, testUser.Password, resp.Password)

}
