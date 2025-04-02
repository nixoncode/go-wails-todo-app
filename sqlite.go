package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type SQLiteStore struct {
	db *sql.DB
}

// NewSQLiteStore creates a new SQLite store
func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create the items table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY, name TEXT, done INTEGER)")
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &SQLiteStore{db: db}, nil
}
func (s *SQLiteStore) Add(ctx context.Context, item Item) (Item, error) {
	result, err := s.db.ExecContext(ctx, "INSERT INTO items (name, done) VALUES (?, ?)", item.Name, item.IsDone)
	if err != nil {
		return Item{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Item{}, err
	}
	item.ID = int(id)
	return item, nil
}

func (s *SQLiteStore) Get(ctx context.Context, id int) (Item, error) {
	row := s.db.QueryRowContext(ctx, "SELECT id, name, done FROM items WHERE id = ?", id)
	var item Item
	err := row.Scan(&item.ID, &item.Name, &item.IsDone)
	if err != nil {
		if err == sql.ErrNoRows {
			return Item{}, fmt.Errorf("item with ID %d not found", id)
		}
		return Item{}, err
	}
	return item, nil
}

func (s *SQLiteStore) GetAll(ctx context.Context) ([]Item, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, done FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.IsDone); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *SQLiteStore) Update(ctx context.Context, item Item) error {
	_, err := s.db.ExecContext(ctx, "UPDATE items SET name = ?, done = ? WHERE id = ?", item.Name, item.IsDone, item.ID)
	return err
}

func (s *SQLiteStore) Delete(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM items WHERE id = ?", id)
	return err
}

func (s *SQLiteStore) Count(ctx context.Context) (int, error) {
	row := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM items")
	var count int
	err := row.Scan(&count)
	return count, err
}
