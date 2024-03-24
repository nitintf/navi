package utils

import (
	"log"
	"sync"
	"time"

	"github.com/briandowns/spinner"
)

var (
	spin *spinner.Spinner
	once sync.Once
)

func GetSpinner() *spinner.Spinner {
	once.Do(func() {
		spin = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		err := spin.Color("cyan")
		if err != nil {
			log.Fatal(err)
		}
	})
	return spin
}
