// To wait for multiple goroutines to finish, we can
// use a *wait group*.

package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	caseFunc := func() (bool, error) {
		return false, nil
	}
	err := PollImmediate(1*time.Second, 2*time.Second, caseFunc)
	fmt.Println("polling in action: (never hits success)", err)

	caseFunc2 := func() (bool, error) {
		return true, nil
	}
	err = PollImmediate(1*time.Second, 2*time.Second, caseFunc2)
	fmt.Println("polling in action: for instant sucess", err)

	caseFunc3 := func() (bool, error) {
		return false, fmt.Errorf("error case")
	}
	err = PollImmediate(1*time.Second, 2*time.Second, caseFunc3)
	fmt.Println("polling in action: for instant error", err)
}

var ErrWaitTimeout = errors.New("timed out waiting for the condition")

func PollImmediate(interval time.Duration, timeout time.Duration, condition func() (bool, error)) error {

	var totalTime time.Duration = 0

	for totalTime < timeout {
		cond, err := condition()
		if cond {
			return nil
		}
		if err != nil {
			return err
		}
		time.Sleep(interval)
		totalTime += interval
	}

	return ErrWaitTimeout
}
