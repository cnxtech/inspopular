package hashtag

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"text/tabwriter"
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
			if err := list[i].popularity(); err != nil {
				log.Println(err)
			}
		}(i)
	}

	w.Wait()
	return &list
}

func (l *List) Print() {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Hashtag", "URL", "Posts")
	fmt.Fprintf(tw, format, "-------", "---", "-----")

	fmt.Println()
	for _, item := range *l {
		fmt.Fprintf(tw, format, item.tag, item.url, item.posts)
	}
	tw.Flush()
}

func newhashtag(tag string) *hashtag {
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
