package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name     string
	Spouse   string
	PrevYear []PrevPerson
}

type PrevPerson struct {
	Name string
	Year int
}

func main() {

	n := OpenFile()

	sN := strings.Split(n, "\n")

	// fmt.Println(len(sN))
	var people []Person
	for _, person := range sN {
		people = append(people, makePerson(person))
	}
}

func makePerson(s string) Person {
	sN := strings.Split(s, ",")
	var person Person

	var prev PrevPerson

	for i := range len(sN) {
		if i == 0 {
			person.Name = sN[i]
			continue
		}
		if i == 1 {
			person.Spouse = sN[i]
			continue
		}

		if i%2 == 0 {
			prev.Name = sN[i]
		}
		if i%2 == 1 {
			year, err := strconv.Atoi(sN[i][:4])
			if err != nil {
				fmt.Println(err)
			}
			prev.Year = year
			person.PrevYear = append(person.PrevYear, prev)
			prev = PrevPerson{}
		}
	}
	fmt.Println(person)
	return person
}

func NewPerson(name string, spouse string, prev []PrevPerson) Person {
	res := Person{
		Name:     name,
		PrevYear: prev,
	}

	if spouse != "" {
		res.Spouse = spouse
	}

	return res
}

func OpenFile() string {
	// open input file
	fi, err := os.Open("LindseyChristmas.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	// // open output file
	// fo, err := os.Create("output.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// // close fo on exit and check for its returned error
	// defer func() {
	// 	if err := fo.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// // make a write buffer
	// w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// // write a chunk
		// if _, err := w.Write(buf[:n]); err != nil {
		// 	panic(err)
		// }
		return string(buf[:n])
	}

	// if err = w.Flush(); err != nil {
	// 	panic(err)
	// }
	return ""
}
