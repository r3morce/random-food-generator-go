package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var minRanking int
	var maxRanking int

	// get all foods with rankings from db
	var foods = [3]Food{{"pizza", 10}, {"burger", 20}, {"salat", 30}}

	// calculate chances
	var rankings []Ranking

	for i := 0; i <= len(foods)-1; i++ {
		food := foods[i]
		maxRanking += food.rank
		rankings = append(rankings, Ranking{food.name, minRanking, maxRanking})
		minRanking += food.rank
	}

	// show chances
	var ranking Ranking
	var chance float32
	fmt.Println("Chances are: ")
	for i := 0; i <= len(rankings)-1; i++ {
		ranking = rankings[i]
		chance = ((float32(ranking.max) - float32(ranking.min)) / float32(maxRanking))
		chance = chance * 100
		fmt.Printf("\n\t%s for\t %2.2f %s", ranking.name, chance, "%")
	}

	// pick chance
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	pick := random.Intn(maxRanking)

	// select food from pick
	var pickedFood string
	for i := 0; i <= len(rankings)-1; i++ {
		ranking := rankings[i]
		if pick >= ranking.min && pick <= ranking.max {
			pickedFood = ranking.name
		}
	}

	// show picked food
	fmt.Printf("\nYou have to eat %s\n", pickedFood)

	// update db
	for i := 0; i <= len(foods)-1; i++ {
		food := foods[i]
		if food.name == pickedFood {
			foods[i].rank++
		}
	}
	// fmt.Println(foods)
}

// Food holds name and current ranking
type Food struct {
	name string
	rank int
}

// Ranking holds food name and rating range for pick
type Ranking struct {
	name string
	min  int
	max  int
}
