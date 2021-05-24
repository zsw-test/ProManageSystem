package model

import (
	"ProManageSystem/model/accessrecord"
	"ProManageSystem/model/announce"
	"ProManageSystem/model/charge"
	"ProManageSystem/model/complaint"
	"ProManageSystem/model/expressage"
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

func PrepareUserDatas() {
	for i := 1; i <= 10; i++ {
		Username := "zsw" + strconv.Itoa(i)

		repair := &repair.Repair{
			Ownername: Username,
			Address:   "37栋2204",
			Reason:    "水龙头坏了",
			Pics:      `["https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=1185614971,854605554&fm=26&gp=0.jpg"]`,
		}
		repair.Create()
		complaint := &complaint.Complaint{
			Ownername: Username,
			Depart:    "维修部门",
			Reason:    "修不好我家的门 希望能够好好解决  态度极差 呵呵~",
		}
		complaint.Create()

		expressage := &expressage.Expressage{
			Ownername:       Username,
			ExpressLocation: "北门门口",
			ExpType:         "生活用品",
			Telephone:       "18627031721",
		}
		expressage.Create()

		manager := &manager.Manager{
			Username:  "admin" + strconv.Itoa(i),
			Password:  "123456",
			Depart:    "维修部",
			Telephone: "18627031313",
			Nickname:  "管理员" + strconv.Itoa(i),
		}
		manager.Create()

		announce := &announce.Announce{
			Title:       "停水公告",
			Text:        "停水通知：本小区因为线路维修于2021年5月29号下午14点-20点停水，请业主们相互转告 谢谢配合！",
			Managername: manager.Username,
		}
		announce.Create()

		accessrecord := accessrecord.Access{
			Name: Username,
			Way:  "人脸识别",
		}
		accessrecord.Create()
		parkinfo := parkmodel.ParkInfo{
			CarNumber: "鄂A123123" + strconv.Itoa(i),
		}
		parkinfo.Create()
		carinfo := parkmodel.CarInfo{
			CarNumber: parkinfo.CarNumber,
			CarType:   "临时车",
		}
		carinfo.Create()
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
						Address:      fmt.Sprintf("%d-%d-%d", i, j, f*100+d),
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
					owner := &owner.Owner{
						Username:  "zsw" + strconv.Itoa(int(house.ID)),
						Password:  "123456",
						Houseid:   int(house.ID),
						Telephone: "18627031721",
						Nickname:  "业主test" + strconv.Itoa(int(house.ID)),
					}
					owner.Create()
					house.Ownerid = int(owner.ID)
					house.Save()

				}
			}
		}
	}
}

func PrepareAll() {
	PrepareHouseAndResident()
	PrepareUserDatas()
	PreparePark()
}
