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

// Exec Statement
// TODO: protect it from multiple access (use some mutex)
func Exec(query string, args ...any) error {

	db, err := connectDatabase(DatabaseConnectionInfo)

	if err != nil {
		return err
	}

	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

// Exec Query and return the rows
// Don't forget to close rows after using
// TODO: protect it from multiple access (use some mutex)
func ExecQuery(query string, args ...any) (*sql.Rows, error) {

	db, err := connectDatabase(DatabaseConnectionInfo)

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// TODO make some generic function that load the query return directly into some variable
