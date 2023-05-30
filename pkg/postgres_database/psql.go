package postgres_database

import (
	"database/sql"
	logr "github.com/sirupsen/logrus"
	"log"
)

func NewConnectionBooks() (*sql.DB, error) {
	connStr := "user=postgres password=12345 dbname=books sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	logr.WithFields(logr.Fields{
		"app_level": "database",
	}).Info("Database opened")

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	logr.WithFields(logr.Fields{
		"app_level": "database",
	}).Info("Database connected")

	return db, nil
}

func NewConnectionUsers() (*sql.DB, error) {
	connStr := "user=postgres password=12345 dbname=users sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	logr.WithFields(logr.Fields{
		"app_level": "database",
	}).Info("Database opened")

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	logr.WithFields(logr.Fields{
		"app_level": "database",
	}).Info("Database connected")

	return db, nil
}
