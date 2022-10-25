package main

import (
	"bufio"
	"fmt"
	"io"
)

func main() {
	fmt.Println("Go!")
}

func countEmptyLine(reader io.Reader) (int, error) {
	// delégué au client
	//file, err := os.Open(filename)
	//if err != nil {
	//	return 0, err
	//}
	//defer file.Close()

	count := 0
	sc := bufio.NewScanner(reader) //by default, split the input per line
	for sc.Scan() {

		if sc.Text() == "" {
			//if len(sc.Text()) == 0 {
			//if len(sc.Bytes()) == 0 {
			//if utf8.RuneCountInString(sc.Text()) == 0 {
			count += 1
		}
	}
	return count, nil
}
