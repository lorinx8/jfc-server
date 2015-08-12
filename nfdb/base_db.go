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

func getDb() (db *sql.DB, err error) {
	if _db != nil {
		fmt.Println("_db not nil, return exists")
		return _db, nil
	} else {
		fmt.Println("_db nil, open it")
		__db, err1 := sql.Open("postgres", "user=pguser1 password=pguser1 host=dev1.papakaka.com port=5431 dbname=paka sslmode=disable")
		if err1 != nil {
			log.Fatal(err1)
		} else {
			_db = __db
		}
	}
	return _db, nil
}
