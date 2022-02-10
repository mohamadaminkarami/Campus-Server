package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Student_20220210_170144 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Student_20220210_170144{}
	m.Created = "20220210_170144"

	migration.Register("Student_20220210_170144", m)
}

// Run the migrations
func (m *Student_20220210_170144) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE student(id serial primary key,student_number TEXT NOT NULL,password TEXT NOT NULL,email TEXT NOT NULL,entrance_year integer DEFAULT NULL,rand integer DEFAULT NULL)")
}

// Reverse the migrations
func (m *Student_20220210_170144) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE student")
}
