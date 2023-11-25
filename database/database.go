package database

import (
	"database/sql"
	"fmt"
	"os"

	utils "github.com/alhazmy20/Go-WebAPI/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var EnvFilePath string

func OpenConnection() *sql.DB {
	if err := godotenv.Load(EnvFilePath); err != nil {
		utils.HandleError(err)
	}
	// ConnectionString for running the api on docker
	connectionString := fmt.Sprintf("%v:%v@tcp(%v)/%v",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_CONTAINER_NAME"), os.Getenv("DB_DATABASE"))

	// Connection for running the api on localhost
	// connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
	// 	os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	//connection for running the api on k8
	// connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
	// 	os.Getenv("DB_USERNAME"), os.Getenv("DB_ROOT_PASSWORD"), os.Getenv("DB_POD_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, err := sql.Open(os.Getenv("DB_CONNECTION"), connectionString)
	if err != nil {
		utils.HandleError(err)
	}
	return db
}
