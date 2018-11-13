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
		url: url + tag,
	}
}

// FIXME:
var regCount = regexp.MustCompile(".?count.?: ?\\d+")
var regPosts = regexp.MustCompile("[0-9]+")

func (h *hashtag) fetch() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", h.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if checkBadStatus(resp.StatusCode) {
		return fmt.Errorf("got %v at %s", resp.StatusCode, h.url)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	posts := regPosts.Find(regCount.Find(body))
	popularity, err := strconv.Atoi(string(posts))
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	h.posts = int64(popularity)
	return nil
}

func checkBadStatus(sc int) bool {
	if sc != 200 {
		return true
	}

	return false
}
