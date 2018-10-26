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
	posts int64
}

func newHashtag(tag, url string) *hashtag {
	return &hashtag{
		tag: tag,
		url: url,
	}
}

func (h *hashtag) fetch() error {
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

	regCount := regexp.MustCompile(".?count.?: ?\\d+")
	regPosts := regexp.MustCompile("[0-9]+")
	posts := regPosts.Find(regCount.Find(body))

	popularity, err := strconv.Atoi(string(posts))

	if err != nil {
		log.Fatal(err)
	}

	h.posts = int64(popularity)
	return nil
}
