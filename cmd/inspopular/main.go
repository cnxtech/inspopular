package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/danielkvist/inspopular/pkg/hashtag"
)

func main() {
	nOrder := flag.Bool("n", true, "ordered by popularity. By default is true")
	flag.Parse()

	tags := flag.Args()
	l := hashtag.CreateList(tags)

	if *nOrder {
		sort.Sort(hashtag.OrderedList(*l))
	}

	fmt.Println(l)
}
