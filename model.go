package main

type Item struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
}
