package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	fmt.Println("hello Golang")
	// rev_slice()
	// rev_map()
	// push_style()
	chain_iter()
}

func rev_slice() {
	nums := []int{1, 2, 3, 4, 5}
	for _, v := range slices.Backward(nums) {
		fmt.Println(v)
	}
}

func rev_map() {
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k, v := range maps.All(m) {
		fmt.Println(k, v)
	}
}

func push_style() {
	// sequense
	seq := func(yield func(int) bool) {
		yield(1)
		fmt.Println("after yeild(1)")
		yield(2)
		fmt.Println("after yeild(2)")
		yield(3)
		fmt.Println("after yeild(3)")
	}

	// やりたい処理はfor内に書くが、関数値としてseqに渡されて実行される。
	for i := range seq {
		fmt.Printf("hello %v番目の少年\n", i)
	}
	// 脱糖するとこんな感じ
	// seq(func(i int) bool {
	// 	fmt.Printf("hello %v番目の少年\n", i)
	// 	return true
	// })

}

func chain_iter() {
	// 0..=9のintを返すイテレータを返す関数
	// numbers := func() func(func(int) bool) {
	// 	return func(yeild func(int) bool) {
	// 		for i := range 10 {
	// 			yeild(i)
	// 		}
	// 	}
	// }

	// 0..=9のintを返すpush型seq
	numbers := func(yield func(int) bool) {
		for i := range 10 {
			yield(i)
		}
	}
	// seq -> seq
	even := func(seq func(func(int) bool)) func(func(int) bool) {
		return func(yield func(int) bool) {
			for i := range seq {
				if i%2 == 0 {
					yield(i)
				}
			}
			// 脱糖したらこんな感じ
			// seq(func(i int) bool {
			// 	if i%2 == 0 {
			// 		return yield(i)
			// 	}
			// 	return true
			// })
		}
	}

	double := func(seq func(func(int) bool)) func(func(int) bool) {
		return func(yield func(int) bool) {
			for i := range seq {
				yield(i * 2)
			}
		}
	}
	for i := range double(even(numbers)) {
		fmt.Println(i)
	}
}
