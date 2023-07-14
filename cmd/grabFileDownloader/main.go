package main

import (
	"bufio"
	"fmt"
	"os"

	downloader "github.com/Xpl0itU/grabFileDownloader"
	"github.com/cavaliergopher/grab/v3"
)

func main() {
	var urls []string

	if len(os.Args) > 1 {
		filePath := os.Args[1]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			url := scanner.Text()
			urls = append(urls, url)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			return
		}
	} else {
		fmt.Print("Enter the URL: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
	}

	client := grab.NewClient()
	for _, url := range urls {
		err := downloader.DownloadFile(client, url, ".")
		if err != nil {
			fmt.Printf("Error downloading file from %s: %s\n", url, err)
		}
	}
}
