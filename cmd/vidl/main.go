package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tdd28/vidl"
)

func main() {
	url, output := os.Args[1], os.Args[2]

	fmt.Printf("Downloading %s to %s\n", url, output)

	vidl, err := vidl.GetMediaPlaylist(url)
	if err != nil {
		log.Fatal(err)
	}
	defer vidl.ReleasePlaylist()

	if err := vidl.Download(output); err != nil {
		log.Fatal(err)
	}
}
