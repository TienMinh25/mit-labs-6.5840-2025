package main

//
// an indexing application "plugin" for MapReduce.
//
// go build -buildmode=plugin indexer.go
//

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/TienMinh25/mit-labs-6-824-2025/mapreduce/types"
)

// The mapping function is called once for each piece of the input.
// In this framework, the key is the name of the file that is being processed,
// and the value is the file's contents. The return value should be a slice of
// key/value pairs, each represented by a mr.KeyValue.
func Map(document string, value string) (res []types.KeyValue) {
	m := make(map[string]bool)
	words := strings.FieldsFunc(value, func(x rune) bool { return !unicode.IsLetter(x) })
	for _, w := range words {
		m[w] = true
	}
	for w := range m {
		kv := types.KeyValue{w, document}
		res = append(res, kv)
	}
	return
}

// The reduce function is called once for each key generated by Map, with a
// list of that key's string value (merged across all inputs). The return value
// should be a single output value for that key.
func Reduce(key string, values []string) string {
	sort.Strings(values)

	i := 0
	res := []string{values[i]}

	for j := i + 1; j < len(values); j++ {
		if values[i] == values[j] {
			continue
		}

		i = j
		res = append(res, values[j])
	}

	return fmt.Sprintf("%d %s", len(res), strings.Join(res, ","))
}
