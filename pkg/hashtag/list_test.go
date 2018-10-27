package hashtag

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(testHandler))
	defer ts.Close()

	tests := []struct {
		tag   string
		url   string
		posts int64
	}{
		{
			tag:   "golang",
			url:   ts.URL + "/golang",
			posts: 420,
		},
		{
			tag:   "docker",
			url:   ts.URL + "/docker",
			posts: 420,
		},
		{
			tag:   "pet",
			url:   ts.URL + "/pet",
			posts: 420,
		},
		{
			tag:   "error",
			url:   ts.URL + "/error",
			posts: 0,
		},
		{
			tag:   "cat",
			url:   ts.URL + "/cat",
			posts: 420,
		},
	}

	hashtags := make([]string, len(tests))

	for i := range tests {
		hashtags[i] = tests[i].tag
	}

	l := CreateList(hashtags, ts.URL+"/")

	for i, tt := range tests {
		listItem := (*l)[i]

		if tt.tag != listItem.tag {
			t.Fatalf("expected tag %q posts. got=%q", tt.tag, listItem.tag)
		}

		if tt.url != listItem.url {
			t.Fatalf("in tag %q expected %s as URL. got=%s", tt.tag, tt.url, listItem.url)
		}

		if tt.posts != listItem.posts {
			t.Fatalf("in tag %q expected %v posts. got=%v", tt.tag, tt.posts, listItem.posts)
		}
	}
}

func BenchmarkCreatrList(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(testHandler))
	defer ts.Close()

	hashtags := []string{"go", "golang", "docker", "cat", "pet"}

	for i := 0; i < b.N; i++ {
		CreateList(hashtags, ts.URL+"/")
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/error":
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "count: %d", 420)
	}
}
