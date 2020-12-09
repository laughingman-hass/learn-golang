package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func testCheckWebsites(t *testing.T) {
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
