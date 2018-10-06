package hashtag

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
)

type List []*hashtag

type hashtag struct {
	tag   string
	url   string
	posts int
}

const instaURL = "https://www.instagram.com/explore/tags/"

func CreateList(tags []string) *List {
	var w sync.WaitGroup
	var list List

	for _, t := range tags {
		list = append(list, newhashtag(t))
	}

	w.Add(len(tags))
	for i := range list {
		go func(i int) {
			defer w.Done()
			list[i].popularity()
		}(i)
	}

	w.Wait()
	return &list
}

func newhashtag(tag string) *hashtag {
	h := &hashtag{tag: tag, url: instaURL + tag}
	return h
}

func (h *hashtag) popularity() {
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
