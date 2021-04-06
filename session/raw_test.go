package session

import (
	"database/sql"
	"geeorm/dialect"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var TestDB *sql.DB

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "/home/dinghe/temp/gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	d, _ := dialect.GetDialect("sqlite3")
	return New(TestDB, d)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text)").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatalf("except 2, but got %d", count)
	}
}

func TestSession_QueryRows(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text)").Exec()
	row := s.Raw("SELECT COUNT(*) FROM User").QueryRaw()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatalf("failed to query db: %v", err)
	}
}
