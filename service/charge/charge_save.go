package charge

import (
	"ProManageSystem/model/charge"
	"ProManageSystem/serializer"
)

type ChargeSaveService struct {
	Houseid int `form:"houseid"`
	//还有多少水
	Water float64 `form:"water"`
	//还有多少度电
	Electric float64 `form:"electric"`
	//还有多少燃气
	Gas float64 `form:"gas"`
	//我的物业费还有多少
	Property float64 `form:"property"`
}
type ChargePayWaterService struct {
	Houseid int `form:"houseid"`
	//费用
	WaterFee float64 `form:"waterfee"`
}
type ChargePayElectricService struct {
	Houseid int `form:"houseid"`
	//费用
	ElectricFee float64 `form:"electricfee"`
}
type ChargePayGasService struct {
	Houseid int `form:"houseid"`
	//费用
	GasFee float64 `form:"gasfee"`
}
type ChargePayPropertyService struct {
	Houseid int `form:"houseid"`
	//费用
	PropertyFee float64 `form:"propertyfee"`
}

func (service *ChargePayWaterService) ChargePayWater() serializer.Response {
	c, err := charge.GetChargebyhouseid(service.Houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	//加上新增的钱的量
	c.Water += service.WaterFee / charge.WaterCountFee
	err = c.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeSave)
	}
	cr := &charge.ChargeRecord{
		Houseid:    service.Houseid,
		Chargetype: "水费",
		Chargefee:  service.WaterFee,
	}
	err = cr.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeRecordCreate)
	}
	return serializer.GetResponse(serializer.Success)
}

func (service *ChargePayElectricService) ChargePayElectric() serializer.Response {
	c, err := charge.GetChargebyhouseid(service.Houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	//加上新增的钱的量
	c.Electric += service.ElectricFee / charge.ElectricCountFee
	err = c.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeSave)
	}
	cr := &charge.ChargeRecord{
		Houseid:    service.Houseid,
		Chargetype: "电费",
		Chargefee:  service.ElectricFee,
	}
	err = cr.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeRecordCreate)
	}
	return serializer.GetResponse(serializer.Success)
}

func (service *ChargePayGasService) ChargePayGas() serializer.Response {
	c, err := charge.GetChargebyhouseid(service.Houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	//加上新增的钱的量
	c.Gas += service.GasFee / charge.GasCountFee
	err = c.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeSave)
	}
	cr := &charge.ChargeRecord{
		Houseid:    service.Houseid,
		Chargetype: "燃气费",
		Chargefee:  service.GasFee,
	}
	err = cr.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeRecordCreate)
	}
	return serializer.GetResponse(serializer.Success)
}
func (service *ChargePayPropertyService) ChargePayProperty() serializer.Response {
	c, err := charge.GetChargebyhouseid(service.Houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	//加上新增的钱的量
	c.Property += service.PropertyFee
	err = c.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeSave)
	}
	cr := &charge.ChargeRecord{
		Houseid:    service.Houseid,
		Chargetype: "物业费",
		Chargefee:  service.PropertyFee,
	}
	err = cr.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeRecordCreate)
	}
	return serializer.GetResponse(serializer.Success)
}

func (service *ChargeSaveService) ChargeSave() serializer.Response {
	charge, err := charge.GetChargebyhouseid(service.Houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	charge.Water = service.Water
	charge.Electric = service.Electric
	charge.Gas = service.Gas
	charge.Property = service.Property
	err = charge.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeSave)
	}
	return serializer.GetResponse(serializer.Success)
}
