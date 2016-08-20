package main

import (
	"bytes"
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

// UpdateFoods increase the ranking of the picked food
func UpdateFoods(pickedFood string, newRanking int) {

	fmt.Printf("Looking for %v\n", pickedFood)

	db, err := bolt.Open("foods.db", 0644, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Println("alright")

	db.Update(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		bucket := tx.Bucket(foodBucketID)
		pick := []byte(pickedFood)
		cursor := bucket.Cursor()

		fmt.Printf("Looking for %v\n", pick)

		for key, value := cursor.Seek(pick); bytes.Equal(key, pick); key, value = cursor.Next() {
			fmt.Printf("key=%s, value=%s\n", key, value)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(foodBucketID)
		pick := []byte(pickedFood)

		err := bucket.Put(pick, []byte(string(newRanking)))
		return err
	})

	return
}
