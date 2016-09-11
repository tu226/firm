package controllers

import "time"

type FirmProduct struct { //每一个设备一个固件对应
	Pid         int       `gorm:"primary_key:pid" json:"pid"` //unique key
	Groupid     string    `gorm:"size:32;index:groupid" json:"groupid"`
	Userid      string    `gorm:"size:32;index:userid" json:"userid"`
	Name        string    `gorm:"size:50" json:"name"`        //产品名称
	Model       string    `gorm:"size:15" json:"model"`       //产品型号
	Cprotocol   int       `gorm:"default:1" json:"cprotocol"` //产品协议
	License     string    `gorm:"size:128" json:"license"`    //license
	Mac         string    `gorm:"size:20" json:"mac"`
	Did         string    `gorm:"size:32;primary_key" json:"did"`   //设备id
	Updatetime  time.Time `gorm:"type:timestamp" json:"updatetime"` // 上一次固件升级时间
	Firmversion string    `gorm:"size:15" json:"firmversion"`       //在用固件版本
	Firmtype    string    `gorm:"size:15" json:"firmtype"`          //固件种类
	Chipid      int       `gorm:"default:0" json:"chipid"`          //芯片的标识
	devtype     string    `gorm:"size:15" json:"devtype"`
}
type DemoappTbl struct { //app表
	Id          int       `gorm:"primary_key:id" json:"id"` //primary key
	Title       string    `gorm:"size:255" json:"title"`
	Provider    string    `gorm:"size:64" json:"provider"`
	System      int       `gorm:"default:1" json:"system"` //1 for android;0 for ios
	Status      bool      `json:"status"`
	Fileid      int       `json:"fileid"`
	Applytime   time.Time `gorm:"type:timestamp" json:"applytime"`
	Username    string    `gorm:"-" json:"username"`
	Versioncode string    `gorm:"size:50" json:"versioncode"`
	Versionname string    `gorm:"size:50" json:"versionname"`
}
type AppDev struct { //app设备关联表
	Userid      string `gorm:"size:32;primary_key" json:"userid"` //userid + pid + did composite primary key
	Pid         int    `json:"pid"`
	Did         string `gorm:"size:32;primary_key" json:"did"`
	Name        string `gorm:"size:50" json:"name"`
	Mac         string `gorm:"size:20" json:"mac"`
	Authid      int    `json:"authid"`
	Authkey     string `gorm:"size:32" json:"authkey"`
	Description string `gorm:"type:text" json:"description"`
	Dpid        string `gorm:"size:50" json:"dpid"`
	Type        int    `json:"type"`
	Password    int    `json:"password"`
	Status      bool   `gorm:"default:false" json:"status"`
}

type Firmware struct {
	Firmversionid   int       `gorm:"primary_key:firmversionid" json:"firmversionid"` //固件版本主键，唯一标识
	Firmversionname string    `gorm:"size:15" json:"firmversion"`
	Firmtype        string    `gorm:"size:15" json:"firmversion"`  //固件类型 固件升降级类型不变
	PermitChip      string    `gorm:"type:text" json:"peimitChip"` //适用的芯片列表
	Checksum        string    `gorm:"size:30" json:"checksum"`
	Svnversion      string    `gorm:"size:15" json:"svnversion"`
	Manuid          string    `gorm:"size:15" json:"manuid"`           //序列号
	Cloudplatform   string    `gorm:"type:text" json:"cloudplatform"`  //适用平台列表
	Sn              string    `gorm:"type:text" json:"cloudplatform"`  //白名单列表
	Info            string    `gorm:"type:text" json:"info"`           //固件描述信息
	Ptime           time.Time `gorm:"type:timestamp" json:"querytime"` //发布时间
	Data            string    `gorm:"type:longblob" json:"data"`       //shuju
}
type ChipInfo struct { //芯片表
	Chipid   int    `gorm:"primary_key:chipid" json:"chipid"` //芯片的主键
	Chipname string `gorm:"size:15" json:"chipname"`
	Chiptype string `gorm:"size:15" json:"Chiptype"`
}
type QueryInfo struct { //请求信息记录Platform
	Id  int    `gorm:"primary_key;auto_increment" json:"id"`
	Did string `gorm:"size:32" json:"did"`
	Pid int    `json:"pid"`
	//authcode string `gorm:"size:128" json:"authcode"` // ""when not exist licens
	Chiptype    string    `gorm:"size:15" json:"Chiptype"`         //芯片类型
	Firmversion int       ` json:"firmversion"`                    //在用固件版本
	Svnversion  string    `gorm:"size:15" json:"svnversion"`       //svn版本
	Manuid      string    `gorm:"size:15" json:"manuid"`           //序列号
	Firmtype    string    `gorm:"size:15" json:"firmversion"`      //固件类型
	Model       string    `gorm:"size:15" json:"model"`            //产品型号
	Sn          string    `gorm:"type:text" json:"Sn"`             //白名单标识
	Querytime   time.Time `gorm:"type:timestamp" json:"querytime"` //请求时间
	devtype     string    `gorm:"size:15" json:"devtype"`
	Reqtype     int       `json:"reqtype"` //0: 自动请求，1 用户触发
}
type DownloadInfo struct {
	Id  int    `gorm:"primary_key;auto_increment" json:"id"`
	Did string `gorm:"size:32" json:"did"`
	Pid int    `json:"pid"`
	//authcode string `gorm:"size:128" json:"authcode"` // ""when not exist licens
	Chiptype    string    `gorm:"size:15" json:"Chiptype"`         //芯片类型
	Firmversion int       ` json:"firmversion"`                    //在用固件版本
	Svnversion  string    `gorm:"size:15" json:"svnversion"`       //svn版本
	Manuid      string    `gorm:"size:15" json:"manuid"`           //序列号
	Firmtype    string    `gorm:"size:15" json:"firmversion"`      //固件类型
	Model       string    `gorm:"size:15" json:"model"`            //产品型号
	Sn          string    `gorm:"type:text" json:"Sn"`             //白名单标识
	Reqtime     time.Time `gorm:"type:timestamp" json:"querytime"` //请求时间
	Devtype     string    `gorm:"size:15" json:"devtype"`
	Reqtype     int       `json:"reqtype"` //0: 自动请求，1 用户触发
	reqversion  int       `json:"reqversion"`
}
