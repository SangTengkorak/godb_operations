package dboperations

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB instance
var DB *sql.DB

func DBconnections() {
	err := godotenv.Load("./local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBusername := os.Getenv("DBUSER")
	DBpassword := os.Getenv("DBUPASS")
	DBname := os.Getenv("DBNAME")
	DBhost := os.Getenv("DBHOST")

	db, err := sql.Open("mysql", DBusername+":"+DBpassword+"@tcp("+DBhost+":3306)/"+DBname)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
