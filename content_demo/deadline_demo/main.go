package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, fun := context.WithDeadline(context.Background(),
		time.Now().Add(10*time.Second))
	defer fun()

	for   {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("hahhahaa")
		case <-ctx.Done():
			return
		}
	}
}
