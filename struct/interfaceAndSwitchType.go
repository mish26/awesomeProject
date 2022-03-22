package main

import (
	"awesomeProject/packageSample"
	"fmt"
	"log"
)

type ItemSelectOption struct {
	ItemSelectOptionGetter
	order packageSample.Order
}

func NewItemSelectOption(optionId ItemSelectOptionGetter, order packageSample.Order) ItemSelectOption {
	return ItemSelectOption{optionId, order}
}

type ItemSelectOptionGetter interface {
	GetId() (id string)
}

type ItemSelectOptionMakerId struct {
	makerId string
}

func NewItemSelectOptionMakerId(makerId string) *ItemSelectOptionMakerId {
	if makerId == "" {
		log.Fatal("makerId is empty")
	}
	return &ItemSelectOptionMakerId{makerId: makerId}
}

func (c ItemSelectOptionMakerId) GetId() string{
	return c.makerId
}

type ItemSelectOptionBrandId struct {
	brandId string
}

func NewItemSelectOptionBrandId(brandId string) *ItemSelectOptionBrandId {
	if brandId == "" {
		log.Fatal("brandId is empty")
	}
	return &ItemSelectOptionBrandId{brandId: brandId}
}

func (c ItemSelectOptionBrandId) GetId() string {
	return c.brandId
}


func main() {
	var o ItemSelectOptionGetter =  NewItemSelectOptionMakerId("maker12345")
	var option ItemSelectOption = NewItemSelectOption(o, "o")
	fmt.Println(option.GetId())

	var o2 ItemSelectOptionGetter =  NewItemSelectOptionBrandId("brand12345")
	var option2 ItemSelectOption = NewItemSelectOption(o2, "o2")
	fmt.Println(option2.GetId())

	var og3 ItemSelectOptionGetter =  NewItemSelectOptionBrandId("")
	var option3 ItemSelectOption = NewItemSelectOption(og3, "o2")
	fmt.Println(option3.GetId())

	checkType(option)
	checkType(option2)
	test(packageSample.Summer)

}

func checkType(og ItemSelectOption) {
	switch og.ItemSelectOptionGetter.(type) {
	case *ItemSelectOptionBrandId:
		fmt.Println("brandId:" + og.GetId())
	case *ItemSelectOptionMakerId:
		fmt.Println("makerId:" + og.GetId())
	default:
		fmt.Printf("UnKnownType")
	}
}

func test(season packageSample.Season)  {
	fmt.Println(season)
}
