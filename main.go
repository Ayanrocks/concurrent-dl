package main

import (
	"fmt"
	"os"

	"github.com/kyokomi/emoji"
)

var Urls = []string{}

func main() {

	// wait for user input
	getFile()
	// ask if more links are there

	// fetch the url concurently

	// store in the file system

	// show the progress in the terminal
}

func downloadFile() {

}

func getFile() {
	emoji.Println("Welcome to Golang Concurrent Downloader :smiley:")
	emoji.Println("Please add links to the files to download. (-1 to confirm :heavy_check_mark: 0 to exit :heavy_multiplication_x:)")
	url := ""
	i := 1
	for {

		emoji.Printf("Enter url %d :link: to download: ", i)
		url = ""
		fmt.Scanln(&url)

		if url == "-1" {
			break
		} else if url == "0" {
			os.Exit(0)
		} else if url == "" || url == "\n" {
			emoji.Printf("Please enter the url %d :rage: to download: ", i)
			fmt.Scanln(&url)
			if url == "" {
				os.Exit(0)
			}
			if url == "-1" {
				break
			}
			Urls = append(Urls, url)
		} else {
			Urls = append(Urls, url)
		}
		i++
	}

	fmt.Println(Urls)
}
