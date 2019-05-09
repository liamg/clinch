package main

import (
	"time"

	"github.com/liamg/clinch/loading"
)

func main() {

	messages := []string{
		"Ruining...",
		"Running: rm -rf /",
		"Wasting time...",
		"Executing pointless task...",
	}

	bar := loading.NewBar()
	defer bar.Close()
	for i := 0.0; i <= 100; i++ {
		bar.Log(messages[int(i)%len(messages)])
		bar.SetPercent(i)
		time.Sleep(time.Millisecond * 100)
	}

}
