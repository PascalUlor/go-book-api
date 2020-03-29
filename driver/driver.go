package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// this package
	_ "github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	var (
		host     = os.Getenv("PG_HOST")
		port     = os.Getenv("PG_PORT")
		user     = os.Getenv("PG_USER")
		password = os.Getenv("PG_PASSWORD")
		dbname   = os.Getenv("PG_DB")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// pgURL, err := pq.ParseURL(psqlInfo)
	// logFatal(err)

	db, err := sql.Open("postgres", psqlInfo)
	logFatal(err)

	// create database if not exists

	// statement := `SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = 'yourDBName');`

	// row := db.QueryRow(statement)
	// var exists bool
	// err = row.Scan(&exists)
	// logFatal(err)

	// if exists == false {
    // statement = `CREATE DATABASE yourDBName;`
    // _, err = db.Exec(statement)
    // logFatal(err)
	// }

	// create table if it does not exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books ( id serial, title varchar(32), author varchar(32), year varchar(32) )")
	if err != nil {
		logFatal(err)
	}

	// _, err = db.Exec("insert into books (title, author, year) values('Juju Rock', 'Cyprian Ekwensi', '1975')")
	// if err != nil {
	// 	logFatal(err)
	// }

	err = db.Ping()
	logFatal(err)

	log.Println(psqlInfo)

	return db
}
