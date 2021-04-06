package main

import (
	"fmt"
	"geeorm"
	_ "geeorm/log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string
}

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "/home/dinghe/temp/gee.db")
	defer engine.Close()
	s := engine.NewSession().Model(&User{})
	s.DropTable()   // _, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	s.CreateTable() // _, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	s.CreateTable() // _, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?);", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec sucess, %d affected\n", count)
}
