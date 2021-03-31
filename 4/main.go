package main

import (
	"fmt"
	"sort"
)

func main(){
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	fmt.Println(groupAnagrams(words))

}


func groupAnagrams(stringCollection []string) [][]string {
	stringMap := make(map[string][]string)
	for _, v := range stringCollection {
		bytes := []byte(v)
		sort.SliceStable(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		str := string(bytes)
		stringMap[str] = append(stringMap[str], v)
	}

	var result [][]string
	for key := range stringMap {
		result = append(result, stringMap[key])
	}
	return result
}
