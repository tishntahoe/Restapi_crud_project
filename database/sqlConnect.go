package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"path/filepath"
)

var DbConnect *sql.DB

func ConnectToDatabase() {
	if DbConnect != nil {
		return
	}
	path, _ := filepath.Abs(filepath.Join("..", "database", "database"))
	fmt.Println(path)
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	DbConnect = db
}
