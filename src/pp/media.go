package pp

import (
	"net/url"
)

type MediaQuality struct {
	Width int
	Height int
}

func (this *MediaQuality) IsHD() bool {
	if this.Width == 1280 && this.Height == 720 {
		return true
	}
	if this.Width == 1920 && this.Height == 1080 {
		return true
	}
	if this.Width == 2560 && this.Height == 1440 {
		return true
	} 
	return false
}

func (this *MediaQuality) IsUHD() bool {
	if this.Width == 2048 && this.Height == 2000 {
		return true
	}
	if this.Width == 3840 && this.Height == 2160 {
		return true
	}
	if this.Width == 4520 && this.Height == 2540 {
		return true
	}
	if this.Width == 4096 && this.Height == 3072 {
		return true
	}
	if this.Width == 7680 && this.Height == 4320 {
		return true
	}

	return false;
}

type Media struct {
	link url.URL
	quality MediaQuality
	imdb int
	kinopoisk int
}