package main

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func setup() {
	var err error
	conn, err = sql.Open("postgres", "<credentials>")
	if err != nil {
		log.Panic(err)
	}
	conn.SetMaxOpenConns(20) // sane default
	conn.SetMaxIdleConns(0)
	conn.SetConnMaxLifetime(time.Nanosecond)

	if _, err = conn.Exec("create table if not exists harness_tests (id integer)"); err != nil {
		log.Panic(err)
	}
}

func teardown() {
	time.Sleep(time.Millisecond)
	if conn == nil {
		log.Panic("conn should not be nil")
	}
	open := conn.Stats().OpenConnections
	if open > 0 {
		log.Panicf("failed to close %d connections", open)
	}
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown() // Can't use defer because os.Exit :(
	os.Exit(retCode)
}

func TestGood(t *testing.T) {
	for i := 0; i < 10; i++ {
		if err := Good(conn); err != nil {
			t.Error(err)
		}
	}
}

func TestBad(t *testing.T) {
	for i := 0; i < 10; i++ {
		if err := Bad(conn); err != nil {
			t.Error(err)
		}
	}
}
