// Αυτά είναι τα αποτελέσματα που πήρα. Αποφάσισα να αυτοματοποιήσω κάπως την διαδικασί
// οπότε λογικά μπορείτε να το τρέξετε και εσείς χωρίς πρόβλημα αν θέλετε να το δοκιμάσετε.
// Αν θέλετε να το τρέξετε θα πρέπει να βάλετε τους κώδικες των υπόλοιπων εργασιών του ίδιου
// φακέλου σε σχόλια εκτός από την γραμμή package main. Μου πήρε λίγη ώρα να το βρω γιαυτό σας το λέω.
// 1house.jpg with 1 threads in 520 ms
// 1house.jpg with 2 threads in 229 ms
// 1house.jpg with 4 threads in 141 ms
// 1house.jpg with 8 threads in 127 ms
// 2aerial.jpg with 1 threads in 1124 ms
// 2aerial.jpg with 2 threads in 627 ms
// 2aerial.jpg with 4 threads in 342 ms
// 2aerial.jpg with 8 threads in 298 ms
// 3tiger.jpg with 1 threads in 2986 ms
// 3tiger.jpg with 2 threads in 1584 ms
// 3tiger.jpg with 4 threads in 859 ms
// 3tiger.jpg with 8 threads in 696 ms
// 4food.jpg with 1 threads in 2788 ms
// 4food.jpg with 2 threads in 1446 ms
// 4food.jpg with 4 threads in 936 ms
// 4food.jpg with 8 threads in 698 ms
// 5landscape.jpg with 1 threads in 2292 ms
// 5landscape.jpg with 2 threads in 1171 ms
// 5landscape.jpg with 4 threads in 773 ms
// 5landscape.jpg with 8 threads in 685 ms
// 6berries.jpg with 1 threads in 2994 ms
// 6berries.jpg with 2 threads in 1459 ms
// 6berries.jpg with 4 threads in 821 ms
// 6berries.jpg with 8 threads in 734 ms
// 7lake.jpg with 1 threads in 3422 ms
// 7lake.jpg with 2 threads in 1748 ms
// 7lake.jpg with 4 threads in 1043 ms
// 7lake.jpg with 8 threads in 816 ms

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	images := []string{
		"1house.jpg",
		"2aerial.jpg",
		"3tiger.jpg",
		"4food.jpg",
		"5landscape.jpg",
		"6berries.jpg",
		"7lake.jpg",
	}

	threadCounts := []int{1, 2, 4, 8}

	for _, imgFile := range images {
		for _, threads := range threadCounts {
			processImage(imgFile, threads)
		}
	}
}

func processImage(inputFile string, numThreads int) {
	base := strings.TrimSuffix(inputFile, ".jpg")
	outputFile := fmt.Sprintf("%sGrey_%d.jpg", base, numThreads)

	// Open input image
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening image %s: %v\n", inputFile, err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image %s: %v\n", inputFile, err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	grayImg := image.NewRGBA(image.Rect(0, 0, width, height))

	const (
		redCoeff   = 0.299
		greenCoeff = 0.587
		blueCoeff  = 0.114
	)

	chunkSize := height / numThreads
	var wg sync.WaitGroup

	start := time.Now()

	for w := 0; w < numThreads; w++ {
		startY := w * chunkSize
		endY := startY + chunkSize
		if w == numThreads-1 {
			endY = height
		}

		wg.Add(1)
		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := 0; x < width; x++ {
					r, g, b, _ := img.At(x, y).RGBA()
					gray := uint8(float64(r>>8)*redCoeff + float64(g>>8)*greenCoeff + float64(b>>8)*blueCoeff)
					grayColor := color.RGBA{gray, gray, gray, 255}
					grayImg.Set(x, y, grayColor)
				}
			}
		}(startY, endY)
	}

	wg.Wait()

	elapsed := time.Since(start)

	// Save output
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file %s: %v\n", outputFile, err)
		return
	}
	defer outFile.Close()

	if err := jpeg.Encode(outFile, grayImg, nil); err != nil {
		fmt.Printf("Error saving image %s: %v\n", outputFile, err)
		return
	}

	// Print result
	fmt.Printf("Processed %s with %d threads in %d ms\n", inputFile, numThreads, elapsed.Milliseconds())
}
