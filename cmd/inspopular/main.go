package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/danielkvist/inspopular/pkg/hashtag"
)

func main() {
	const url = "https://www.instagram.com/explore/tags/"

	sorted := flag.Bool("s", true, "sort results by popularity. By default is true.")
	flag.Parse()

	tags := flag.Args()
	l := hashtag.Create(tags, url)

	if *sorted {
		sort.Sort(hashtag.OrderedList(*l))
	}

	fmt.Println(l)
}
