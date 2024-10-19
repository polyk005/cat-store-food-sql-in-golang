package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// функция подключения к базе данных
func Connect() error {
	connStr := "user=username dbname=catfoodstore sslmode=disable password=yourpassword"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к БД: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	DB = db
	log.Println("Успешно подключились к базе данных")
	return nil
}
