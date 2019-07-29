package e

type mErr int 

func (e mErr) Error() string {
	msg, ok := MsgFlags[e]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

const (
	SUCCESS    mErr    = 200
	ERROR      mErr    = 500
	INVALID_PARAMS mErr = 400

	ERROR_EXIST_TAG    mErr     = 10001
	ERROR_NOT_EXIST_TAG   mErr  = 10002
	ERROR_NOT_EXIST_ARTICLE mErr = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL  mErr  = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT  mErr = 20002
	ERROR_AUTH_TOKEN         mErr      = 20003
	ERROR_AUTH                mErr     = 20004
	MISS_AUTH_TOKEN           mErr     = 20005

	ERROR_UPLOAD_SAVE_IMAGE_FAIL mErr = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL mErr= 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT mErr = 30003
)
