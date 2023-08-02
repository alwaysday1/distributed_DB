package main

import (
	"fmt"
	"strings"
)

func Map(texts []string, mapper func(string) map[string]int) []map[string]int {
	var intermediate []map[string]int
	for _, text := range texts {
		result := mapper(text)
		intermediate = append(intermediate, result)
	}
	return intermediate
}

func wordCountMap(text string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		result[word]++
	}
	return result
}

func Reduce(intermediate []map[string]int, reducer func(map[string]int) map[string]int) map[string]int {
	result := make(map[string]int)
	for _, item := range intermediate {
		reduced := reducer(item)
		for k, v := range reduced {
			result[k] += v
		}
	}
	return result
}

func wordCountReduce(item map[string]int) map[string]int {
	return item // 在这个例子中，Reduce 直接返回输入
}

func main() {
	texts := []string{
		"hello world",
		"world is beautiful",
		"hello Go",
	}

	// Map 阶段
	intermediate := Map(texts, wordCountMap)

	// Reduce 阶段
	result := Reduce(intermediate, wordCountReduce)

	fmt.Println(result) // 输出：map[Go:1 beautiful:1 hello:2 is:1 world:2]
}
