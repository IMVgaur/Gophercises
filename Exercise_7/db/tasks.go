package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var taskBucket = []byte("tasks")

type Task struct {
	Key   int
	Value string
}

//db initialization
//input : path of db data file
//return error
func Init(dbPath string) (*bolt.DB, error) {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return db, db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

//AddTask is a metod to add task to the task list
//input : task string
//return error
func AddTask(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		i, _ := b.NextSequence()
		id := int(i)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
}

//GetAllTasks is a method to list tasks in the database
//return task slice and error
func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	return tasks, err
}

//DeleteTask is a method which takes task id and delete the task from database
//input : key integer
//return error
func DeleteTask(k int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(k))
	})
}

//helper function
//input : integer key
//return : binary representation of integer key
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//helper function
//input : Binary key
//return : Integer representation of binary key
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
