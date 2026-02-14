package store

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type InMemoryStore struct {
	list []Product
	mu   sync.Mutex
}

func New() *InMemoryStore {
	return &InMemoryStore{
		list: []Product{},
		mu:   sync.Mutex{},
	}
}

func (s *InMemoryStore) Create(record Product) (Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	record.Id = uuid.New().String()
	s.list = append(s.list, record)
	return record, nil
}

func (s *InMemoryStore) GetAll() (Products, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return Products{List: s.list}, nil
}

func (s *InMemoryStore) GetById(id string) (Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, record := range s.list {
		if record.Id == id {
			return record, nil
		}
	}
	return Product{}, errors.New("record not found")
}

func (s *InMemoryStore) UpdateById(id string, modifiedRecord Product) (Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, record := range s.list {
		if record.Id == id {
			s.list[i] = modifiedRecord
			return s.list[i], nil
		}
	}
	return Product{}, errors.New("record not found")
}

func (s *InMemoryStore) DeleteById(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, record := range s.list {
		if record.Id == id {
			s.list = append(s.list[:i], s.list[i+1:]...)
			return nil
		}
	}
	return errors.New("record not found")
}
