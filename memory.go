package main

import (
	"context"
	"fmt"
	"sync"
)

type InMemoryStore struct {
	items  map[int]Item
	nextID int
	mu     sync.Mutex
}

// NewInMemoryStore creates a new in-memory store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		items:  make(map[int]Item),
		nextID: 1,
	}
}

// Add adds an item to the store
func (s *InMemoryStore) Add(ctx context.Context, item Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	item.ID = s.nextID
	s.items[s.nextID] = item
	s.nextID++

	return nil
}

// Get retrieves an item by ID
func (s *InMemoryStore) Get(ctx context.Context, id int) (Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[id]
	if !exists {
		return Item{}, fmt.Errorf("item with ID %d not found", id)
	}

	return item, nil
}

// GetAll retrieves all items from the store
func (s *InMemoryStore) GetAll(ctx context.Context) ([]Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	items := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}

	return items, nil
}

// Update updates an existing item in the store
func (s *InMemoryStore) Update(ctx context.Context, item Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.items[item.ID]
	if !exists {
		return fmt.Errorf("item with ID %d not found", item.ID)
	}

	s.items[item.ID] = item
	return nil
}

// Delete removes an item from the store by ID
func (s *InMemoryStore) Delete(ctx context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.items[id]
	if !exists {
		return fmt.Errorf("item with ID %d not found", id)
	}

	delete(s.items, id)
	return nil
}

// Count returns the number of items in the store
func (s *InMemoryStore) Count(ctx context.Context) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.items), nil
}
