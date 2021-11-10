package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1]

	if len(os.Args) != 4 || (os.Args[2] != "standard" && os.Args[2] != "shadow" && os.Args[2] != "thinkertoy") {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	} else {

		file, err := os.Open(os.Args[2] + ".txt")
		if err != nil {
			fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
			fmt.Println("EX: go run . something standard --output=<fileName.txt>")
			os.Exit(0)
		}
		defer file.Close()

		scanned := bufio.NewScanner(file) // reading file
		scanned.Split(bufio.ScanLines)

		var lines []string

		for scanned.Scan() {
			lines = append(lines, scanned.Text())
		}

		file.Close()

		asciiChrs := make(map[int][]string)
		id := 31

		for _, line := range lines {
			if string(line) == "" {
				id++
			} else {
				asciiChrs[id] = append(asciiChrs[id], line)
			}
		}

		for i := 0; i < len(args); i++ {
			if args[i] == 92 && args[i+1] == 110 {
				a := Newline(string(args[:i]), asciiChrs)

				b := os.WriteFile(os.Args[3][9:], []byte(a), 0644)
				if b != nil {
					panic(b)
				}
				c := Newline(string(args[i+2:]), asciiChrs)
				f, _ := os.OpenFile(os.Args[3][9:], os.O_APPEND|os.O_WRONLY, 0644)
				n, err := f.WriteString(c)
				if err != nil {
					panic(n)
				}
			}
		}

		// checking for new line within arguments
		if !strings.Contains(args, "\\n") {
			z := Newline(args, asciiChrs)
			k := os.WriteFile(os.Args[3][9:], []byte(z), 0644)
			if k != nil {
				panic(k)
			}
		}

	}

}
func Newline(n string, y map[int][]string) string {
	a := []string{}
	// prints horizontally
	for j := 0; j < len(y[32]); j++ {
		for _, letter := range n {
			a = append(a, y[int(letter)][j])
		}
		a = append(a, "\n")
	}
	b := strings.Join(a, "")
	return b
}
