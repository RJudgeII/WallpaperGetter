package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type postData struct {
	Data upperData `json:"data"`
}

type upperData struct {
	Children []childData `json:"children"`
}

type childData struct {
	Data lowerData `json:"data"`
}

type lowerData struct {
	URL string `json:"url"`
}

// GetURL - Gets the url of the current "top of the last 24 hours"
//		post from the /r/wallpapers subreddit
func GetURL() (result string) {

	url := "https://www.reddit.com/r/wallpapers/top/.json?"

	wpClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "wallpaper-getter")

	res, getErr := wpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	posts := postData{}
	jsonErr := json.Unmarshal(body, &posts)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for _, l := range posts.Data.Children {
		link := l.Data.URL
		if link[len(link)-4:] == ".jpg" || link[len(link)-4:] == ".png" {
			result = link
			return
		}
	}

	return
}
