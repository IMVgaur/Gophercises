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

func Init(dbPath string) (*bolt.DB, error) {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
	return db, err
}

func AddTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	return id, err
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   int(btoui(k)),
				Value: string(v),
			})
		}
		return nil
	})
	return tasks, err
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func btoui(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
