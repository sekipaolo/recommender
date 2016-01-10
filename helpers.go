package main

import (
	"strconv"
	"fmt"
	"os"
)

func sliceContains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}


func deleteFromSlice(s []string, v string) []string {
	var r []string
	for i := range s {
		item := s[i]
		if item != v {
			r = append(r, item)
		}
	}
	return r
}

func stringToInt(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
    fmt.Printf("error converting %v to integer \n", s, err)
    os.Exit(1)
	}
	return res
}
