package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"kratos-bbs-demo/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo, NewProfileRepo)


// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db:db}, cleanup, nil
}


func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{
			SkipDefaultTransaction: true,  // 关闭默认事务
			PrepareStmt: true,             // 缓存预编译语句
			DisableForeignKeyConstraintWhenMigrating: true,  // 不允许外键
	})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&Article{}); err != nil {
		panic(err)
	} 
	
	return db
}