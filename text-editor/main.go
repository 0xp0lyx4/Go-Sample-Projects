package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	color.Yellow("File name >")
	var filename string
	fmt.Scanln(&filename)
	file, err := os.Create(filename)
	if err != nil {
		println(err.Error)
		return
	}
	for {
		fmt.Printf("%s :", hostname)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		_, err := file.WriteString(text)
		if err != nil {
			println(err.Error)
			file.Close()
			return
		}
		color.Green("File is Successfully Updated!")

	}
}
