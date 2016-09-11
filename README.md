#OTA升级是目前所有IOT产品的基本功能，所有产品必须具备的功能。固件升级由于涉及硬件的变更，所以是一件风险很大的事情，新版本必须经过严格验证和确认才能全面推广到所有用户。

针对这些情况，固件升级需要使用灰度发布的机制。服务器发布固件时，先在部分用户中提示，验证通过后，在全面放开。整体发布流程如下：
 

详细需求：
1.	固件编译完成后，需要输出（固件头和固件名称）如下信息
-	芯片类型
-	平台
-	产品
-	版本
-	校验和

2.	上线发布
有管理员确认新版本固件的发布目标（SN，目前版本，平台等）

3.	固件在启动或者用户查询新版本时携带如下信息
-	芯片类型
-	固件版本
-	SVN版本
-	设备类型
-	第三方平台信息
-	查询触发方式（自动查询，用户触发）

4.	升级时做如下校验事宜
-	检查芯片是否兼容
-	校验镜像是否完整

5.	云端在接收到固件升级相关请求时，需要做如下事宜
-	纪录请求的公网IP地址
-	纪录固件上报过来的信息
-	做出响应

详细设计：
1.	固件版本查询
服务器地址： DeviceTypeimage.broadlink.com.cn/firmware/version
方法： POST
参数：
{
		“did”: “xxxx”, //没有时填mac，格式”112233445566” ？
		“pid”: “xxxx”, //xx
		“authcode”: “xxx”, //没有License时填空 
		“reqcode”:”xxx” //请求数据

}
请求数据的格式如下：
Base64(AES(platform)),
如果固件有license，则aes的key使用license库生成，否则key和iv分别用：
Md5(did), md5(pid).
Platform的格式如下：
{
				“bl”:{},
				“hi”:{},
				“u+”{},
				“jd”:{},
				“wx”:{}
“reqtype”: 0|1 //0: 自动请求，1 用户触发
}
其中bl里面包括如下字段：
“chip”: “xxx”, //芯片类型
“fwversion”: “20024”, //固件版本 ？
“revision”: “xxxx”, //svn版本
Hi里面包括如下字段：
“version”:””,
“did”:””,
“manuid”:””,
“devtype”:”xxx”,
“model”:””,
“sn”:””,
“st”:””
请求返回：
返回数据	意义
{“error”:”0”,”version”:”xxx”}	成功
{“error”:”-2”,”version”:”0”}	失败

2.	固件升级下载请求
服务器地址： DeviceTypeimage.broadlink.com.cn/firmware/download
方法： POST
参数：
{
		“did”: “xxxx”, //没有时填mac，格式”112233445566”
		“pid”: “xxxx”, //xx
		“authcode”: “xxx”, //没有License时填空
		“reqcode”:”xxx” //请求数据
}
请求数据的格式如下：
Base64(AES(platform)),
如果固件有license，则aes的key使用license库生成，否则key和iv分别用：
Md5(did), md5(pid).
Platform的格式如下：
{
				“bl”:{},
				“hi”:{},
				“u+”{},
				“jd”:{},
				“wx”:{}
“reqtype”: 0|1 //0: 自动请求，1 用户触发
“reqversion”: “”, 请求下载的版本，没有为空
}
其中bl里面包括如下字段：
“chip”: “xxx”, //芯片类型
“fwversion”: “20024”, //固件版本
“revision”: “xxxx”, //svn版本
Hi里面包括如下字段：
“version”:””,
“did”:””,
“manuid”:””,
“devtype”:”xxx”,
“model”:””,
“sn”:””,
“st”:””

返回镜像文件数据，同时在响应的头中设置签名：
IMGSIG: “”, 
签名计算方法为使用License库中的签名函数对镜像的checksum进行签名，固件收到后，校验签名通过后，才继续升级。

所有HTTP请求，在HTTP 的code为200时，才需要继续解析数据。

3.	服务器发布固件时，需要添加固件可见白名单，针对华为平台，白名单列表为SN，没有设置白名单表示全部可见。

添加发布固件接口
接口地址：/firmware/publish
请求方法post
请求参数{
			“Firmversionid”:10003    //固件版本主键，唯一标识
			“Firmversionname“:”xxxx”
			“Firmtype”:”A1”          	//固件类型 固件升降级类型不变（芯片类型）
			“PermitChip”:	“A1,A2”			//适用的芯片列表     
			“Checksum”:”XXXXX”        //校验和
	 	     “Svnversion“:”4562”          //svn版本号
			“Manuid“:”XXXXXXXX”          //  序列号
			“Cloudplatform”:”XXXX,XXXX,XXX” //适用平台列表
				“Sn”:”XXXX,XXXX,XXX”         //白名单列表
			}
返回参数{
Code：
Msg:
}
