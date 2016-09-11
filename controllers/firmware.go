package controllers

import (
	"time"

	"github.com/astaxie/beego"
	js "github.com/go-simplejson"
)

type FirmController struct {
	beego.Controller
}

func (this *FirmController) FindNewVer() {
	var data = map[string]interface{}{
		"code": -2,
		"msg":  "fail",
	}
	defer func() {
		this.Data["json"] = data
		this.ServeJson()
	}()
	obj, err := js.NewJson(this.Ctx.Input.CopyBody())
	if err != nil {
		beego.Warn("js.NewJson:", err)
		data["code"] = RequestParse
		data["msg"] = codemsg[RequestParse]
		return
	}
	//did := obj.Get("did").MustString()
	//pid := obj.Get("pid").MustInt()
	//暂时做身份判断 authcode:=obj.Get("authcode").MustString()

	var pl QueryInfo
	pl.Pid = obj.Get("pid").MustInt()
	pl.Did = obj.Get("reqcode").Get("did").MustString()
	if pl.Did == "" {
		beego.Warn("pl.Did invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	} //
	pl.Chiptype = obj.Get("reqcode").Get("chiptype").MustString()
	if pl.Chiptype == "" {
		beego.Warn("pl.Chiptype invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	pl.Firmtype = obj.Get("reqcode").Get("firmtype").MustString()
	if pl.Firmtype == "" {
		beego.Warn("pl.Firmtype invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	pl.Firmversion = obj.Get("reqcode").Get("firmversion").MustInt()
	if pl.Firmversion == 0 {
		beego.Warn("pl.Firmversion invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	pl.Manuid = obj.Get("reqcode").Get("manuid").MustString()
	if pl.Manuid == "" {
		beego.Warn("pl.Manuid invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	pl.Model = obj.Get("reqcode").Get("model").MustString()
	if pl.Model == "" {
		beego.Warn("pl.Model invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	pl.Querytime = time.Now()
	pl.Sn = obj.Get("reqcode").Get("sn").MustString() //设备所在sn分组
	pl.Svnversion = obj.Get("reqcode").Get("Svnversion").MustString()
	if pl.Svnversion == "" {
		beego.Warn("Svnversion invaid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	pl.devtype = obj.Get("reqcode").Get("devtype").MustString()
	if pl.devtype == "" {
		beego.Warn("pl.devtype invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	pl.Reqtype = obj.Get("reqcode").Get("reqtype").MustInt()
	tx := dbpool.db.Begin()
	err = tx.Create(&pl).Error
	if err != nil {
		beego.Warn("tx.Create:", err)
		tx.Rollback()
		return
	}
	tx.Commit()
	//查询Chiptype兼容的且符合白名单的firmware所有版本的list
	var firmlist []Firmware
	var firmlistNum int
	err = dbpool.db.Where("permit_chip like ? and (sn like ? or sn is null)", "%"+pl.Chiptype+"%", "%"+pl.Sn+"%").Order("ptime").Find(&firmlist).Count(&firmlistNum).Error
	if err != nil {
		beego.Warn("find firmlist:", err)
		return
	}
	beego.Debug(firmlist, firmlistNum)
	var outList = make([]map[string]interface{}, 0, firmlistNum)
	for _, v := range firmlist {
		outList = append(outList, map[string]interface{}{
			"firmversionid": v.Firmversionid,
			"info":          v.Info,
		})
	}
	data["code"] = 0
	data["msg"] = outList

}
func (this *FirmController) NewVer() {
	var data = map[string]interface{}{
		"code": ServerBusy,
		"msg":  codemsg[ServerBusy],
	}
	defer func() {
		this.Data["json"] = data
		this.ServeJson()
	}()
	fw := Firmware{}
	//	if err := this.ParseForm(&fw); err != nil {
	//		beego.Warn("get fw worry")
	//		data["code"] = RequestParams
	//		data["msg"] = codemsg[RequestParams]
	//	}
	//	beego.Debug("fw")
	obj, err := js.NewJson(this.Ctx.Input.CopyBody())
	if err != nil {
		beego.Warn("js.NewJson:", err)
		data["code"] = RequestParse
		data["msg"] = codemsg[RequestParse]
		return
	}
	fw.Firmversionid = obj.Get("Firmversionid").MustInt()
	var num int
	dbpool.db.Where("Firmversionid= ? ", fw.Firmversionid).Model(fw).Count(&num)
	if num == 1 {
		data["code"] = RequestParse
		data["msg"] = "该固件id已经存在"
		return
	}
	fw.Firmversionname = obj.Get("Firmversionname").MustString()
	if fw.Firmversionname == "" {
		beego.Warn("Firmversionname invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	fw.Firmtype = obj.Get("Firmtype").MustString()
	if fw.Firmtype == "" {
		beego.Warn("fw.Firmtype invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	fw.PermitChip = obj.Get("PermitChip").MustString()
	if fw.PermitChip == "" {
		beego.Warn("fw.PermitChip invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	fw.Checksum = obj.Get("Checksum").MustString()
	if fw.Checksum == "" {
		beego.Warn("fw.Checksum invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	fw.Svnversion = obj.Get("Svnversion").MustString()
	if fw.Svnversion == "" {
		beego.Warn("fw.Svnversion invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	fw.Manuid = obj.Get("Manuid").MustString()
	if fw.Manuid == "" {
		beego.Warn("fw.Manuid invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	fw.Cloudplatform = obj.Get("Cloudplatform").MustString()
	if fw.Cloudplatform == "" {
		beego.Warn("fw.Cloudplatform invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	fw.Sn = obj.Get("Sn").MustString()
	fw.Info = obj.Get("Info").MustString()
	if fw.Info == "" {
		beego.Warn("fw.Info invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	fw.Data = obj.Get("data").MustString()
	if fw.Data == "" {
		beego.Warn("fw.Data invalid")
		data["code"] = RequestParams
		data["msg"] = codemsg[RequestParams]
		return
	}
	beego.Debug(fw)

	beego.Debug("dbpool.db:", dbpool.db)

	tx := dbpool.db.Begin()
	err = tx.Create(&fw).Error
	if err != nil {
		beego.Warn("tx.Create:", err)
		tx.Rollback()
		return
	}
	tx.Commit()
	data["code"] = Success
	data["msg"] = codemsg[Success]

}
func (this *FirmController) DownloadFirm() {
	var data = map[string]interface{}{
		"code": ServerBusy,
		"msg":  codemsg[ServerBusy],
	}
	defer func() {
		this.Data["json"] = data
		this.ServeJson()
	}()
	obj, err := js.NewJson(this.Ctx.Input.CopyBody())
	if err != nil {
		beego.Warn("js.NewJson:", err)
		data["code"] = RequestParse
		data["msg"] = codemsg[RequestParse]
		return
	}
	var fw Firmware

	fw.Firmversionid = obj.Get("reqcode").Get("reqversion").MustInt()
	var dinfo DownloadInfo
	dinfo.Chiptype = obj.Get("reqcode").Get("chiptype").MustString()
	dinfo.Devtype = obj.Get("reqcode").Get("devtype").MustString()
	dinfo.Did = obj.Get("reqcode").Get("did").MustString()
	dinfo.Firmtype = obj.Get("reqcode").Get("firmtype").MustString()
	dinfo.Firmversion = obj.Get("reqcode").Get("firmversion").MustInt()
	dinfo.Manuid = obj.Get("reqcode").Get("manuid").MustString()
	dinfo.Model = obj.Get("reqcode").Get("model").MustString()
	dinfo.Pid = obj.Get("pid").MustInt()
	dinfo.Reqtime = time.Now()
	dinfo.Reqtype = obj.Get("reqcode").Get("reqtype").MustInt()
	dinfo.reqversion = obj.Get("reqversion").MustInt()
	dinfo.Sn = obj.Get("reqcode").Get("sn").MustString()
	dinfo.Svnversion = obj.Get("reqcode").Get("svnversion").MustString()
	tx := dbpool.db.Begin()
	err = tx.Create(&dinfo).Error
	if err != nil {
		beego.Warn("create dinfo :", err)
		tx.Rollback()
		return
	}
	tx.Commit()
	if dbpool.db.Where("firmversionid= ? and (sn like ? or sn is null)", fw.Firmversionid, "%"+dinfo.Sn+"%").Find(&fw).RecordNotFound() == true {
		data["code"] = -1
		data["msg"] = "没有此版本可供下载"
		return
	}
	data["code"] = 0
	data["msg"] = fw.Data

}
