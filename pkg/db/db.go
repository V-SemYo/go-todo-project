package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
id INTEGER PRIMARY KEY AUTOINCREMENT,
date CHAR(8) NOT NULL DEFAULT "",
title VARCHAR(255) NOT NULL DEFAULT "",
comment TEXT NOT NULL DEFAULT "",
repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX IF NOT EXISTS index_date ON scheduler(date);
`

func InitDB(dbFile string) error {
	var err error
	DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("database %s opening error: %w", dbFile, err)
	}
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("database ping error: %w", err)
	}

	_, err = DB.Exec(schema)
	if err != nil {
		return fmt.Errorf("table creating error: %w", err)
	}
	return nil
}
