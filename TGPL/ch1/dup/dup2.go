package dup

import (
	"bufio"
	"fmt"
	"os"
)

func dup2() {
	counts := make(map[string]int)
	fileMap := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "duq2: %v \n", err)
				continue
			}
			countLines(f, counts, fileMap)
		}
	}
	for line, count := range counts {
		if count > 0 {
			fmt.Printf("%d \t %s \t %v\n", count, line, fileMap[line])
		}
	}

}
func countLines(f *os.File, counts map[string]int, fileMap map[string][]string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fileName := f.Name()
		counts[line]++
		if !contains(fileMap[line], fileName) {
			fileMap[line] = append(fileMap[line], fileName)

		}

	}
}

func contains(s []string, fileName string) bool {
	for _, str := range s {
		if str == fileName {
			return true
		}
	}
	return false
}
