package mysqlDataSource

import (
	"fmt"

	"github.com/IqbalLx/alterra-agmc/config"

	s "github.com/IqbalLx/alterra-agmc/dataSources/mysql/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&tls=true",
		config.Database.User,
		config.Database.Pwd,
		config.Database.Host,
		config.Database.Port,
		config.Database.Table,
	)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(&s.User{}, &s.Book{})

	return DB
}
