package utils

import (
	"fmt"
	"time"
)

func Loading() (stopLoading func()) {
	ch := make(chan struct{})

	pointsPrinted := 0
	fmt.Print("Carregando")
	go func() {
		for {
			select {
			case <-ch:
				return
			default:
				if pointsPrinted == 4 {
					fmt.Print("\b \b")
					fmt.Print("\b \b")
					fmt.Print("\b \b")
					pointsPrinted = 1
				} else {
					fmt.Print(".")
					pointsPrinted++
				}
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	return func() {
		ch <- struct{}{}
		close(ch)
	}
}
