package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main contains the business logic
func main() {

	// only initially
	CreateFoods()

	// get all foods with rankings from db
	var foods = GetFoods()

	// calculate chances
	var rankings, maxRanking = GetRankingFromFoods(foods)

	// show
	PrintChances(rankings, maxRanking)

	// pick chance
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	pick := random.Intn(maxRanking)

	// select food from pick
	pickedFood := GetFoodFromPick(rankings, pick)

	// show picked food
	fmt.Printf("\n\nYou have to eat %s\n", pickedFood)

	// update db
	UpdateFoods(pickedFood)
}

// GetRankingFromFoods return rankings, minRanking and maxRanking
func GetRankingFromFoods(foods []Food) ([]Ranking, int) {
	var rankings []Ranking
	var minRanking int
	var maxRanking int

	for i := 0; i <= len(foods)-1; i++ {
		food := foods[i]
		maxRanking += food.rank
		rankings = append(rankings, Ranking{food.name, minRanking, maxRanking})
		minRanking += food.rank
	}

	return rankings, maxRanking
}

// PrintChances simply print out chances in percent
func PrintChances(rankings []Ranking, maxRanking int) {
	var ranking Ranking
	var chance float32

	fmt.Println("Chances are: ")

	for i := 0; i <= len(rankings)-1; i++ {
		ranking = rankings[i]
		chance = ((float32(ranking.max) - float32(ranking.min)) / float32(maxRanking))
		chance = chance * 100
		fmt.Printf("\n\t%s for\t %2.2f %s", ranking.name, chance, "%")
	}
}

// GetFoodFromPick find and return food name from pick
func GetFoodFromPick(rankings []Ranking, pick int) string {

	for i := 0; i <= len(rankings)-1; i++ {
		ranking := rankings[i]
		if pick >= ranking.min && pick <= ranking.max {
			return rankings[i].name
		}
	}
	return ""
}
