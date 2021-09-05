package models

import "fmt"

type DB struct {
	*Config
}

func NewDB(c *Config) *DB {
	return &DB{}
}

func (db *DB) Print()  {
	fmt.Println("Success")
}
