package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calcFuel(m int64) (fuel int64) {
	return m/3 - 2
}

func calcTotalFuel(mass []int64) (fuel int64) {
	for _, m := range mass {
		fuel += calcFuel(m)
	}
	return fuel
}

func calcFuelOnFuel(mass []int64) int64 {

	var (
		fuel     int64
		itemFuel int64
		allFuel  int64
	)

	for _, m := range mass {
		itemFuel = 0

		fuel = calcFuel(m)
		for fuel > 0 {
			fuel = calcFuel(m)
			m = fuel
			if fuel < 0 {
				break
			}
			itemFuel += fuel
		}
		allFuel += itemFuel
	}

	return allFuel
}

func main() {
	// Read file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var inputs []int64
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, num)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Calculate fuel
	fuel := calcTotalFuel(inputs)
	fmt.Printf("Part 1 - Fuel(%d)\n", fuel)

	// Calculate fuel on fuel
	allFuel := calcFuelOnFuel(inputs)
	fmt.Printf("Part 2 - Fuel(%d)\n", allFuel)

}
