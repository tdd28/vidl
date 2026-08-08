package vidl

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Eyevinn/hls-m3u8/m3u8"
)

var ErrNotMediaPlaylist = errors.New("not a media playlist")

type MediaPlaylist struct {
	*m3u8.MediaPlaylist

	m3u8URL *url.URL
}

func GetMediaPlaylist(m3u8url string) (*MediaPlaylist, error) {
	m3u8URL, err := url.Parse(m3u8url)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(m3u8URL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	p, lt, err := m3u8.DecodeFrom(resp.Body, true)
	if err != nil {
		return nil, err
	}

	switch lt {
	case m3u8.MEDIA:
		return &MediaPlaylist{
			MediaPlaylist: p.(*m3u8.MediaPlaylist),
			m3u8URL:       m3u8URL,
		}, nil
	default:
		return nil, ErrNotMediaPlaylist
	}
}

func (p *MediaPlaylist) Download(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, seg := range p.GetAllSegments() {
		segURL, err := url.Parse(seg.URI)
		if err != nil {
			return err
		}

		resp, err := http.Get(p.m3u8URL.ResolveReference(segURL).String())
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if _, err := io.Copy(file, resp.Body); err != nil {
			return err
		}
	}

	return nil
}
