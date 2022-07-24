package singletonpattern

import (
	"fmt"
	"sync"
)

// type Database struct {
// 	Driver string
// }

// func (db Database) Connect() {
// 	fmt.Printf("Conecting to database through driver: %s\n", db.Driver)
// 	time.Sleep(100 * time.Millisecond)
// }

// var Db *Database

var Once sync.Once

func newSafeDatabaseConnection() *Database {
	Once.Do(
		func() {
			if Db == nil {
				Db = &Database{Driver: "mysqlDriver"}
				Db.Connect()
			}
		})

	return Db
}

func RunThreadSafeExample() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			newSafeDatabaseConnection()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Thread safe singleton: we are always instantiating the same connection")
}
