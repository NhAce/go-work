package main

import (
	"fmt"
)

type Latlon struct {
	Lat, Long float64
}

var m map[string]Latlon
var n = map[string]Latlon{
	"Bell Labs": Latlon{40.312354, 120.336780},
	"Google": Latlon{37.422323, 122.986723},
}
func main() {
	m = make(map[string]Latlon)
	m["Bell Labs"] = Latlon{40.312354, 120.336780}
	fmt.Println(m["Bell Labs"])

	fmt.Println(n)
}

