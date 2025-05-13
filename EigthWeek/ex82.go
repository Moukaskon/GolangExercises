// Μετά απο διάφορα τεστ που έκανα, threshold = 5000000, 5000, 50000, 500000 παρατήρησα πως
// η ταχύτητα εκτέλεσης του προγράμματος δεν αλλάζει σημαντικά. Ίσως το Serial mergesort
// να είναι λίγο πιο αργό όσο ανεβαίνει υο threshold, αλλά δεν είναι κάτι το σημαντικό.
// Κατι που όντως έκανε διαφορά είναι να χρησιμοποιήσω το runtime.GOMAXPROCS(runtime.NumCPU())
// για να εκμεταλλευτώ όλους τους πυρήνες του υπολογιστή μου. Με τις μισές CPU ήταν όντως πιο αργό το πρόγραμμα.

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)



func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}


func parallelMergeSort(arr []int, wg *sync.WaitGroup, threshold int) []int {
	if wg != nil {
		defer wg.Done()
	}

	if len(arr) <= 1 {
		return arr
	}

	if len(arr) <= threshold {
		return mergeSort(arr) 
	}

	mid := len(arr) / 2
	var left, right []int
	var innerWg sync.WaitGroup

	innerWg.Add(2)
	go func() {
		left = parallelMergeSort(arr[:mid], &innerWg, threshold)
	}()
	go func() {
		right = parallelMergeSort(arr[mid:], &innerWg, threshold)
	}()
	innerWg.Wait()

	return merge(left, right)
}


func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func generateRandomSlice(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size)
	}
	return arr
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
 	threshold := 5000000

	runtime.GOMAXPROCS(runtime.NumCPU())

	size := 10_000_000
	data := generateRandomSlice(size)

	copy1 := make([]int, size)
	copy(copy1, data)

	start := time.Now()
	sorted1 := mergeSort(copy1)
	serialTime := time.Since(start)
	fmt.Printf("Serial MergeSort time: %v\n", serialTime)

	copy2 := make([]int, size)
	copy(copy2, data)

	var wg sync.WaitGroup
	wg.Add(1)
	start = time.Now()
	sorted2 := parallelMergeSort(copy2, &wg, threshold)
	wg.Wait()
	parallelTime := time.Since(start)
	fmt.Printf("Parallel MergeSort time: %v\n", parallelTime)

	if slicesEqual(sorted1, sorted2) {
		fmt.Println("Sorting verified. Arrays are equal.")
	} else {
		fmt.Println("Error: Sorted arrays do not match.")
	}
}
