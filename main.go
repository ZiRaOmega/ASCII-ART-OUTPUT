package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func perror() {
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("Usage: go run . " + "\033[32m" + "[STRING] " + "\033[34m" + "[OPTION] " + "\033[35m" + "[FILE]" + "\033[0m")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("EX: go run . something standard --output=<fileName.txt>")
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("Option must be in this list :")
	fmt.Println("{\033[33m standard \033[33m shadow \033[33m thinkertoy\033[0m }")
	fmt.Println("------------------------------------------------------------------------------------------------")
}

func main() {
	var s []string
	if len(os.Args)-1 != 3 {
		perror()
		return
	}
	if strings.Contains(os.Args[3], "--output") && strings.Contains(os.Args[3], "=") {
		os.Args[3] = strings.Replace(os.Args[3], "--output", "", 1)
		os.Args[3] = strings.Replace(os.Args[3], "=", "", 10)
	} else {
		perror()
		return
	}
	file2, err2 := os.Create(os.Args[3])
	if err2 != nil {
		return
	}
	font := os.Args[2]
	if font == "standard" {
		dat, err1 := os.Open("standard.txt")
		reader := bufio.NewScanner(dat)
		for reader.Scan() {
			s = append(s, reader.Text())
		}
		if err1 != nil {
			return
		}
	} else if font == "thinkertoy" {
		dat, err1 := os.Open("thinkertoy.txt")
		reader := bufio.NewScanner(dat)
		for reader.Scan() {
			s = append(s, reader.Text())
		}
		if err1 != nil {
			return
		}
	} else if font == "shadow" {
		dat, err1 := os.Open("shadow.txt")
		reader := bufio.NewScanner(dat)
		for reader.Scan() {
			s = append(s, reader.Text())
		}
		if err1 != nil {
			return
		}
	} else {
		perror()
		return
	}
	var l1 []string
	var l2 []string
	var l3 []string
	var l4 []string
	var l5 []string
	var l6 []string
	var l7 []string
	var l8 []string
	var l9 []string
	pass := 0
	for j := range os.Args[1] {
		if pass == 1 {
			pass--
			continue
		}
		char := os.Args[1][j]
		ascii := int(char)
		stop := (ascii - 32) * 9
		if j < len(os.Args[1])-1 {
			nextascii := int(os.Args[1][j+1])
			if ascii == 92 && nextascii == 110 {
				if j+1 == len(os.Args[1])-1 {
					PrintSlice(l1, file2)
					file2.WriteString("\n")
					PrintSlice(l2, file2)
					file2.WriteString("\n")
					PrintSlice(l3, file2)
					file2.WriteString("\n")
					PrintSlice(l4, file2)
					file2.WriteString("\n")
					PrintSlice(l5, file2)
					file2.WriteString("\n")
					PrintSlice(l6, file2)
					file2.WriteString("\n")
					PrintSlice(l7, file2)
					file2.WriteString("\n")
					PrintSlice(l8, file2)
					file2.WriteString("\n")
					PrintSlice(l9, file2)
					file2.WriteString("\n")
					file2.WriteString("\n")
					return

				} else {
					PrintSlice(l2, file2)
					file2.WriteString("\n")
					PrintSlice(l3, file2)
					file2.WriteString("\n")
					PrintSlice(l4, file2)
					file2.WriteString("\n")
					PrintSlice(l5, file2)
					file2.WriteString("\n")
					PrintSlice(l6, file2)
					file2.WriteString("\n")
					PrintSlice(l7, file2)
					file2.WriteString("\n")
					PrintSlice(l8, file2)
					file2.WriteString("\n")
					PrintSlice(l9, file2)
					file2.WriteString("\n")
					l1 = l1[:0]
					l2 = l2[:0]
					l3 = l3[:0]
					l4 = l4[:0]
					l5 = l5[:0]
					l6 = l6[:0]
					l7 = l7[:0]
					l8 = l8[:0]
					l9 = l9[:0]
					pass++
					continue

				}
			}
		}
		for i := stop; i <= stop+8; i++ {
			if i == stop {
				l1 = append(l1, " ")
			} else if i == stop+1 {
				l2 = append(l2, s[i])
			} else if i == stop+2 {
				l3 = append(l3, s[i])
			} else if i == stop+3 {
				l4 = append(l4, s[i])
			} else if i == stop+4 {
				l5 = append(l5, s[i])
			} else if i == stop+5 {
				l6 = append(l6, s[i])
			} else if i == stop+6 {
				l7 = append(l7, s[i])
			} else if i == stop+7 {
				l8 = append(l8, s[i])
			} else if i == stop+8 {
				l9 = append(l9, s[i])
			}
		}
	}
	PrintSlice(l1, file2)
	file2.WriteString("\n")
	PrintSlice(l2, file2)
	file2.WriteString("\n")
	PrintSlice(l3, file2)
	file2.WriteString("\n")
	PrintSlice(l4, file2)
	file2.WriteString("\n")
	PrintSlice(l5, file2)
	file2.WriteString("\n")
	PrintSlice(l6, file2)
	file2.WriteString("\n")
	PrintSlice(l7, file2)
	file2.WriteString("\n")
	PrintSlice(l8, file2)
	file2.WriteString("\n")
	PrintSlice(l9, file2)
	file2.WriteString("\n")
}

func PrintSlice(s []string, file2 *os.File) {
	for i := range s {
		file2.WriteString(s[i])
	}
}
