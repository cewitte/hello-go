package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		err := sqrtError{
			lat:  "23 degrees 28 minutes 33.2 seconds South",
			long: "46 degres 49 minutes 56.5 secpnds West",
			err:  errors.New("\nPlease use a positive number"),
		}

		return 0.00, err

	}
	return 42, nil
}
