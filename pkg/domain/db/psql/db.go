package psql

import (
	"access-management/pkg/config"
	"fmt"
)

type DB struct {
	*config.Config
}

func NewDB(c *config.Config) *DB {
	return &DB{}
}

func (db *DB) Print() {
	fmt.Println("Success")
}
