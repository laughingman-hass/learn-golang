package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

func TestRepeatFor(t *testing.T) {
    repeated := RepeatFor("h", 8)
    expected := "hhhhhhhh"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}

func ExampleRepeat() {
    repeatedString := Repeat("a")
    fmt.Println(repeatedString)
    // Output: aaaaa
}

func ExampleRepeatFor() {
    repeatedString := RepeatFor("h", 8)
    fmt.Println(repeatedString)
    // Output: hhhhhhhh
}
