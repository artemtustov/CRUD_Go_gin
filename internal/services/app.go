package services

import (
	"CRUD_Go_gin/internal/database"
	"CRUD_Go_gin/internal/domain"
)

func CreateBook(b *domain.Book) error {
	err := database.CreateBook(b)
	if err != nil {
		return err
	}
	return nil
}

func GetBook(id int) (*domain.Book, error) {
	b, err := database.GetBook(id)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateBook(b *domain.Book) error {
	err := database.UpdateBook(b)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(id int) error {
	err := database.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}
