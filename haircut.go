// Solving Haircut problem from Google Code Jam 2015 Round 1A
//
// https://code.google.com/codejam/contest/4224486/dashboard#s=p1

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Each Barber takes a certain amount of time to cut hair (cutTime).
// barberNumbers should start at 1.
type Barber struct {
	cutTime, barberNumber int
}

// Returns the number of customers who can be serviced by a given time.
//
// A serviced customer is one who has been completed or is actively having their
// hair cut.
func numCustomersServiced(barbers []Barber, minutes int) int {
	// When we start, <number of barbers> amount of customers are being serviced
	sum := len(barbers)
	// Figure out how many people each individual barber could have serviced in
	// this much time.
	for _, b := range barbers {
		sum += minutes / b.cutTime
	}
	return sum
}

func splitString(str string) []int {
	splitString := strings.Fields(str)
	ints := make([]int, len(splitString))
	for i, s := range splitString {
		output, err := strconv.Atoi(s)
		check(err)
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
		// barberNumbers start at 1
		barbers[b] = Barber{barberTime, b + 1}
	}

	return barbers, customerNumber, nil
}

func solve(barbers []Barber, customerNumber int) int {
	// If the customer is one of the first, just return their place in line.
	if customerNumber <= len(barbers) {
		return customerNumber
	}

	// First we find an upper bound - a number of minutes in which we KNOW our
	// customer must be taken in.
	upperBound := 1
	for numCustomersServiced(barbers, upperBound) < customerNumber {
		upperBound *= 2
	}

	// Now we loop until we find (lowerBound, upperBound) pair s.t. we know the
	// upperBound is the time at which our customer would be taken.

	// We know it is the correct time when the lowerBound would not be enough time
	// and the upperBound is enough time. If we continually cut them in half, this
	// will happen when they're one minute apart.
	lowerBound := 0
	for upperBound-lowerBound != 1 {
		// Test a new bound of an int directly in the middle of our two bounds
		newBound := (upperBound + lowerBound) / 2
		if numCustomersServiced(barbers, newBound) < customerNumber {
			lowerBound = newBound
		} else {
			upperBound = newBound
		}
	}

	customerCutTime := upperBound
	log.Printf("Our customer will go at minute %d", customerCutTime)

	numCustomersBeingServiced := numCustomersServiced(barbers, customerCutTime-1)
	log.Printf("When minute %d starts, %d customers are already being serviced",
		customerCutTime, numCustomersBeingServiced)

	// We know how many customers are being serviced when our minute starts.
	// Assign them to barbers in order until we reach our customer number.
	for _, barber := range barbers {
		if math.Mod(float64(customerCutTime), float64(barber.cutTime)) == 0 {
			numCustomersBeingServiced++
			if numCustomersBeingServiced == customerNumber {
				return barber.barberNumber
			}
		}
	}

	panic(errors.New("Couldn't find a matching barber"))
}

func main() {
	filenamePointer := flag.String(
		"input-file", "input-test.txt", "filename to read input")
	flag.Parse()
	file, err := os.Open(*filenamePointer)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// First line is the number of problems
	read_from_file := scanner.Scan()
	if !read_from_file {
		panic("couldn't read input")
	}
	num_problems, err := strconv.Atoi(scanner.Text())
	check(err)
	log.Printf("num_problems: %d\n", num_problems)

	out, err := os.Create("output.txt")
	check(err)
	defer out.Sync()
	for i := 0; i < num_problems; i++ {
		barbers, customerNumber, done := readNextProblem(scanner)
		if done != nil {
			break
		}

		barberNumber := solve(barbers, customerNumber)
		log.Printf("Our customer uses barber number %d", barberNumber)
		// Cases are 1-s based
		out.WriteString(fmt.Sprintf("Case #%d: %d\n", i+1, barberNumber))
	}
}
