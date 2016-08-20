package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

var foodBucketID = []byte("foodBucket")

// CreateFoods create the BoltDB
func CreateFoods() {
	db, err := bolt.Open("foods.db", 0644, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// store some data
	err = db.Update(func(tx *bolt.Tx) error { // .Tx is a transaction object
		foodBucket, err := tx.CreateBucketIfNotExists(foodBucketID)
		if err != nil {
			return err
		}

		key := []byte("pizza")
		value := []byte("10")
		err = foodBucket.Put(key, value)

		key = []byte("burger")
		value = []byte("20")
		err = foodBucket.Put(key, value)

		key = []byte("salat")
		value = []byte("30")
		err = foodBucket.Put(key, value)

		if err != nil {
			return err
		}

		return nil
	})
}

// GetFoods return all stored food options
func GetFoods() []Food {
	// retrieve the data
	var foods []Food
	db, err := bolt.Open("foods.db", 0644, nil)
	// var value []byte

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		foodBucket := tx.Bucket(foodBucketID)

		cursor := foodBucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			name := string(key)
			rank, err := strconv.ParseInt(string(value), 10, 64)
			if err != nil {
				return err
			}
			foods = append(foods, Food{name: name, rank: int(rank)})

			fmt.Printf("name=%s, rank=%v\n", name, rank)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return foods
}
