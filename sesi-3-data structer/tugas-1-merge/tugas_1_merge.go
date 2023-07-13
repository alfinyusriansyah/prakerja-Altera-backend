package main

import "fmt"

func main() {
	array1 := []string{"John", "Alice", "Bob", "Mary"}
	array2 := []string{"Alice", "David", "Emma", "John"}

	mergedArray := mergeArrays(array1, array2)
	uniqueArray := removeDuplicates(mergedArray)

	fmt.Println("Hasil penggabungan dan penghapusan duplikat:")
	fmt.Println(uniqueArray)
}

func mergeArrays(arr1, arr2 []string) []string {
	merged := append(arr1, arr2...)
	return merged
}

func removeDuplicates(arr []string) []string {
	unique := make([]string, 0)
	visited := make(map[string]bool)

	for _, val := range arr {
		if !visited[val] {
			unique = append(unique, val)
			visited[val] = true
		}
	}

	return unique
}
