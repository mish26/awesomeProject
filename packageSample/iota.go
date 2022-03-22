package packageSample

import "fmt"

const (
	c1 = iota
	c2 = iota
	c3 = iota
)

const (
	_ = iota
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
)

type Order string

//var TimeNewlyListed = Order{"TimeNewlyListed"}
const (
	TimeNewlyListed Order = "TimeNewlyListed"
//	PriceLowestFirst = Order("PriceLowestFirst")
//	PriceHighestFirst =  Order("PriceHighestFirst")
)


type Direction string

const (
	North = Direction("North")
	East = Direction("East")
	South = Direction("South")
	West = Direction("West")
)

type Season struct{ value string }

var Summer Season = Season{"Summer"}
//var Autumn = Season{"Autumn"}
//var Winter = Season{"Winter"}

func (s Season) String() string {
	if s.value == "" {
		return "未定義"
	}
	return s.value
}


func main() {

	fmt.Println(c1,c2,c3)
	fmt.Println(KB,MB,GB,TB)
	testDirection(North)
	testDirection("test")
	testOrder(TimeNewlyListed)
	testOrder("test")
	testSeason(Summer)
	Summer = Season{"Spring"}
	testSeason(Summer)

}

func testDirection(dir Direction)  {
	fmt.Println(dir)
}

func testOrder(order Order) {
	fmt.Println(order)
}

func testSeason(s Season)  {
	fmt.Println(s)
}



