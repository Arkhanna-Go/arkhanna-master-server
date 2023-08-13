package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB_INFO struct {
	Username string
	Password string
	Database string
}

type CONN_INFO struct {
	DB_INFO
	Protocol string `default:"tcp"`
	Hostname string `default:"127.0.0.1"`
	Port     uint16 `default:"3306"`
}

const BaseConnectionSourceFormat string = "%s:%s@%s(%s:%d)/%s"

var DatabaseConnectionInfo CONN_INFO

// Convert CONN_INFO to connection source string
func (conn_info CONN_INFO) databaseInfoToString() string {
	return fmt.Sprintf(BaseConnectionSourceFormat,
		conn_info.Username,
		conn_info.Password,
		conn_info.Protocol,
		conn_info.Hostname,
		conn_info.Port,
		conn_info.Database)
}

// Connect to database
func connectDatabase(conn_info CONN_INFO) (*sql.DB, error) {
	db, err := sql.Open("mysql", conn_info.databaseInfoToString())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
