package auth

import (
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbin(db *gorm.DB) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("casbin adapter init failed: %v", err)
	}
	e, err := casbin.NewEnforcer("configs/model.conf", adapter)
	if err != nil {
		log.Fatalf("casbin enforcer init failed: %v", err)
	}
	if err := e.LoadPolicy(); err != nil {
		log.Fatalf("casbin load policy failed: %v", err)
	}
	return e
}
