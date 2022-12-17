package retry

import (
	"fmt"
)

type DefaultRetriable struct {
	maxAttempts uint
}

func NewDefaultRetriable(maxAttempts uint) Retriable {
	return DefaultRetriable{maxAttempts: maxAttempts}
}

func (dr DefaultRetriable) GetMaxAttempts() uint {
	return dr.maxAttempts
}

func (dr DefaultRetriable) Retry(action func() bool) {
	var i uint
	for i = 1; i <= dr.maxAttempts; i++ {
		var attemptResult = action()
		if attemptResult {
			fmt.Printf("Attempt number %d: success!\n", i)
			return
		} else {
			fmt.Printf("Attempt number %d: failed\n", i)
		}
	}
	fmt.Println("All of the", dr.maxAttempts, "attempts failed")
}