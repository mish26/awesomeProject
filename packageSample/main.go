package packageSample

import (
	"awesomeProject/packageSample/test"
	"fmt"
)

func main() {
	s := []int{10,20,30,40}
	fmt.Println(test.Average(s))
	test.Say()
	animal := test.Animal{Name: "256",Age: 10}
	fmt.Println(animal)
	fmt.Println(test.PublicVal)

	// 先頭、lowerのstructは同じパッケージからしか参照できない
	// 他のパッケージからは大文字のものしか参照できない
	// 変数も同様
	// animal2 := test.animal2{Name: "256",Age: 10}
	// fmt.Println(test.privateVal)

}


