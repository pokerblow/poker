# Poker

[![CircleCI](https://circleci.com/gh/chehsunliu/poker/tree/master.svg?style=shield&circle-token=abebd63b852ce8ecdcdf3f7e597be743d07402e4)](https://circleci.com/gh/chehsunliu/poker/tree/master) [![GoDoc](https://godoc.org/github.com/chehsunliu/poker?status.svg)](https://godoc.org/github.com/chehsunliu/poker) [![codecov](https://codecov.io/gh/chehsunliu/poker/branch/master/graph/badge.svg)](https://codecov.io/gh/chehsunliu/poker)

Poker is cloned from the Python library [worldveil/deuces](https://github.com/worldveil/deuces).

## Installation

Use `go get` to install Poker:

```sh
$ go get github.com/pokerblow/poker
```

## Usage

Support 5-, 6-, and 7-card evalutions:

```go
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
```

```sh
$ go run ./main.go
1
Straight Flush
[As Ks Qs Js Ts]
```

## Performance

Compared with [notnil/joker](https://github.com/notnil/joker), Poker is 160x faster on 5-card evaluation, and drops to 40x faster on 7-card evaluation.

```sh
go test -bench=. -benchtime 5s
goos: darwin
goarch: amd64
pkg: github.com/chehsunliu/poker
BenchmarkFivePoker-4    	23396181	       253 ns/op
BenchmarkFiveJoker-4    	  141036	     41662 ns/op
BenchmarkSixPoker-4     	 3037298	      1949 ns/op
BenchmarkSixJoker-4     	   28158	    211533 ns/op
BenchmarkSevenPoker-4   	  356448	     16357 ns/op
BenchmarkSevenJoker-4   	    7143	    759394 ns/op
PASS
ok  	github.com/chehsunliu/poker	40.111s
```
