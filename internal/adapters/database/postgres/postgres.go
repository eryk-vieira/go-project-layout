package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var databaseInstance *sql.DB

func init() {
	databaseInstance = createDatabaseInstance()
}

func createDatabaseInstance() *sql.DB {
	var (
		host    = os.Getenv("DATABASE_HOST")
		port    = os.Getenv("DATABASE_PORT")
		user    = os.Getenv("DATABASE_USER")
		pass    = os.Getenv("DATABASE_PASS")
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, pass)
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error while opening postgresql reader database: %s\n", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error while ping postgresql reader database: %s\n", err.Error())
	}

	log.Println("Connected to reader database:", host)

	return db
}

func GetDatabaseInstance() *sql.DB {
	return databaseInstance
}
