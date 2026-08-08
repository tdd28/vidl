package main

import (
	"log"

	"github.com/tdd28/vidl"
)

const (
	url        = "https://test-streams.mux.dev/x36xhzz/url_6/193039199_mp4_h264_aac_hq_7.m3u8"
	outputFile = "output.ts"
)

func main() {
	mp, err := vidl.GetMediaPlaylist(url)
	if err != nil {
		log.Fatal(err)
	}
	defer mp.ReleasePlaylist()

	if err := mp.Download(outputFile); err != nil {
		log.Fatal(err)
	}
}
