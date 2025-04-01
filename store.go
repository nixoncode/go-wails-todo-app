package main

import "context"

type Store interface {
	Add(ctx context.Context, item Item) error
	Get(ctx context.Context, id int) (Item, error)
	GetAll(ctx context.Context) ([]Item, error)
	Update(ctx context.Context, item Item) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int, error)
}
