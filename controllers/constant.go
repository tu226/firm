package controllers

const (
	Success           = 0
	RequestMethod     = 400
	RequestParse      = 401
	RespondParse      = 402
	RequestParams     = 403
	RespondParams     = 404
	ReadRecord        = 405
	WriteRecord       = 406
	UpdateRecord      = 407
	NetRequest        = 408
	NetRespond        = 409
	NetRead           = 410
	NetWrite          = 411
	NetFile           = 412
	NetInteractive    = 413
	ProtocolConvert   = 414
	LoginStatus       = 415
	EncBizSign        = 416
	DecBizSign        = 417
	PicVerifyCode     = 418
	RegisterStatus    = 419
	PhoneCodeStatus   = 420
	ServerBusy        = 421
	GeneOfficialPid   = 422
	ExistOfficialPid  = 423
	ExistPublished    = 424
	NotExistUser      = 425
	NotAdmin          = 426
	NotExistGroup     = 427
	ExistFirm         = 428
	MappingNotMatch   = 429
	ItemExist         = 430
	NotAuditLicense   = 431
	PendAuditRelation = 432
	UserNotInGroup    = 433
	UserInviteSelf    = 434
	UserIsAdmin       = 435
	ProNotInGroup     = 436
	CodeExist         = 437
	EmailNoVerify     = 438
	HadRegistered     = 439
)

const (
	DEVELOPSTATUS        string = "develop"
	INNERFAILURESTATUS   string = "failureinneraudit"
	INNERSTATUS          string = "inneraudit"
	BLAUDITSTATUS        string = "blaudit"
	PUBLISHSTATUS        string = "published"
	PUBLISHFAILURESTATUS string = "failurepublished"

	DO   byte = 0
	UNDO byte = 1

	PUBLISHHISTORYFLAG byte = 1
)

var codemsg = map[int]string{
	ServerBusy:        "server busy",
	Success:           "success",
	RequestMethod:     "request method error",
	RequestParse:      "request params parse error",
	RespondParse:      "respond params parse error",
	RequestParams:     "request params invalid",
	RespondParams:     "respond params invalid",
	ReadRecord:        "no datarecord",
	WriteRecord:       "write datarecord error",
	UpdateRecord:      "update datarecord error",
	NetRequest:        "net request error",
	NetRespond:        "net respond error",
	NetRead:           "net read error",
	NetWrite:          "net write error",
	NetFile:           "net file error",
	NetInteractive:    "network interactive error",
	ProtocolConvert:   "protocol convert error",
	LoginStatus:       "login error",
	EncBizSign:        "encode bizsign error",
	DecBizSign:        "decode bizsign error",
	PicVerifyCode:     "verify code error",
	RegisterStatus:    "register error",
	PhoneCodeStatus:   "phone verify code error",
	GeneOfficialPid:   "generate officialpid error",
	ExistOfficialPid:  "officialpid exist",
	ExistPublished:    "product had published",
	NotExistUser:      "user existn't",
	NotAdmin:          "Not Admin",
	NotExistGroup:     "Group existn't",
	ExistFirm:         "The Firm Already exist",
	MappingNotMatch:   "mapping not match",
	ItemExist:         "Item Exist",
	NotAuditLicense:   "not audit license",
	PendAuditRelation: "user pending audit",
	UserNotInGroup:    "User Not In Group",
	UserInviteSelf:    "user don't invite self",
	UserIsAdmin:       "current user is admin",
	ProNotInGroup:     "product not belong group",
	CodeExist:         "code exist",
	EmailNoVerify:     "email no verfication",
	HadRegistered:     "user had registered",
}
