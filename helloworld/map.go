package main

import "fmt"

func main() {
	m := map[string]int{"apple":100, "banana":200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["apple"] = 300
	fmt.Println(m)
	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, existsKey := m["nothing"]
	fmt.Println(v2, existsKey)

	// メモリ上に空のmapを作る
	m2 := make(map[string]int)
	m2["sp"] = 5000
	fmt.Println(m2)

	/* メモリ上に、領域を確保していないのでエラーになる
	var m3 map[string]int
	m3["pc"] = 5000
	fmt.Println(m3)
	*/

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}

	b := []byte{74, 75}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("JK")
	fmt.Println(c)
}


