package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*gorm.DB,error){
	if err := godotenv.Load();err != nil{
		log.Print("ENV file not found", err)
	}

	user:= os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASS")
	host:= os.Getenv("DB_HOST")
	port:= os.Getenv("DB_PORT")
	db_name:= os.Getenv("DB_NAME")


	if user == "" || host == "" || db_name == ""{
		return nil,fmt.Errorf("Missing ENV fields")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		pwd,
		host,
		port,
		db_name,
	)

	db , err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		return nil,err
	}

	sqlDB,err:= db.DB()
	if err != nil{
		return nil,err
	}

	if err := sqlDB.Ping();err!= nil{
		return nil, err
	}

	return db,nil
}