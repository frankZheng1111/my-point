package main

import (
	"fmt"
	"time"
)

type SqrtError struct {
	msg string
}

func (se *SqrtError) Error() string {
	return se.msg
}

func sqrt(num int) (result int, err error) {
	invalidNumErr := &SqrtError{"invalidNum"}
	if num < 0 {
		return 0, invalidNumErr
	}
	if num == 1 || num == 0 {
		return num, nil
	}
	head, tail := 0, num
	for {
		time.Sleep(time.Second / 2)
		fmt.Println("head: ", head, ", tail: ", tail)
		if tail-head <= 1 {
			return 0, &SqrtError{"result is not integer"}
		}
		result = (head + tail) / 2
		tmpNum := result * result
		if tmpNum == num {
			break
		} else if tmpNum > num {
			tail = result
		} else {
			head = result
		}
	}
	return result, nil
}

func main() {
	if result, err := sqrt(25); err == nil {
		fmt.Println("get result: ", result)
	} else {
		fmt.Println("get error: ", err)
	}
}
