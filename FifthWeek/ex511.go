// Εδώ είναι ο πολ/μος πινάκων. Ο κώδικας χωρίς πραλληλισμό σε Java δεν έτρεξε σε 1 
// λεπτό, που είχα την  υπομονή να περιμένω, με παραλληλισμό εδώ τρέχει σε περίπου 1 sec (τοπικά)
// και 3 με 4 sec στο playground, αλλά φαντάζομαι έιναι και το δίκτυο που παίζει ρόλο.
// Είδα πως δεν χρειάζεται να το δοκιμάσουμε απλά το έκανα απο περιέργια και είπα να το αναφέρω.
// Έτρεξα και για μιρότερα sizes αν σας ενδιαφέρει.

package main

import (
	"fmt"
	"sync"
)

func main() {
	size := 1000

	a := make([][]float64, size)
	b := make([][]float64, size)
	c := make([][]float64, size)

	for i := 0; i < size; i++ {
		a[i] = make([]float64, size)
		b[i] = make([]float64, size)
		c[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			a[i][j] = 0.1
			b[i][j] = 0.3
			c[i][j] = 0.5
		}
	}

	var wg sync.WaitGroup

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			for j := 0; j < size; j++ {
				a[row][j] = b[row][j] + c[row][j]
			}
		}(i)
	}

	wg.Wait()

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%.1f ", a[i][j])
		}
		fmt.Println()
	}
}