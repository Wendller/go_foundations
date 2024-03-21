package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := bookHotel(ctx)
	if err != nil {
		fmt.Printf("Failed to book hotel: %v\n", err)
	}
}

func bookHotel(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("hotel booking cancelled: timeout reached")
	case <-time.After(time.Second):
		fmt.Println("Hotel booked")
		return nil
	}
}
