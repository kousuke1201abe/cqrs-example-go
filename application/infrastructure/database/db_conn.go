package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	sqlMode                        = "'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'"
	mySQLDuplicateEntryErrorNumber = 1062
	connMaxLifetime                = 30 * time.Second
	connMaxIdleTime                = 30 * time.Second
)

type DBConn struct {
	db *sqlx.DB
}

func Connect(dbName, dbUser, dbPass, dbHost string, dbPort, maxIdleConns, maxOpneConns int) DBConn {
	cf := mysql.NewConfig()
	cf.DBName = dbName
	cf.User = dbUser
	cf.Passwd = dbPass
	cf.Addr = dbHost + ":" + strconv.Itoa(dbPort)
	cf.Net = "tcp"
	cf.ParseTime = true
	cf.Params = map[string]string{"sql_mode": sqlMode, "multiStatements": "true"}

	db, err := sqlx.Connect("mysql", cf.FormatDSN())
	if err != nil {
		cf.Passwd = "*****"
		fmt.Printf("unable to use data source name. dsn: %s", cf.FormatDSN())
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpneConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	return DBConn{db: db}
}

func (conn *DBConn) Close() error {
	return conn.db.Close()
}

func (conn *DBConn) Transaction(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) (err error) {
	tx, err := conn.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				fmt.Print("transactional function panicked and then failed to rollback")
			}
			panic(p)
		}
		if err != nil {
			if err := tx.Rollback(); err != nil {
				fmt.Print("failed to do transactional function and then failed to rollback")
			}
			return
		}

		if err := tx.Commit(); err != nil {
			fmt.Print("succeeded to do transactional function but failed to commit")
			if err := tx.Rollback(); err != nil {
				fmt.Print("failed to commit and then failed to rollback")
			}
		}
	}()

	if err := f(ctx, tx); err != nil {
		return err
	}
	return nil
}
