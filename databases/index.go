package databases

import "gorm.io/gorm"

var (
	Db *gorm.DB
)

func init() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
