package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
)

type format struct {
	name   string
	width  int
	height int
}

var formats = map[string]format{
	"A10":     {"A10", 26, 37},
	"A9":      {"A9", 37, 52},
	"A8":      {"A8", 52, 74},
	"A7":      {"A7", 74, 105},
	"A6":      {"A6", 105, 148},
	"A5":      {"A5", 148, 210},
	"A4":      {"A4", 210, 297},
	"A3":      {"A3", 297, 420},
	"A2":      {"A2", 420, 594},
	"A1":      {"A1", 594, 841},
	"A0":      {"A0", 841, 1189},
	"Epson 1": {"Epson 1", 1580, 1580},
	"Epson 2": {"Epson 2", 1580, 1580 * 2},
	"Epson 3": {"Epson 3", 1580, 1580 * 3},
	"Epson 4": {"Epson 4", 1580, 1580 * 4},
	"Epson 5": {"Epson 5", 1580, 1580 * 5},
	"Epson 6": {"Epson 4", 1580, 1580 * 6},
}

const (
	maxWidth  = 1580
	maxHeight = 1580 * 6
)

func closestFormat(width, height *int) (*format, error) {
	if *width > maxWidth || *height > maxHeight {
		return nil, errors.New("Larger than largest allowed format")
	}
	var w, h int
	if *width > *height {
		// It is Landscape. Let's make it Portrait so we can compare apples to apples.
		w = *height
		h = *width
	} else {
		w = *width
		h = *height
	}

	diffs := make(map[int]string)
	vals := []int{}
	for k, v := range formats {
		if v.width >= w && v.height >= h {
			val := (v.height * v.width) - (h * w)
			if _, ok := diffs[val]; !ok {
				diffs[val] = k
				vals = append(vals, val)
			}
		}
	}

	sort.Ints(vals)
	str := diffs[vals[0]]
	f := formats[str]
	return &f, nil
}

func (f *format) String() string {
	return fmt.Sprintf("Format=%s\tWidth=%d\tHeight=%d", f.name, f.width, f.height)
}

func printResult(width, height int) {
	v, err := closestFormat(&width, &height)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("The closest format to width=%d and height=%d is %s\n", width, height, v)
}

func main() {
	w := flag.Int("w", 210, "-w=210")
	h := flag.Int("h", 297, "-h=297")
	count := flag.Int("c", 100, "-c=100")
	flag.Parse()
	printResult(*w, *h)

	fmt.Println("Testing all formats now...")
	for _, v := range formats {
		printResult(v.width, v.height)
	}

	fmt.Println("\n\nTesting ramdomly...")
	for i := 0; i < *count; i++ {
		w := rand.Intn(maxWidth)
		h := rand.Intn(maxHeight)
		if w > 0 && h > 0 {
			if w > h {
				// It's landscape. Let's rotate it.
				printResult(h, w)
			} else {
				printResult(w, h)
			}

		}
	}
}
