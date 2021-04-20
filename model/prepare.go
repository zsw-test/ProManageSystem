package model

import (
	"ProManageSystem/model/charge"
	"ProManageSystem/model/house"
	"ProManageSystem/model/manager"
	"ProManageSystem/model/owner"
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/model/repair"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func PreparePark() {
	for i := 0; i < 100; i++ {
		park := &parkmodel.Park{
			Status:   parkmodel.Empty,
			Location: "A-" + strconv.Itoa(i),
		}
		park.Create()
	}
}
func PrepareComplaint() {

}

func PrepareUsersAndRepair() {
	for i := 0; i < 10; i++ {
		owner := &owner.Owner{
			Username:  "zsw" + strconv.Itoa(i),
			Password:  "123456",
			Houseid:   i + 1,
			Telephone: "12332112332",
			Nickname:  "业主test" + strconv.Itoa(i),
		}
		err := owner.Create()
		if err != nil {
			fmt.Println(err.Error())
		}
		repair := &repair.Repair{
			Ownername: owner.Username,
			Address:   "37栋2204",
			Reason:    "水龙头坏了",
			Pics:      `["https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fp2.zx.680.com%2F2016-07%2F20%2F20160720155702715888.png&refer=http%3A%2F%2Fp2.zx.680.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1619315604&t=fc55cb0a7fe354b0920156e917c070eb"]`,
		}
		repair.Create()

		manager := &manager.Manager{
			Username:  "admin" + strconv.Itoa(i),
			Password:  "123456",
			Depart:    "维修部",
			Telephone: "12332112332",
			Nickname:  "管理员test" + strconv.Itoa(i),
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
					resident.IdCard = "420105199903173612"
					resident.HouseId = int(house.ID)
					resident.Create()
					rand.Seed(time.Now().Unix())
					// 每个房屋必须有费用表
					charge := &charge.Charge{
						Houseid:  int(house.ID),
						Water:    float64(rand.Intn(50)),
						Electric: float64(rand.Intn(50)),
						Gas:      float64(rand.Intn(50)),
						Property: 0,
					}
					charge.Create()
				}
			}
		}
	}
}

func PrepareAll() {
	PrepareHouseAndResident()
	PrepareUsersAndRepair()
	PreparePark()
}
