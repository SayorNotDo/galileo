package errResponse

import (
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

const ReasonSuccess = "SUCCESS"
const ReasonUnknownError = "UNKNOWN_ERROR"

const ReasonMissingParams = "MISSING_PARAMS"
const ReasonMissingId = "MISSING_ID"
const ReasonParamsError = "PARAMS_ERROR"
const TimeFormatError = "TIME_FORMAT_ERROR"
const ReasonRecordNotFound = "RECORD_NOT_FOUND"
const ReasonRecordExists = "RECORD_EXISTS"

const ReasonUserNotFound = "USER_NOT_FOUND"
const ReasonUserPasswordError = "USER_PASSWORD_ERROR"
const ReasonUserForbidden = "USER_FORBIDDEN"
const ReasonUserDeleted = "USER_DELETED"
const ReasonUserUsernameExist = "USER_USERNAME_EXIST"
const ReasonUserPhoneExist = "USER_PHONE_EXIST"
const ReasonUnauthorizedUser = "UNAUTHORIZED_USER"

const ReasonAdministratorNotFound = "ADMINISTRATOR_NOT_FOUND"
const ReasonAdministratorPasswordError = "ADMINISTRATOR_PASSWORD_ERROR"
const ReasonAdministratorForbidden = "ADMINISTRATOR_FORBIDDEN"
const ReasonAdministratorDeleted = "ADMINISTRATOR_DELETED"
const ReasonAdministratorUsernameExist = "ADMINISTRATOR_USERNAME_EXIST"
const ReasonAdministratorMobileExist = "ADMINISTRATOR_MOBILE_EXIST"

const ReasonUnauthorizedRole = "UNAUTHORIZED_ROLE"
const ReasonUnauthorizedInfoMissing = "UNAUTHORIZED_INFO_MISSING"
const ReasonAuthorizedDataMissing = "AUTHORIZATION_DATA_MISSING"
const ReasonAuthorizationRoleExist = "AUTHORIZATION_ROLE_EXIST"
const ReasonAuthorizationRoleNotFound = "AUTHORIZATION_ROLE_NOT_FOUND"
const ReasonAuthorizationUserHasRoleAlready = "AUTHORIZATION_USE_HAS_ROLE_ALREADY"

const ReasonAuthorizationApiNotFound = "AUTHORIZATION_API_NOT_FOUND"
const ReasonAuthorizationApiExist = "AUTHORIZATION_API_EXIST"

const ReasonFileNameMissing = "FILE_NAME_MISSING"
const ReasonFileMissing = "FILE_MISSING"
const ReasonFileOverLimitSize = "FILE_OVER_LIMIT_SIZE"
const ReasonOssConfigWrong = "OSS_CONFIG_WRONG"
const ReasonOssBucketWrong = "OSS_BUCKET_WRONG"
const ReasonOssPutObjectFail = "OSS_PUT_OBJECT_FILE"

const ReasonUserUnauthorized = "UNAUTHORIZED"

const ReasonSystemError = "SYSTEM_ERROR"
const ReasonServiceGatewayTimeout = "SERVICE_GATEWAY_TIMEOUT"
const ReasonServiceUnavailable = "SERVICE_GATEWAY_UNAVAILABLE"

var reasonMessageAll = map[string]string{
	ReasonUserNotFound:      "user not found",
	ReasonUserPasswordError: "user password error",
	ReasonUserForbidden:     "user forbidden",
	ReasonUserDeleted:       "user has been deleted",
	ReasonUserUsernameExist: "username is already exist",
	ReasonUserPhoneExist:    "phone is already exist",
	ReasonUnauthorizedUser:  "user is unauthorized",

	ReasonSuccess:      "success",
	ReasonUnknownError: "unknown error",

	ReasonParamsError:    "params error",
	ReasonMissingParams:  "params missing",
	TimeFormatError:      "时间格式错误",
	ReasonMissingId:      "id不得为空",
	ReasonRecordNotFound: "数据不存在",
	ReasonRecordExists:   "record already exists",

	ReasonAdministratorNotFound:      "管理员数据不存在",
	ReasonAdministratorPasswordError: "管理员密码错误",
	ReasonAdministratorForbidden:     "管理员已禁用",
	ReasonAdministratorDeleted:       "管理员已删除",
	ReasonAdministratorUsernameExist: "管理员用户名已存在",
	ReasonAdministratorMobileExist:   "管理员手机号已存在",

	ReasonUserUnauthorized:        "user is unauthorized",
	ReasonUnauthorizedInfoMissing: "角色授权信息不存在",
	ReasonAuthorizedDataMissing:   "权限数据不存在",
	ReasonUnauthorizedRole:        "角色未授权",

	ReasonAuthorizationRoleExist:          "角色已存在",
	ReasonAuthorizationRoleNotFound:       "角色不存在",
	ReasonAuthorizationUserHasRoleAlready: "用户已经拥有角色",

	ReasonAuthorizationApiNotFound: "API不存在",
	ReasonAuthorizationApiExist:    "API已存在",

	ReasonFileNameMissing:   "文件名称不得为空",
	ReasonFileMissing:       "upload file is missing",
	ReasonFileOverLimitSize: "文件超过最大限制",

	ReasonOssConfigWrong:   "OSS配置错误",
	ReasonOssBucketWrong:   "OSS Bucket配置错误",
	ReasonOssPutObjectFail: "OSS文件上传失败",

	ReasonSystemError:           "系统繁忙,请稍后再试",
	ReasonServiceGatewayTimeout: "服务不可达",
	ReasonServiceUnavailable:    "服务不可达",
}

var reasonCodeAll = map[string]int{
	ReasonSuccess:      0,
	ReasonUnknownError: 1,

	ReasonParamsError:    4000,
	ReasonMissingParams:  1001,
	ReasonMissingId:      1002,
	ReasonRecordNotFound: 4004,
	ReasonRecordExists:   4005,
	TimeFormatError:      1004,

	ReasonAdministratorNotFound:      2001,
	ReasonAdministratorPasswordError: 2002,
	ReasonAdministratorForbidden:     2003,
	ReasonAdministratorDeleted:       2004,
	ReasonAdministratorUsernameExist: 2005,
	ReasonAdministratorMobileExist:   2006,

	ReasonUserNotFound:      4000,
	ReasonUserPasswordError: 4001,
	ReasonUserForbidden:     4003,
	ReasonUserDeleted:       4004,
	ReasonUserUsernameExist: 2002,
	ReasonUserPhoneExist:    2002,

	ReasonUserUnauthorized:                4000,
	ReasonUnauthorizedRole:                4001,
	ReasonUnauthorizedInfoMissing:         4002,
	ReasonAuthorizedDataMissing:           4003,
	ReasonAuthorizationRoleExist:          4020,
	ReasonAuthorizationRoleNotFound:       4024,
	ReasonAuthorizationUserHasRoleAlready: 4022,

	ReasonAuthorizationApiExist:    4032,
	ReasonAuthorizationApiNotFound: 4034,

	ReasonFileNameMissing:   4041,
	ReasonFileMissing:       4042,
	ReasonFileOverLimitSize: 4043,
	ReasonOssConfigWrong:    4044,
	ReasonOssBucketWrong:    4045,
	ReasonOssPutObjectFail:  4046,

	ReasonSystemError:           5000,
	ReasonServiceGatewayTimeout: 5004,
	ReasonServiceUnavailable:    5003,
}

var reasonGrpcCodeAll = map[string]int{
	ReasonSuccess:      http.StatusOK,
	ReasonUnknownError: http.StatusInternalServerError,

	ReasonParamsError:   http.StatusBadRequest,
	ReasonMissingParams: http.StatusBadRequest,
	ReasonMissingId:     http.StatusBadRequest,

	ReasonAdministratorNotFound:      http.StatusBadRequest,
	ReasonAdministratorPasswordError: http.StatusBadRequest,
	ReasonAdministratorForbidden:     http.StatusBadRequest,
	ReasonAdministratorDeleted:       http.StatusBadRequest,
	ReasonAdministratorUsernameExist: http.StatusBadRequest,
	ReasonAdministratorMobileExist:   http.StatusBadRequest,

	ReasonUserUnauthorized: http.StatusUnauthorized,

	ReasonSystemError:           http.StatusInternalServerError,
	ReasonServiceGatewayTimeout: http.StatusGatewayTimeout,
}

// SetCustomizeErrInfo 根据err.Reason返回自定义包装错误
func SetCustomizeErrInfo(err error) error {
	e := errors.FromError(err)
	// 如果 e.Code = 504 则是服务不可达
	if e.Code == http.StatusGatewayTimeout {
		return SetCustomizeErrInfoByReason(ReasonServiceGatewayTimeout)
	}
	// 如果 e.Code = 503 则是服务连接被拒绝
	if e.Code == http.StatusServiceUnavailable {
		return SetCustomizeErrInfoByReason(ReasonServiceUnavailable)
	}
	reason := e.Reason
	if reason == "" {
		reason = ReasonUnknownError
	}
	if _, ok := reasonCodeAll[reason]; !ok {
		return err
	}
	// 如果是参数错误， 则检查err是否有值， 有则直接参会
	if e.Reason == ReasonParamsError && e.Message != "" {
		return errors.New(reasonCodeAll[reason], reason, e.Message)
	}
	return SetCustomizeErrInfoByReason(e.Reason)
}

// SetCustomizeErrInfoByReason 根据err.Reason返回自定义包装错误
func SetCustomizeErrInfoByReason(reason string) error {
	code, message := reasonCodeAll[reason], reasonMessageAll[reason]
	return errors.New(code, reason, message)
}

// SetCustomizeErrMsg 根据err.Reason返回自定义包装错误（message动态修改）
func SetCustomizeErrMsg(reason, msg string) error {
	code := reasonCodeAll[reason]
	return errors.New(code, reason, msg)
}

// GetErrInfoByReason 根据err.Reason返回自定义包装错误
func GetErrInfoByReason(reason string) string {
	return reasonMessageAll[reason]
}

// SetErrByReason 根据err.Reason返回自定义包装错误
func SetErrByReason(reason string) error {
	code, message := reasonGrpcCodeAll[reason], reasonMessageAll[reason]
	return errors.New(code, reason, message)
}
