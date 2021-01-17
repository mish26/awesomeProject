package main

import (
	"fmt"
)

type ItemSelectOption struct {
	ItemSelectOptionGetter
	order          string
}

func NewItemSelectOption(optionId ItemSelectOptionGetter, order string) ItemSelectOption {
	return ItemSelectOption{optionId, order}
}

type ItemSelectOptionGetter interface {
	GetId() (id string)
}

type ItemSelectOptionMakerId struct {
	makerId string
}

func NewItemSelectOptionMakerId(makerId string) *ItemSelectOptionMakerId {
	return &ItemSelectOptionMakerId{makerId: makerId}
}

func (c ItemSelectOptionMakerId) GetId() string{
	return c.makerId
}

type ItemSelectOptionBrandId struct {
	brandId string
}

func NewItemSelectOptionBrandId(brandId string) *ItemSelectOptionBrandId {
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

	checkType(option)
	checkType(option2)
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


