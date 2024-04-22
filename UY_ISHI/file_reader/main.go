/*
2
Bir vaqtning o'zida bir nechta fayllarni o'qiydigan, ularning contentini o'qiydigan natijalarni birlashtiradigan Go dasturini yozing.
Codelarni github ga joylang
*/

package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func readFile(filename string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch <- string(content)
}

func mergeContents(filenames []string) string {
	var wg sync.WaitGroup
	ch := make(chan string, len(filenames))
	defer close(ch)

	for _, filename := range filenames {
		wg.Add(1)
		go readFile("UY_ISHI/file_reader/"+filename, &wg, ch)
	}

	wg.Wait()

	var mergedContent string
	for i := 0; i < len(filenames); i++ {
		mergedContent += <-ch
	}

	return mergedContent
}

func main() {
	filenames := []string{"nodirali.txt", "abdushukur.txt", "holiqbe.txt"}

	mergedContent := mergeContents(filenames)

	fmt.Println("contents :")
	fmt.Println(mergedContent)
}
