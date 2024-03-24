package utils

import (
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
		spin.Color("cyan")
	})
	return spin
}
