package main

import (
	"fmt"
	"github.com/pokerblow/poker"
)

func main() {
	res := poker.Eval([]string{"As", "Ks", "Qs", "Js", "Ts", "Ad", "Ac"})
	fmt.Println(res.Rank)
	fmt.Println(res.Hand)
	fmt.Println(res.Cards)
}