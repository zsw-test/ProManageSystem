package model

import (
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

func PrepareHouseAndResident() {
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			for f := 1; f <= 10; f++ {
				for d := 1; d <= 4; d++ {
					resident := &house.Resident{}
					house := &house.House{
						Building:     i,
						Unit:         j,
						Door:         f*100 + d,
						Area:         90,
						HouseType:    "三室一厅",
						Prorityright: 70,
					}
					house.Create()
					resident.Name = "zsw" + strconv.Itoa(int(house.ID))
					resident.Age = 18
					resident.Sex = "男"
					resident.Work = "工程师"
					resident.IdCard = 420105199903173612
					resident.HouseId = int(house.ID)
					resident.Create()
				}
			}
		}
	}
}
