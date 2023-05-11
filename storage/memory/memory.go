package memory

import "github.com/Lubwama-Emmanuel/phoneBook_app/models"

type MemoryStorage struct {
	data map[string]models.DataObject
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]models.DataObject)}
}
