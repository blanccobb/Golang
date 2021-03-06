package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
)

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	
	err := ReadFrom(r, func(line string) {
			lines = append(lines,line)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
}
