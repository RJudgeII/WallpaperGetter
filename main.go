package main

import (
	"./fetcher"
	"./wallpaper"
)

func main() {

	url := fetcher.GetURL()
	wallpaper.SetFromURL(url)

}
