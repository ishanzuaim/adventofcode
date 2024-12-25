package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	strs := readFile()

	m := make(map[string][]string)

	for _, val := range strs {
		vals := strings.Split(val, "-")
		_, ok := m[vals[0]]
		if !ok {
			m[vals[0]] = []string{}
		}
		_, ok = m[vals[1]]
		if !ok {
			m[vals[1]] = []string{}
		}
		m[vals[0]] = append(m[vals[0]], vals[1])
		m[vals[1]] = append(m[vals[1]], vals[0])
	}
    // __AUTO_GENERATED_PRINT_VAR_START__
    fmt.Println(fmt.Sprintf("main m: %v", m)) // __AUTO_GENERATED_PRINT_VAR_END__
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	total := make([][]string, 0)
	checked := make(map[string]bool, 0)
	for _, key := range keys {
		for _, vals := range m[key] {
			pairs := intersection(key, m[key], m[vals])
			ls := []string{key, vals}
			check := append(ls, pairs...)
			// for _, val := range pairs {
			srl := serialize(check)
			_, ok := checked[srl]
			if !ok {
				total = append(total, check)
				checked[srl] = true
			}
			// }
		}
  }
  // __AUTO_GENERATED_PRINT_VAR_START__
  fmt.Println(fmt.Sprintf("main total: %v", total)) // __AUTO_GENERATED_PRINT_VAR_END__
	}

func anyContains(ls []string, chr byte) bool {
	for _, val := range ls {
		if val[0] == chr {
			return true
		}
	}
	return false
}

func serialize(nsew []string) string {
  ls := slices.Clone(nsew)
  slices.Sort(ls)
	return strings.Join(ls, "-")
}

func intersection(b string, als, bls []string) []string {
	new := make([]string, 0)
	for _, val := range als {
		if val != b && slices.Contains(bls, val) {
			new = append(new, val)
		}
	}
	return new
}
