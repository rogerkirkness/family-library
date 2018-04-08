package main

// Structs for each table. Functions for each query.

type Book struct {
	Index  float64 `json:"index"`
	Author string  `json:"author"`
	Title  string  `json:"title"`
	Theme  string  `json:"theme"`
}
