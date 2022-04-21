package gromimplement

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDBCockroach() *gorm.DB {
	addr := "postgresql://fuckmeyes:3Lfyik3NHBJ7v8s4Hb5RaQ@free-tier8.aws-ap-southeast-1.cockroachlabs.cloud:26257/defaultdb?sslmode=prefer&options=--cluster%3Dkind-cheetah-1261"
	db, err := gorm.Open(postgres.Open(addr))
	if err != nil {
		log.Fatal("Cockroach 链接失败")
	}
	log.Println("CockroachDB connected")
	// db.AutoMigrate(&Episode{}, &Movie{}, &Performer{}, &Episodes2Movie{}, &Episodes2Performers{})
	return db
}
