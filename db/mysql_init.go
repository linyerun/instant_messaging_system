package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var MySQLClient = InitMySQL(strings.Join([]string{"root", ":", "123456", "@tcp(", "localhost", ":", "3306", ")/", "im_system", "?charset=utf8mb4&parseTime=true&loc=Local&timeout=10s&readTimeout=30s&writeTimeout=60s"}, ""))

func InitMySQL(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true, //预编译语句
		Logger:                 newGormLogger(),
	})
	if err != nil {
		panic(err)
	}
	return db
}

func newGormLogger() logger.Interface {
	//创建文件夹
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileDir := rootDir + "/im_system_logs"
	if _, err = os.Stat(fileDir); os.IsNotExist(err) {
		//不存在这个文件夹就创建
		err := os.MkdirAll(fileDir, 0666) //可读写，不可执行
		if err != nil {
			panic(err)
		}
	}

	//创建文件
	fileName := "gorm.log"
	filePath := path.Join(fileDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	return logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,
		},
	)
}
