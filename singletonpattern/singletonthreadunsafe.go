package singletonpattern

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {
	Driver string
}

func (db Database) Connect() {
	fmt.Printf("Conecting to database through driver: %s\n", db.Driver)
	time.Sleep(100 * time.Millisecond)
}

var Db *Database

func newDatabaseConnection() *Database {
	if Db == nil {
		Db = &Database{Driver: "mysqlDriver"}
		Db.Connect()
	} else {
		fmt.Println("Using previously created database connection")
	}

	return Db
}

func RunThreadUnsafeExample() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			newDatabaseConnection()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Thread unsafe singleton: Running this programs multiple times will show multiple connections")
}
