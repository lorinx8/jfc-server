package nfdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	_db *sql.DB
)

func GetConn() (db *sql.DB, err error) {
	if _db != nil {
		err := _db.Ping()
		if err != nil {
			fmt.Println("_db ping failed")
			_db.Close()
			open()
		}
		return _db, nil
	} else {
		open()
	}
	return _db, nil
}

func open() {
	__db, err1 := sql.Open("postgres", "user=pguser1 password=pguser1 host=dev1.papakaka.com port=5431 dbname=j2map sslmode=disable")
	__db.SetMaxIdleConns(10)
	__db.SetMaxOpenConns(50)
	if err1 != nil {
		log.Fatal(err1)
	} else {
		_db = __db
	}
}
