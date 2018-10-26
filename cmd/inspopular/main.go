package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/danielkvist/inspopular/pkg/hashtag"
)

func main() {
	const instaURL = "https://www.instagram.com/explore/tags/"

	nOrder := flag.Bool("n", true, "sort results by popularity. By default is true.")
	flag.Parse()

	tags := flag.Args()
	l := hashtag.CreateList(tags, instaURL)

	if *nOrder {
		sort.Sort(hashtag.OrderedList(*l))
	}

	fmt.Println(l)
}
