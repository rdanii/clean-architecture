package response

import (
	"net/http"
	"strings"
)

func BuildSuccessResponse(data interface{}) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusOK,
		ErrorCode:       &ERROR_CODE_SUCCESS,
		ResponseCode:    &RESPONSE_CODE_SUCCESS,
		ResponseMessage: &RESPONSE_MESSAGE_SUCCESS,
		Errors:          nil,
		Data:            data,
		Info:            nil,
	}
}

func BuildSuccessResponseWithInfo(data interface{}, info *ResponseInfo) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusOK,
		ErrorCode:       &ERROR_CODE_SUCCESS,
		ResponseCode:    &RESPONSE_CODE_SUCCESS,
		ResponseMessage: &RESPONSE_MESSAGE_SUCCESS,
		Errors:          nil,
		Data:            data,
		Info:            info,
	}
}

func BuildDataNotFoundResponse() *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusNotFound,
		ErrorCode:       &ERROR_CODE_DATA_NOT_FOUND,
		ResponseCode:    &RESPONSE_CODE_NOT_FOUND,
		ResponseMessage: &RESPONSE_MESSAGE_DATA_NOT_FOUND,
		Errors:          nil,
		Data:            nil,
		Info:            nil,
	}
}

func BuildDataNotFoundResponseWithMessage(msg string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusNotFound,
		ErrorCode:       &ERROR_CODE_DATA_NOT_FOUND,
		ResponseCode:    &RESPONSE_CODE_NOT_FOUND,
		ResponseMessage: &RESPONSE_MESSAGE_DATA_NOT_FOUND,
		Errors:          strings.Split(msg, "\n"),
		Data:            nil,
		Info:            nil,
	}
}

func BuildBadRequestResponse(errCode, respCode, errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusBadRequest,
		ErrorCode:       &errCode,
		ResponseCode:    &respCode,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
		Info:            nil,
	}
}

func BuildInternalErrorResponse(errCode, respCode, errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusInternalServerError,
		ErrorCode:       &errCode,
		ResponseCode:    &respCode,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
		Info:            nil,
	}
}

func BuildRouteNotFoundResponse() *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusNotFound,
		ErrorCode:       &ERROR_CODE_DATA_NOT_FOUND,
		ResponseCode:    &RESPONSE_CODE_NOT_FOUND,
		ResponseMessage: &RESPONSE_MESSAGE_ROUTE_NOT_FOUND,
		Errors:          nil,
		Data:            nil,
		Info:            nil,
	}
}

func BuildEmptyBodyReqResponse(errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusBadRequest,
		ErrorCode:       &ERROR_CODE_BODY_REQUEST_EMPTY,
		ResponseCode:    &RESPONSE_CODE_BAD_REQUEST,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}

func BuildInvalidTypeResponse(errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusBadRequest,
		ErrorCode:       &ERROR_CODE_INVALID_DATA_TYPE,
		ResponseCode:    &RESPONSE_CODE_BAD_REQUEST,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}

func BuildUnauthorizedResponse(errMessage, throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusUnauthorized,
		ErrorCode:       &ERROR_CODE_INVALID_AUTH_TOKEN,
		ResponseCode:    &RESPONSE_CODE_AUTH_ERROR,
		ResponseMessage: &errMessage,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}

func BuildTimeoutResponse(throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusRequestTimeout,
		ErrorCode:       &ERROR_CODE_TIMEOUT,
		ResponseCode:    &RESPONSE_CODE_INTERNAL_ERROR,
		ResponseMessage: &RESPONSE_MESSAGE_TIMEOUT,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}

func BuildForbiddenAccessResponse(throwable string) *ResponseContainer {
	return &ResponseContainer{
		StatusCode:      http.StatusForbidden,
		ErrorCode:       &ERROR_CODE_FORBIDDEN_ACCESS,
		ResponseCode:    &RESPONSE_CODE_AUTH_ERROR,
		ResponseMessage: &RESPONSE_MESSAGE_FORBIDDEN_ACCESS,
		Errors:          strings.Split(throwable, "\n"),
		Data:            nil,
	}
}
