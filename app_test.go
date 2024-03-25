package main

import (
	"log"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	err := a.Initialise(DBUser, DBPassword, "test")
	if err != nil {
		log.Fatal("An error occurred while initializing the database")
	}

	m.Run()

}

func createTable() {

}
