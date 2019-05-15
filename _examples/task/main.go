package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/liamg/clinch/task"
)

func main() {

	fmt.Println("")

	if err := task.New(
		"build",
		"compiling...",
		func(t *task.Task) error {
			time.Sleep(time.Second * 3)
			return nil
		},
	).Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	for i := 0; i < 5; i++ {

		if err := task.New(
			"test",
			fmt.Sprintf("running test #%d...", i),
			func(t *task.Task) error {
				time.Sleep(time.Second * time.Duration(rand.Intn(3)))
				return nil
			},
		).Run(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}

	if err := task.New(
		"release",
		"packaging...",
		func(t *task.Task) error {
			time.Sleep(time.Second * 3)
			return nil
		},
	).Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	if err := task.New(
		"release",
		"uploading...",
		func(t *task.Task) error {
			time.Sleep(time.Second * 3)
			return nil
		},
	).Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Println("")
}
