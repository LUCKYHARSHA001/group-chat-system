package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Task finished successfully!")
	case <-ctx.Done():
		fmt.Println("Task canceled because it took too long:", ctx.Err())
	}
}