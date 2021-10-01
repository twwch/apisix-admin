package user

import (
	"apisix-admin/component/repo/user"
	userEntity "apisix-admin/entity/user"
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

func (service *UserService) FindUserByIdOrEmail(ctx context.Context, key string, isUserId bool) (*userEntity.User, error) {
	userTemp := new(userEntity.User)
	var err error
	if isUserId {
		err = service.userRepo.GetBase().FindOne(ctx, bson.M{"user_id": key}, &userTemp)
	} else {
		key = strings.Replace(key, "@xiaoduotech.com", "", -1)
		email := fmt.Sprintf("%v@xiaoduotech.com", key)
		err = service.userRepo.GetBase().FindOne(ctx, bson.M{"email": email}, &userTemp)
	}
	return userTemp, err
}
