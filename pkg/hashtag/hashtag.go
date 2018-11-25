// Package hashtag is a simple utility
// to get the popularity of different hashtags at Instagram.
package hashtag

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type hashtag struct {
	tag   string
	url   string
	posts int
}

func newHashtag(tag, url string) *hashtag {
	return &hashtag{
		tag: tag,
		url: url + tag,
	}
}

func fetch() func(string) (int, error) {
	regCount := regexp.MustCompile(".?count.?: ?\\d+")
	regPosts := regexp.MustCompile("[0-9]+")

	return func(url string) (int, error) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return 0, fmt.Errorf("error at creating new request: %v", err)
		}

		resp, err := client.Do(req)
		if checkBadStatus(resp.StatusCode) {
			return 0, fmt.Errorf("got %v at url %s", resp.StatusCode, url)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 0, fmt.Errorf("error at reading the response body: %v", err)
		}

		posts := regPosts.Find(regCount.Find(body))
		popularity, err := strconv.Atoi(string(posts))
		if err != nil {
			return 0, fmt.Errorf("error at converting the number of posts from string to an int: %v", err)
		}

		return popularity, nil
	}
}

func checkBadStatus(sc int) bool {
	if sc != http.StatusOK {
		return true
	}

	return false
}
