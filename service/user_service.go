package service

import (
	"context"
	"fmt"

	"grpcTest4/configs"
	pb "grpcTest4/proto"
)

var db = configs.NewDBHandler()

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (service *UserServiceServer) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.UserResponse, error) {
	resp, err := db.GetUser(req.Id)

	if err != nil {
		return nil, fmt.Errorf("get user by id service error: %s", err)
	}

	return &pb.UserResponse{Id: resp.Id.String(), Name: resp.Name, Email: resp.Email, Password: resp.Password}, nil
}

func (service *UserServiceServer) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.UserResponse, error) {
	resp, err := db.GetUserByEmail(req.Email)

	if err != nil {
		return nil, fmt.Errorf("get user by email service error: %s", err)
	}

	return &pb.UserResponse{Id: resp.Id.String(), Name: resp.Name, Email: resp.Email, Password: resp.Password}, nil
}

func (service *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := configs.User{Name: req.Name, Email: req.Email, Password: req.Password}
	_, err := db.CreateUser(newUser)

	if err != nil {
		return nil, fmt.Errorf("create user service error: %s", err)
	}

	return &pb.CreateUserResponse{Data: "User created successful"}, nil
}

func (service *UserServiceServer) GetUsers(context.Context, *pb.Empty) (*pb.GetAllUsersResponse, error) {
	resp, err := db.GetAllUsers()
	var users []*pb.UserResponse

	if err != nil {
		return nil, fmt.Errorf("get users service error: %s", err)
	}

	for _, v := range resp {
		var singleUser = &pb.UserResponse{
			Id:       v.Id.String(),
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
		}
		users = append(users, singleUser)
	}

	return &pb.GetAllUsersResponse{Users: users}, nil
}

func (service *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	newUser := configs.User{Name: req.Name, Email: req.Email, Password: req.Password}
	_, err := db.UpdateUser(req.Id, newUser)

	if err != nil {
		return nil, fmt.Errorf("update user service error: %s", err)
	}

	return &pb.UpdateUserResponse{Data: "Users updated successful"}, nil
}

func (service *UserServiceServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := db.DeleteUser(req.Id)

	if err != nil {
		return nil, fmt.Errorf("delete user service error: %s", err)
	}

	return &pb.DeleteUserResponse{Data: "User detail delete successful"}, nil
}
