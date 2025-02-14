package types

const (
	SUCCESS_CODE          = 20000      //成功的状态码
	FAIL_CODE             = 30000      //失败的状态码
	MD5_PREFIX            = "jkfldfsf" //MD5加密前缀字符串
	TOKEN_KEY             = "X-Token"  //页面token键名
	USER_ID_Key           = "X-USERID" //页面用户ID键名
	USER_UUID_Key         = "X-UUID"   //页面UUID键名
	SUPER_ADMIN_ID uint64 = 956986     // 超级管理员账号ID
)

type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseModelBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponsePageData struct {
	Total uint64      `json:"total"`
	Items interface{} `json:"items"`
}

type ResponsePage struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    ResponsePageData `json:"data"`
}
