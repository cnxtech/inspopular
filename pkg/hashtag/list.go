package hashtag

import (
	"fmt"
	"log"
	"os"
	"sync"
	"text/tabwriter"
)

type List []*hashtag

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
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 4, 4, ' ', 0)
	fmt.Fprintf(tw, format, "Hashtag", "URL", "Posts")
	fmt.Fprintf(tw, format, "-------", "---", "-----")

	fmt.Println()
	for _, item := range *l {
		fmt.Fprintf(tw, format, item.tag, item.url, item.posts)
	}
	fmt.Println()
	tw.Flush()
}

type OrderedList []*hashtag

func (ol OrderedList) Len() int {
	return len(ol)
}

func (ol OrderedList) Less(i, j int) bool {
	return ol[i].posts > ol[j].posts
}

func (ol OrderedList) Swap(i, j int) {
	ol[i], ol[j] = ol[j], ol[i]
}
