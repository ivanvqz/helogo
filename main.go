package main

import (
	"fmt"

	quoteV3 "rsc.io/quote/v3"

	"github.com/ivanvqz/helogo/greet"

)

func main() {
	fmt.Println(greet.Japanese())
	fmt.Println(quoteV3.HelloV3())
	fmt.Println(quoteV3.Concurrency())
}
