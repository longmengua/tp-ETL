package common

import "log"

func Async(goFunc func()) {
	go func() {
		defer func() {
			r := recover()
			if r != nil {
				log.Panic()
			}
		}()
		goFunc()
	}()
}
