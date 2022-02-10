package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type School_20220210_163618 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &School_20220210_163618{}
	m.Created = "20220210_163618"

	migration.Register("School_20220210_163618", m)
}

// Run the migrations
func (m *School_20220210_163618) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE school(id serial primary key,name TEXT NOT NULL)")
}

// Reverse the migrations
func (m *School_20220210_163618) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE school")
}
