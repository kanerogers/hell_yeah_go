package common

type Status struct {
	Status string
}

type Listing struct {
	ID   string `json:"listing_id"`
	Name string `json:"name"`
}

type User struct {
	ID   string `json:"user_id"`
	Name string `json:"name"`
}

var JWTKEY = "heyheyhey"
