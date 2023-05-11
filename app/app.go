package app

import "github.com/Lubwama-Emmanuel/phoneBook_app/models"

type Storage interface {
	Create(data models.DataObject) error
	Read(path string) (models.DataObject, error)
	ReadAll(path string) ([]models.DataObject, error)
	Update(data models.DataObject) error
	Delete(path string) error
}

type App struct {
	storage Storage
}

func NewApp(storage Storage) *App {
	return &App{storage: storage}
}
