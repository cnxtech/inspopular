// Package hashtag is a simple utility
// to get the popularity of different hashtags at Instagram.
package hashtag

import (
	"fmt"
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

const instaURL = "https://www.instagram.com/explore/tags/"

func newHashtag(tag string) *hashtag {
	h := &hashtag{tag: tag, url: instaURL + tag}
	return h
}

func (h *hashtag) popularity() error {
	resp, err := http.Get(h.url)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		err := fmt.Errorf("got %v at %s", http.StatusNotFound, h.url)
		return err
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
	return nil
}
