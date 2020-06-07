package main

import (
	"encoding/json"
	"fmt"
)

type Creature interface {
	hello()
}

type Person struct {
	Name string `json:"name"`
	Age int `json:"age,string"`
	Nicknames []string `json:"nicknames"`
	PostalCode string `json:"postal_code"`
}

type Person2 struct {
	Name string `json:"name"`
	Age int `json:"age,omitempty"`
	Nicknames []string `json:"nicknames,omitempty"`
	PostalCode string `json:"postal_code"`
	T *T `json:"T,omitempty"`
	T2 T `json:"T2,omitempty"`
}
type T struct {}

func (p Person) hello() {
	fmt.Println("---Person Start---")
}
func (p Person2) hello() {
	fmt.Println("---Person2 Start---")
}

// json.Marshalをラップした処理がかける
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person3 struct {
		Name string
	}
	var p3 Person3
	err := json.Unmarshal(b, &p3)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p3.Name + "!"
	return err
}

func (p Person) MarshalJSON() ([]byte, error) {
	y, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return y,err
}

func main() {
	b := []byte(`{"name":"Mike","age":"20", "nicknames":["a","b","c"], "postal_code":"123-4567"}`)
	var person Person
	outputMarshalAndUnmarshal(b, &person)

	b2 := []byte(`{"name":"mike","age":0, "nicknames":[], "postal_code":"123-4567"}`)
	var person2 Person2
	outputMarshalAndUnmarshal(b2, &person2)

}

func outputMarshalAndUnmarshal(b []byte, creature Creature)  {
	creature.hello()
	if err := json.Unmarshal(b, &creature); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Unmarshal:", creature)
	v, _ := json.Marshal(creature)
	fmt.Println("Marshal:", string(v), "\n")
}
