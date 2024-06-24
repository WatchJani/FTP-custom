package trie

import (
	"strings"
	"testing"
)

type Handler func()

func BenchmarkSpeedTest(b *testing.B) {
	b.StopTimer()

	trie := NewTrie[Handler]()

	allEndpoints := []struct {
		key   string
		value Handler
	}{
		{
			key: "dree/{more}/macka",
			value: func() {
				// fmt.Println("NEAH")
			},
		},
	}

	for _, endpoint := range allEndpoints {
		key := strings.Split(endpoint.key, "/")
		trie.Insert(key, endpoint.value)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		value, err := trie.Search([]string{"dree", "janko", "macka"})
		if err != nil {
			return
		}

		value()
	}
}
