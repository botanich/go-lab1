package main

import (
	"fmt"
	"math/rand"
	"github.com/botanich/go-lab1/retry"
)

func ret_false() bool {
	return false
}

func rand_bool() bool {
	return rand.Intn(5) == 0  // 0.2 probability of true
}

func main() {
	retriable := retry.NewDefaultRetriable(7)
	retriable.Retry(ret_false)
	retriable.Retry(rand_bool)

	
	fmt.Println("--------------------------------------------")
}