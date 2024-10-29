package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortWord(s string) string {
	s = strings.ToLower(s)
	tmp := strings.Split(s, "")
	sort.Strings(tmp)
	return strings.Join(tmp, "")
}

func makeUniqSlice(bucket []string) []string {
	set := make(map[string]bool)
	var setToSlice []string
	for _, v := range bucket {
		v = strings.ToLower(v)
		set[v] = true
	}
	for i := range set {
		setToSlice = append(setToSlice, i)
	}
	sort.Strings(setToSlice)
	return setToSlice
}
func filterAnagrams(res map[string][]string) map[string][]string {
	filtered := make(map[string][]string)
	for key, values := range res {
		if len(values) > 1 {
			filtered[key] = values
		}
	}
	return filtered
}

func findAnagram(arr *[]string) map[string][]string {
	anag := make(map[string][]string, 0)
	anagV := make(map[string][]int, 0)
	res := make(map[string][]string, 0)
	anagKeys := make([]string, 0)

	for i, v := range *arr {
		sorted := sortWord(v)
		anagV[sorted] = append(anagV[sorted], i)
		anag[sorted] = append(anag[sorted], v)

	}

	for k := range anag {
		anagKeys = append(anagKeys, k)
	}

	for _, v := range anagKeys {
		min := anagV[v][0]
		res[(*arr)[min]] = makeUniqSlice(anag[v])
	}

	res = filterAnagrams(res)

	return res
}

func printAnagrams(bucket map[string][]string) {
	for k, v := range bucket {
		fmt.Printf("%v : %v\n", k, v)

	}
}

func main() {
	words := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик", "ПяТаК", "ЛиСтОк", "тест", "рог", "гор"}
	anagrams := findAnagram(&words)
	printAnagrams(anagrams)
}
