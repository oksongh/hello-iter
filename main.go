package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

func main() {
	fmt.Println("hello Golang")
	// rev_slice()
	// rev_map()
	// push_style()
	// chain_iter_push()
	// pull_style()
	// may_raise_panic()
	tree_ex()

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

// https://zenn.dev/team_soda/articles/understanding-iterators-in-go
func push_style() {
	// sequence
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

func chain_iter_push() {
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
	even := func(seq iter.Seq[int]) iter.Seq[int] {
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

	double := func(seq iter.Seq[int]) iter.Seq[int] {
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

func pull_style() {
	seq := slices.Values([]int{1, 2, 3})
	next, stop := iter.Pull(seq)
	defer stop() // 誰？
	for {
		i, ok := next()
		fmt.Println(next)
		if !ok {
			fmt.Println("end")
			break
		}

		fmt.Println(i)
	}
}

func may_raise_panic() {
	numbers := func(yield func(int) bool) {
		for i := range 10 {
			// yieldの戻り値は
			// true: 続くとき
			// false: 終了のとき

			// yield(i) // panic!!!
			if !yield(i) {
				fmt.Println("!yield!")
				return
			}
		}
	}

	fmt.Println("break")
	for i := range numbers {
		fmt.Println(i)
		if i == 4 {
			break
		}
	}

	fmt.Println("continue")
	for i := range numbers {
		if i%2 == 0 || i == 7 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("return")
	for i := range numbers {
		if i == 4 {
			return
		}
		fmt.Println(i)
	}

}

// https://go.dev/blog/range-functions#binary-tree
type Tree[E any] struct {
	value       E
	left, right *Tree[E]
}

func (t Tree[E]) all() iter.Seq[E] {
	return func(yield func(E) bool) {
		t.push(yield)
	}
}
func (t *Tree[E]) push(yield func(E) bool) bool {
	if t == nil {
		return true
	}
	// 左、中央、右の順で実行
	return t.left.push(yield) &&
		yield(t.value) &&
		t.right.push(yield)
}

func tree_ex() {
	tree := Tree[string]{"root",
		&Tree[string]{"left1", nil, nil},
		&Tree[string]{"right1",
			&Tree[string]{"r1-l2", nil, nil}, &Tree[string]{"r1-r2", nil, nil}}}

	fmt.Println(tree)
	for v := range tree.all() {
		fmt.Println(v)
	}
}
