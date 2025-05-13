// // Αυτομαροποίησα τα τεστ για να δούμε το "μεγάλωμα" του αρχείου και την ταχύτητα του προγράμματος.
// // Χρησιμοποίησα το αρχείο bible.txt που προτείνετε και το έκανα 4, 6, 8, 10, 12, 14, 16, 18 και 20 φορές μεγαλύτερο.
// // Το πρόγραμμα τρέχει με 8 goroutines και μετράει τις λέξεις του κειμένου.
// // Γενικά το πρόγραμμα είναι γρήγορο και η ταχύτητα του δεν επηρεάζεται από το μέγεθος του αρχείου ιδιαίτερα.
// // Για την ακρίβεια παρατηρώ μια γραμμική, αναμενώμενη, αύξηση.

package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// 	"sync"
// 	"time"
// )

// func mapper(lines []string, out chan<- map[string]int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	wordCount := make(map[string]int)
// 	for _, line := range lines {
// 		words := strings.Fields(line)
// 		for _, word := range words {
// 			word = strings.ToLower(strings.Trim(word, ".,!?:;\"'()[]{}"))
// 			wordCount[word]++
// 		}
// 	}
// 	out <- wordCount
// }

// func reducer(in <-chan map[string]int, out map[string]int, done *sync.WaitGroup) {
// 	defer done.Done()
// 	for wc := range in {
// 		for word, count := range wc {
// 			out[word] += count
// 		}
// 	}
// }

// func main() {
// 	file, err := os.ReadFile("bible.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	for multiplier := 4; multiplier <= 20; multiplier += 2 {
// 		fmt.Printf("Running with %dx file size:\n", multiplier)

// 		content := string(file)
// 		var multiplied string
// 		for i := 0; i < multiplier; i++ {
// 			multiplied += content
// 		}
// 		lines := strings.Split(multiplied, "\n")

// 		start := time.Now()

// 		numWorkers := 8
// 		chunkSize := len(lines) / numWorkers
// 		mapChannel := make(chan map[string]int, numWorkers)

// 		var mapWG sync.WaitGroup
// 		for i := 0; i < numWorkers; i++ {
// 			start := i * chunkSize
// 			end := start + chunkSize
// 			if i == numWorkers-1 {
// 				end = len(lines)
// 			}
// 			mapWG.Add(1)
// 			go mapper(lines[start:end], mapChannel, &mapWG)
// 		}

// 		go func() {
// 			mapWG.Wait()
// 			close(mapChannel)
// 		}()

// 		finalCounts := make(map[string]int)
// 		var reduceWG sync.WaitGroup
// 		reduceWG.Add(1)
// 		go reducer(mapChannel, finalCounts, &reduceWG)

// 		reduceWG.Wait()
// 		elapsed := time.Since(start)

// 		fmt.Printf("Execution Time: %s\n\n", elapsed)
// 	}
// }
