package main

func (a *App) AddTodo(todo Item) error {
	return a.store.Add(a.ctx, todo)
}

func (a *App) GetTodo(id int) (Item, error) {
	return a.store.Get(a.ctx, id)
}

func (a *App) GetAllTodos() ([]Item, error) {
	return a.store.GetAll(a.ctx)
}

func (a *App) UpdateTodo(todo Item) error {
	return a.store.Update(a.ctx, todo)
}

func (a *App) DeleteTodo(id int) error {
	return a.store.Delete(a.ctx, id)
}

func (a *App) CountTodos() (int, error) {
	return a.store.Count(a.ctx)
}
