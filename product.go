package main

import(
	"sort"
)

type product_struct struct {
	freq int
	product string
}

// by_freq implements sort.Interface for []product_struct based on the freq field
type by_freq []product_struct
func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq > a[j].freq }

// iterate over the slice of items and populate
// the map with item:frequency pairs
func freq_count(items []string) map[string]int {
	freq := make(map[string]int)
	// index is not needed, so use blank identifier _
	for _, item := range items {
		// check if item is already in the map
		_, ok := freq[item]
		// if true add 1 to frequency (value of map)
		// else start frequency at 1
		if ok == true {
			freq[item] += 1
		} else {
			freq[item] = 1
		}
	}
	return freq
}

// Sort a map with item:frequency pairs by frequency
func sort_map(similar_map map[string]int) []product_struct {
	pws := new(product_struct)
	similar_sorted := make([]product_struct, len(similar_map))
	ix := 0
	for key, val := range similar_map {
		pws.freq = val
		pws.product = key
		// test, %+v shows field names
		//fmt.Printf("%v %v  %+v\n", pws.freq, pws.word, *pws)
		similar_sorted[ix] = *pws
		ix++
	}
	sort.Sort(by_freq(similar_sorted))
	return similar_sorted
}