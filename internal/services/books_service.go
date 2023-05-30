package services

import (
	"CRUD_Go_gin/internal/database"
	"CRUD_Go_gin/internal/domain"
)

var Books *database.Books

func CreateBook(b *domain.Book) error {
	err := Books.CreateBook(b)
	if err != nil {
		return err
	}
	return nil
}

func GetBook(id int) (*domain.Book, error) {
	b, err := Books.GetBook(id)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateBook(b *domain.Book) error {
	err := Books.UpdateBook(b)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(id int) error {
	err := Books.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}
