package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowStebWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://razgriz.io",
		"waat://furhurterwe.geds",
	}

	expected := map[string]bool{
		"http://google.com":       true,
		"http://razgriz.io":       true,
		"waat://furhurterwe.geds": false,
	}

	actual := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("actual %v, expected %v", actual, expected)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStebWebsiteChecker, urls)
	}
}
