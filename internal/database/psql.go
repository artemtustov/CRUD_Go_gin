package database

import (
	"CRUD_Go_gin/internal/domain"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDatabase() error {
	connStr := "user=postgres password=12345 dbname=books sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func CreateBook(b *domain.Book) error {
	_, err := db.Exec("INSERT INTO books (title, author, year, publisher, count) VALUES ($1, $2, $3, $4, $5)",
		b.Title, b.Author, b.Year, b.Publisher, b.Count)
	if err != nil {
		return err
	}
	return nil
}

func GetBook(id int) (*domain.Book, error) {
	row := db.QueryRow("SELECT title, author, year, publisher, count FROM books WHERE id = $1", id)
	b := domain.Book{}
	err := row.Scan(&b.Title, &b.Author, &b.Year, &b.Publisher, &b.Count)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func UpdateBook(b *domain.Book) error {
	_, err := db.Exec("UPDATE books SET count = $1 WHERE id = $2", b.Count, b.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(id int) error {
	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
