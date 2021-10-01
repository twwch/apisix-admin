package user

import "apisix-admin/component/repo/base"

type IRepository interface {
	GetBase() *base.MongoBase
}

var __repoInstance IRepository

func SetRepoInstance(repo IRepository) {
	__repoInstance = repo
}

func GetRepoInstance() IRepository {
	return __repoInstance
}