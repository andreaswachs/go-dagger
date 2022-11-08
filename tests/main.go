package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	args := os.Args

	for i, arg := range args {
		if i != 0 {
			fmt.Printf("Got arg %d: ", i)
			fmt.Println(arg)

		}
	}

	if err := test(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func test(ctx context.Context) error {
	fmt.Println("Testing with Dagger")

	// initialize Dagger client
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	return nil
}
