package main

import "fmt"

func LogError (info string, error error) {
	if error == nil {
		return
	}
	fmt.Errorf(info + ": %s", error)
}
