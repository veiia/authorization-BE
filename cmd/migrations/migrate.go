package main

import (
	"github.com/golang-migrate/migrate"
	"log"
)

func main() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5436/example?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
