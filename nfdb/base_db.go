package nfdb

import (
	"database/sql"
	"jfcsrv/nfconst"
	"jfcsrv/nflog"

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
			jlog.Error("_db ping failed", err)
			_db.Close()
			_db = nil
			open()
		}
		return _db, nil
	} else {
		open()
	}
	return _db, nil
}

func open() {
	__db, err1 := sql.Open("postgres", nfconst.JCfg.DbConnString)
	__db.SetMaxIdleConns(5)
	__db.SetMaxOpenConns(50)
	if err1 != nil {
		jlog.Critical("db open failed:", err1)
		__db.Close()
		__db = nil
	} else {
		_db = __db
	}
}
