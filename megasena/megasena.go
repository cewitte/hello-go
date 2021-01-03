package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 4 quadrants in the megasena game
var q1 = []int{1, 2, 3, 4, 5, 11, 12, 13, 14, 15, 21, 22, 23, 24, 25}
var q2 = []int{6, 7, 8, 9, 10, 16, 17, 18, 19, 20, 26, 27, 28, 29, 30}
var q3 = []int{31, 32, 33, 34, 35, 41, 42, 43, 44, 45, 51, 52, 53, 54, 55}
var q4 = []int{36, 37, 38, 39, 40, 46, 47, 48, 49, 50, 56, 57, 58, 59, 60}

func drawNumbers(dozens int) (*[]int, error) {
	raffle := make([]int, dozens)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	raffle[0] = q1[r1.Intn(len(q1)-1)]
	raffle[1] = q2[r1.Intn(len(q2)-1)]
	raffle[2] = q3[r1.Intn(len(q3)-1)]
	raffle[3] = q4[r1.Intn(len(q4)-1)]

	// now pick quadrants to draw the remaining numbers
	for i := 4; i < dozens; i++ {
		switch r1.Intn(3) {
		case 0:
			raffle[i] = q1[r1.Intn(len(q1)-1)]
		case 1:
			raffle[i] = q2[r1.Intn(len(q2)-1)]
		case 2:
			raffle[i] = q3[r1.Intn(len(q3)-1)]
		case 3:
			raffle[i] = q4[r1.Intn(len(q4)-1)]
		}
	}

	sort.Ints(raffle)

	err := checkRepeated(&raffle)
	if err != nil {
		return nil, err
	}

	err = checkComeOutLittle(&raffle)
	if err != nil {
		return nil, err
	}

	err = checkSequencedNumbers(&raffle)
	if err != nil {
		return nil, err
	}

	err = balanceEvenAndOdd(&raffle)
	if err != nil {
		return nil, err
	}

	err = checkVerticalSequence(&raffle)
	if err != nil {
		return nil, err
	}

	return &raffle, nil
}

func checkRepeated(raffle *[]int) error {
	for k1, v1 := range *raffle {
		for k2, v2 := range *raffle {
			if k1 != k2 && v1 == v2 {
				return fmt.Errorf("%v and %v are repeated numbers", v1, v2)
			}
		}
	}

	return nil
}

func checkComeOutLittle(raffle *[]int) error {
	toAvoid := []int{1, 2, 3, 11, 22, 44, 55, 48, 57}

	for _, drawn := range *raffle {
		for _, avoided := range toAvoid {
			if drawn == avoided {
				return fmt.Errorf("%v is a bad bet", drawn)
			}
		}
	}

	return nil
}

func checkSequencedNumbers(raffle *[]int) error {
	for i, v := range *raffle {
		if i < len(*raffle)-1 {
			n := i + 1
			if (*raffle)[n]-v == 1 {
				return fmt.Errorf("there is a sequence in %v and %v", (*raffle)[i], (*raffle)[n])
			}
		}
	}

	return nil
}

func balanceEvenAndOdd(raffle *[]int) error {
	evenNums := 0

	for _, v := range *raffle {
		if v%2 == 0 {
			evenNums++
		}
	}

	if evenNums < 2 || evenNums > 4 {
		return fmt.Errorf("even and odds are not balanced with %v even numbers", evenNums)
	}

	return nil
}

func checkVerticalSequence(raffle *[]int) error {
	for i, v := range *raffle {
		if i < len(*raffle)-1 {
			n := i + 1
			if (*raffle)[n]-v == 10 {
				return fmt.Errorf("there is a vertical sequence in %v and %v", (*raffle)[i], (*raffle)[n])
			}
		}
	}

	return nil
}

func main() {
	num, err := drawNumbers(6)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*num)

}
