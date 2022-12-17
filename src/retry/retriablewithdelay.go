package retry

import (
	"fmt"
	"time"
)

type RetriableWithDelay struct {
	maxAttempts uint
	delayInSec time.Duration
}

func NewRetriableWithDelay(maxAttempts uint, delayInSec uint) Retriable {
	return RetriableWithDelay{maxAttempts: maxAttempts, delayInSec: time.Duration(delayInSec)}
}

func (rd RetriableWithDelay) GetMaxAttempts() uint {
	return rd.maxAttempts
}

func (rd RetriableWithDelay) Retry(action func() bool) {
	var i uint
	for i = 1; i <= rd.maxAttempts; i++ {
		var attemptResult = action()
		if attemptResult {
			fmt.Printf("Attempt number %d: success!\n", i)
			return
		} else {
			fmt.Printf("Attempt number %d: failed\n", i)
			time.Sleep(rd.delayInSec * time.Second)
		}
	}
	fmt.Println("All of the", rd.maxAttempts, "attempts failed")
}