package database

import (
	"CRUD_Go_gin/internal/domain"
	"CRUD_Go_gin/pkg/postgres_database"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Books struct {
	db *sql.DB
}

func NewBooks() *Books {
	var b Books
	db, err := postgres_database.NewConnectionBooks()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"level": "internal database",
		}).Info("Can't init connection to DB books")
	}
	b.db = db
	return &b
}

func (books *Books) CreateBook(b *domain.Book) error {
	_, err := books.db.Exec("INSERT INTO books (title, author, year, publisher, count) VALUES ($1, $2, $3, $4, $5)",
		b.Title, b.Author, b.Year, b.Publisher, b.Count)
	if err != nil {
		return err
	}
	return nil
}

func (books *Books) GetBook(id int) (*domain.Book, error) {
	row := books.db.QueryRow("SELECT title, author, year, publisher, count FROM books WHERE id = $1", id)
	b := domain.Book{}
	err := row.Scan(&b.Title, &b.Author, &b.Year, &b.Publisher, &b.Count)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (books *Books) UpdateBook(b *domain.Book) error {
	_, err := books.db.Exec("UPDATE books SET count = $1 WHERE id = $2", b.Count, b.Id)
	if err != nil {
		return err
	}
	return nil
}

func (books *Books) DeleteBook(id int) error {
	_, err := books.db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
