package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	circles := []Circle{
		{ID: 1, Rayon: 2, Latence: time.Second},
		{ID: 2, Rayon: 4, Latence: time.Second * 5},
		{ID: 3, Rayon: 6, Latence: time.Second * 5},
		{ID: 4, Rayon: 8, Latence: time.Second * 5},
		{ID: 5, Rayon: 10, Latence: time.Second * 10},
	}
	results, err := Handler(context.Background(), circles)
	if err != nil {
		fmt.Println("error occured : ", err)
		return
	}

	for _, result := range results {
		fmt.Printf("Perimeter = %f\n", result.Perimeter)
	}
}

type Circle struct {
	ID      int
	Rayon   float64
	Latence time.Duration // Simule une latence
}

type Result struct {
	Perimeter float64
}

func Handler(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))

	g, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		i := i
		circle := circle
		g.Go(func() error {

			ctx, cancel := context.WithCancel(ctx)
			defer cancel() // releases resources if the following code completes before timeout elapses
			result, err := ExtrernalServiceCalculatorWithCancellation(ctx, circle)

			//result, err := ExtrernalServiceCalculator(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

func ExtrernalServiceCalculator(ctx context.Context, circle Circle) (Result, error) {

	time.Sleep(circle.Latence)
	diameter := 2 * circle.Rayon
	result := Result{Perimeter: math.Pi * diameter}

	// simule une erreur dès le premier traitement
	//if circle.ID == 1 {
	//	return Result{}, errors.New(fmt.Sprintf("error on circleID %d", circle.ID))
	//}

	return result, nil
}

func ExtrernalServiceCalculatorWithCancellation(ctx context.Context, circle Circle) (Result, error) {

	select {
	case <-ctx.Done():
		return Result{}, ctx.Err()

	case <-time.After(circle.Latence):
		diameter := 2 * circle.Rayon
		result := Result{Perimeter: math.Pi * diameter}

		// simule une erreur dès le premier traitement
		if circle.ID == 1 {
			return Result{}, errors.New(fmt.Sprintf("error on circleID %d", circle.ID))
		}

		return result, nil
	}

}
