package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("")

	var minRanking int
	var maxRanking int

	// get all foods with rankings from db
	var foods = [3]Food{{"pizza", 5}, {"burger", 8}, {"salat", 16}}

	// calculate chances
	var rankings []Ranking

	for i := 0; i <= len(foods)-1; i++ {
		fmt.Printf("%v\n", foods[i])
		food := foods[i]
		maxRanking += food.rank
		rankings = append(rankings, Ranking{food.name, minRanking, maxRanking})
		minRanking += food.rank
	}

	// pick chance
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	pick := random.Intn(maxRanking)
	fmt.Printf("pick is %v\n", pick)

	// select food from pick
	var pickedFood string
	for i := 0; i <= len(rankings)-1; i++ {
		rank := rankings[i]
		if pick >= rank.min && pick <= rank.max {
			pickedFood = rank.name
		}
	}
	// show picked food
	fmt.Printf("You have to eat %s\n", pickedFood)
	// update db

}

// Food holds name and ranking
type Food struct {
	name string
	rank int
}

// Ranking holds food namen and min / max ranking
type Ranking struct {
	name string
	min  int
	max  int
}

// GetName return name of the food
func (food Food) GetName() string {
	return food.name
}
