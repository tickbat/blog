package e

// Response {code int msg string} 不能为常量
type Response int

func (e Response) Error() string {
	val, ok := msg[e]
	if ok {
		return val
	}
	return msg[ERROR]
}

func (e Response) Code() int {
	return int(e)
}

const (
	SUCCESS        Response = 200
	ERROR          Response = 500
	INVALID_PARAMS Response = 400

	ERROR_EXIST_TAG         Response = 10001
	ERROR_NOT_EXIST_TAG     Response = 10002
	ERROR_NOT_EXIST_ARTICLE Response = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    Response = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT Response = 20002
	ERROR_AUTH_TOKEN               Response = 20003
	ERROR_AUTH                     Response = 20004
	MISS_AUTH_TOKEN                Response = 20005

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    Response = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   Response = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT Response = 30003
)
