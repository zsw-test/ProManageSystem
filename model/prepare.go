package model

import (
	"ProManageSystem/model/charge"
	"ProManageSystem/model/house"
	"ProManageSystem/model/manager"
	"ProManageSystem/model/owner"
	"fmt"
	"strconv"
)

func PrepareUsers() {
	for i := 0; i < 10; i++ {
		owner := &owner.Owner{
			Username:  "zsw" + strconv.Itoa(i),
			Password:  "123456",
			Houseid:   i + 1,
			Telephone: "12332112332",
		}
		err := owner.Create()
		if err != nil {
			fmt.Println(err.Error())
		}
		manager := &manager.Manager{
			Username:  "admin" + strconv.Itoa(i),
			Password:  "123456",
			Depart:    "维修部",
			Telephone: "12332112332",
		}
		err = manager.Create()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func PrepareHouse() {
	for i := 0; i < 10; i++ {
		house := &house.House{
			Building:     i + 1,
			Unit:         i % 2,
			Area:         90,
			Prorityright: 100,
			HouseType:    "三室一厅",
		}
		err := house.Create()
		if err != nil {
			fmt.Println(err.Error())
		}
		charge := charge.Charge{
			Houseid:  int(house.ID),
			Water:    0,
			Electric: 0,
			Gas:      0,
			Property: 0,
		}
		err = charge.Create()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
