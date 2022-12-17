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
	x := rand.Intn(5)
	return x == 0  // 0.2 probability of true
}

func main() {
	retriable := retry.NewDefaultRetriable(7)
	retriable.Retry(ret_false)
	retriable.Retry(rand_bool)

	fmt.Println("--------------------------------------------")

	delayRetriable := retry.NewRetriableWithDelay(7, 1)
	delayRetriable.Retry(ret_false)
	delayRetriable.Retry(rand_bool)
	fmt.Println("--------------------------------------------")
}