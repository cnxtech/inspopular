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
		posts int
	}{
		{
			tag:   "golang",
			url:   ts.URL + "/golang",
			posts: 1200,
		},
		{
			tag:   "error",
			url:   ts.URL + "/error",
			posts: 0,
		},
	}

	hashtags := make([]string, len(tests))
	for i := range tests {
		hashtags[i] = tests[i].tag
	}

	l := Create(hashtags, ts.URL+"/")
	for i, tt := range tests {
		listItem := (*l)[i]

		if tt.tag != listItem.tag {
			t.Errorf("expected tag %q posts. got=%q", tt.tag, listItem.tag)
		}

		if tt.url != listItem.url {
			t.Errorf("in tag %q expected %s as URL. got=%s", tt.tag, tt.url, listItem.url)
		}

		if tt.posts != listItem.posts {
			t.Errorf("in tag %q expected %v posts. got=%v", tt.tag, tt.posts, listItem.posts)
		}
	}
}

func BenchmarkCreate(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(testHandler))
	defer ts.Close()

	url := ts.URL + "/"
	hashtags := []string{"go", "golang", "docker", "cat", "pet"}
	for i := 0; i < b.N; i++ {
		Create(hashtags, url)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	switch r.URL.Path {
	case "/error":
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<h1>Simulated counter</h1><p>\"count\": %d<p><p>count: 250</p><p>counting : 500</p>}", 1200)
	}
}
