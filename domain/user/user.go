package user

import (
	"apisix-admin/component/repo/user"
	"apisix-admin/entity/organization/user/mongo"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type UserService struct {
	userRepo user.IRepository
}

func NewUserService() *UserService {
	return &UserService{userRepo: user.GetRepoInstance()}
}

func (service *UserService) FindUserByPhoneOrEmail(ctx context.Context, key, password string) (*mongo.User, error) {
	userTemp := new(mongo.User)
	key = strings.Replace(key, "@xiaoduotech.com", "", -1)
	email := fmt.Sprintf("%v@xiaoduotech.com", key)
	err := service.userRepo.GetBase().FindOne(ctx, bson.M{"$or": []bson.M{
		{"email": email, "password": password},
		{"mobile": key, "password": password},
	}}, &userTemp)
	return userTemp, err
}

func  (service *UserService) FindUserByUserId (ctx context.Context, userId string) (*mongo.User, error) {
	userTemp := new(mongo.User)
	err := service.userRepo.GetBase().FindOne(ctx, bson.M{"user_id": userId}, &userTemp)
	return userTemp, err
}