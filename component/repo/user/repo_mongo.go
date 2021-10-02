package user

import (
	"apisix-admin/component/repo/base"
	"apisix-admin/config"
)

type repoMongo struct {
	*base.MongoBase
}

// NewRepoMongo NewRepo creates a new repository.
func NewRepoMongo() *repoMongo {
	return &repoMongo{
		base.NewMongoBase(config.Get().MongoConf["default"], "eff-publish", "user").SetSoftDelete(false),
	}
}

func (r *repoMongo) GetBase() *base.MongoBase{
	return r.MongoBase
}