package hashtag

import (
	"bytes"
	"fmt"
	"log"
	"sync"
	"text/tabwriter"
)

// List is a slice of *hashtag.
type List []*hashtag

// Create creates a new List of *hashtags
// and fetch the popularity of each one.
func Create(tags []string, url string) *List {
	var w sync.WaitGroup
	var list List

	list = make([]*hashtag, len(tags))
	for i, t := range tags {
		list[i] = newHashtag(t, url)
	}

	get := fetch()

	w.Add(len(tags))
	for _, h := range list {
		go func(h *hashtag) {
			defer w.Done()
			posts, err := get(h.url)
			if err != nil {
				log.Println(err)
			}
			h.posts = posts
		}(h)
	}

	w.Wait()
	return &list
}

// String is a method of the List type to satisfy
// the stringer interface.
func (l *List) String() string {
	var out bytes.Buffer

	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(&out, 0, 4, 4, ' ', 0)
	fmt.Fprintf(tw, format, "Hashtag", "URL", "Posts")
	fmt.Fprintf(tw, format, "-------", "---", "-----")

	for _, item := range *l {
		fmt.Fprintf(tw, format, item.tag, item.url, item.posts)
	}

	tw.Flush()
	out.WriteString("\n")

	for _, item := range *l {
		h := fmt.Sprintf("#%s ", item.tag)
		out.WriteString(h)
	}

	return out.String()
}

// OrderedList type is a slice of *hashtag that
// implements the sort interface.
type OrderedList []*hashtag

// Len is a method of the OrderedList type to satisfy
// the sort interface.
func (ol OrderedList) Len() int {
	return len(ol)
}

// Less is a method of the OrderedList type to satisfy
// the sort interface.
func (ol OrderedList) Less(i, j int) bool {
	return ol[i].posts > ol[j].posts
}

// Swap is a method of the OrderedList type to satisfy
// the sort interface.
func (ol OrderedList) Swap(i, j int) {
	ol[i], ol[j] = ol[j], ol[i]
}
