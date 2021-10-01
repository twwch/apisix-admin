package repo

import "apisix-admin/component/repo/user"

func Init()  {
	userRepo := user.NewRepoMongo()
	user.SetRepoInstance(userRepo)
}
