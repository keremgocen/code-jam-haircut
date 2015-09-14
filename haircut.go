// Solving Haircut problem from Google Code Jam 2015 Round 1A
//
// https://code.google.com/codejam/contest/4224486/dashboard#s=p1

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Barber struct {
	cut_time, barber_number int
}

// Returns the number of customers who can be completed in the given time.
func numCustomersCompleted(barbers []Barber, minutes int) int {
	sum := 0
	for _, b := range barbers {
		sum += minutes / b.cut_time
	}
	return sum
}

func splitString(str string) []int {
	splitString := strings.Fields(str)
	ints := make([]int, len(splitString))
	for i, s := range splitString {
		output, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = output
	}
	return ints
}

// Returns the relevant information for the next problem
//
// Each problem is described on two lines:
//  1. First line includes the number of barbers and the customer position
//  2. Second line has the cutting times for each barber
//
// Returns a tuple of (<all the barbers in a []Barber list, the customerNumber).
// These have all be converted into ints (instead of strings).
//
// Returns error when EOF is reached (ie: no more problems).
func readNextProblem(scanner *bufio.Scanner) ([]Barber, int, error) {
	if !scanner.Scan() {
		return nil, 0, errors.New("All done!")
	}

	firstLine := splitString(scanner.Text())
	numBarbers := firstLine[0]
	customerNumber := firstLine[1]

	if !scanner.Scan() {
		panic("EOF at unexpected time")
	}
	barberTimes := splitString(scanner.Text())

	if len(barberTimes) != numBarbers {
		log.Panicf("Found %d barbers, expected to find %d\n",
			len(barberTimes), numBarbers)
	}

	barbers := make([]Barber, numBarbers)
	for b, barberTime := range barberTimes {
		barbers[b] = Barber{barberTime, b}
	}

	return barbers, customerNumber, nil
}

func main() {
	file, err := os.Open("input-test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// First line is the number of problems
	read_from_file := scanner.Scan()
	if !read_from_file {
		panic("couldn't read input")
	}
	num_problems, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	fmt.Printf("num_problems: %d\n", num_problems)

	for i := 0; i < num_problems; i++ {
		barbers, customerNumber, done := readNextProblem(scanner)
		if done != nil {
			break
		}
		log.Printf("Problem %d. barbers: %v, customerNumber: %d\n",
			i, barbers, customerNumber)
	}
}
