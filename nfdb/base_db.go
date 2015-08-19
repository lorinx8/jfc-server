package nfdb

import (
	"database/sql"
	"jfcsrv/nflog"
	"os"

	_ "github.com/lib/pq"
)

var (
	_db *sql.DB
)

var jlog = nflog.Logger

func GetConn() (db *sql.DB, err error) {
	if _db != nil {
		err := _db.Ping()
		if err != nil {
			jlog.Error("_db ping failed")
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
		jlog.Critical("db open failed:", err1)
		os.Exit(1)
	} else {
		_db = __db
	}
}
