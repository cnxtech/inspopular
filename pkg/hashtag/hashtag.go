package hashtag

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type hashtag struct {
	tag   string
	url   string
	posts int
}

var instaURL = "https://www.instagram.com/explore/tags/"

func New(tag string) *hashtag {
	return &hashtag{tag: tag, url: instaURL + tag}
}

func (h *hashtag) Popularity() {
	resp, err := http.Get(h.url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	reCount := regexp.MustCompile(`"count":\d+`)
	rePosts := regexp.MustCompile("[0-9]+")
	posts := rePosts.Find(reCount.Find(body))

	popularity, err := strconv.Atoi(string(posts))

	if err != nil {
		log.Fatal(err)
	}

	h.posts = popularity
}
