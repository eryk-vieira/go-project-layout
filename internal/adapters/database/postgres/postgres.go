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
		host    = os.Getenv("DATABASE_READER_HOST")
		port    = os.Getenv("DATABASE_READER_PORT")
		user    = os.Getenv("DATABASE_READER_USER")
		pass    = os.Getenv("DATABASE_READER_PASS")
		dbname  = os.Getenv("DATABASE_READER_NAME")
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error while opening '%s' postgresql reader database: %s\n", dbname, err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error while ping '%s' postgresql reader database: %s\n", dbname, err.Error())
	}

	log.Println("Connected to reader database:", host)

	return db
}

func GetDatabaseInstance() *sql.DB {
	return databaseInstance
}
