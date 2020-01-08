// Package mathOps performs math operations
package mathOps

// average returns the average of a slice of integers
func average(x ...int) int {
	a := len(x)
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum / a
}

// func main() {

// 	fmt.Println("The average of 2, 4 and 6 is...", average(2, 4, 6))
// 	fmt.Println("The average of 3, 5 is...", average(3, 5))
// 	fmt.Println("The average of 12, 24, 42, 54 and 36 is...", average(12, 24, 42, 54, 36))
// }
