package main

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
