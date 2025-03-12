package db

import (
	"database/sql" //Đây là gói tiêu chuẩn của Golang để làm việc với database t
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) { // khoi tao tables
	var err error
	DB, err = sql.Open("sqlite", "database.db")
	if err != nil {
		return DB, fmt.Errorf("could not connect to database: %w", err)
	}
	CreateTables()
	return DB, nil
}

func CreateTables() {
	CreateImageTable := `
	CREATE TABLE IF NOT EXISTS images (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT ,
		path1 TEXT NOT NULL,
		path2 TEXT NOT NULL,
		text TEXT NOT NULL,
		width INTEGER NOT NULL,
		uploaded_at TEXT DEFAULT CURRENT_TIMESTAMP,
		updated_at TEXT DEFAULT CURRENT_TIMESTAMP 
	);`
	_, err := DB.Exec(CreateImageTable)
	if err != nil {
		panic(err.Error())
	}
	CreateTrigger := `
	CREATE TRIGGER IF NOT EXISTS update_images_timestamp
	AFTER UPDATE ON images
	FOR EACH ROW
	BEGIN
	    UPDATE images SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
	END;`
	_, err = DB.Exec(CreateTrigger)
	if err != nil {
		panic(err.Error())
	}
}
