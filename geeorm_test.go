package geeorm

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "/home/dinghe/temp/gee.db")
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	e := OpenDB(t)
	defer e.Close()
}
