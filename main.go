package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Write_To_File(file_name string, ascii string) {
	file, err := os.Create(file_name)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(ascii)
}
func Get_ascii_char(caractere string) []string {
	output := []string{}
	for _, element := range caractere {
		char := rune(element - 32)
		file, err := os.Open("standard.txt")
		if err != nil {
			os.Exit(1)
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(file)
		ascii := []string{}
		for scanner.Scan() {
			ascii = append(ascii, scanner.Text())
		}
		for i := char * 9; i < (char*9)+9; i++ {
			output = append(output, ascii[i])
		}
	}
	return output
}
func Show_ascii(ascii_char []string) string {
	nbr_lettre := len(ascii_char) / 9
	var word_array []string
	var result string
	for y := 0; y < 9; y++ {
		for i := y; i < nbr_lettre*9; i += 9 {
			word_array = append(word_array, ascii_char[i])
			if ascii_char[i] != "\n" {
				result += ascii_char[i]
			}
		}
		result += "\n"
	}
	return result
}
func main() {
	var ascii_char []string
	var file_name string
	if len(os.Args) > 2 {
		//fmt.Println(os.Args[2][0:9])[2]
		if len(os.Args[2]) > 8 {
			if os.Args[2][0:9] == "--output=" {
				ascii_char = Get_ascii_char(os.Args[1])
				file_name = os.Args[2][9:]
				Write_To_File(file_name, Show_ascii(ascii_char))
			} else {
				return
			}
		}

	}
}
