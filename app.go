package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx   context.Context
	store Store
}

// NewApp creates a new App application struct
func NewApp() *App {
	//store := NewInMemoryStore()
	store, err := NewSQLiteStore("todo.db")
	if err != nil {
		fmt.Printf("Failed to create SQLite store: %v\n", err)
		panic("Failed to create SQLite store")
	}

	// check if the store is nil
	if store == nil {
		panic("Store is nil")
	}

	return &App{
		store: store,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	fmt.Println("App is starting up...")
}

func (a *App) shutdown(ctx context.Context) {
	// Perform any cleanup tasks here
	fmt.Println("App is shutting down...")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// sum two numbers
func (a *App) Sum(one, two int) int {
	return one + two
}
