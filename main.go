package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/kyokomi/emoji"
)

var Urls = []string{}

type WriteCounter struct {
	n   int // bytes read so far
	bar *pb.ProgressBar
}

const RefreshRate = time.Millisecond * 100

func NewWriteCounter(total int) *WriteCounter {
	b := pb.New(total)
	b.SetRefreshRate(RefreshRate)

	return &WriteCounter{
		bar: b,
	}
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	wc.n += len(p)
	wc.bar.Set(wc.n, nil)
	return wc.n, nil
}

func (wc *WriteCounter) Start() {
	wc.bar.Start()
}

func (wc *WriteCounter) Finish() {
	wc.bar.Finish()
}

func main() {

	// wait for user input
	getFile()
	// ask if more links are there

	// fetch the url concurently
	for i := 0; i < len(Urls); i++ {
		downloadFile(Urls[i])
	}

	// store in the file system

	// show the progress in the terminal
}

func downloadFile(url string) {
	filePaths := strings.Split(url, "/")
	filename := filePaths[len(filePaths)-1]
	out, err := os.Create(filename + ".tmp")

	if err != nil {
		panic(err)
	}

	defer out.Close()

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fsize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	// Create our progress reporter and pass it to be used alongside our writer
	counter := NewWriteCounter(fsize)
	counter.Start()

	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		panic(err)
	}

	counter.Finish()
	err = os.Rename(filename+".tmp", filename)
	if err != nil {
		panic(err)
	}
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

		ifURL, _ := regexp.MatchString("https?://(www.)?", url)

		if url == "-1" {
			break
		} else if url == "0" {
			os.Exit(0)
		} else {
			if ifURL {
				Urls = append(Urls, url)
				i++

			} else {
				fmt.Println("\nPlease Try Again... ")
			}
		}

	}

	fmt.Println(Urls)
}
