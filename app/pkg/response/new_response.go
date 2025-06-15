package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message    string      `json:"message,omitempty"`
	Code       int         `json:"code,omitempty"`
	StatusBool bool        `json:"status_bool,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	ErrorCode  string      `json:"error_code,omitempty"`
	Timestamp  string      `json:"timestamp,omitempty"`
	Path       string      `json:"path,omitempty"`
	RequestID  string      `json:"request_id,omitempty"`
	Total      int64       `json:"total_row,omitempty"`
	View       echo.Map    `json:"view,omitempty"`
}

type builderResponse struct {
	messageX    string
	codeX       int
	dataX       interface{}
	statusBoolx bool
	totalX      int64
	reqIDx      string
	viewX       echo.Map
}

func NewResponseBuilder() *builderResponse {
	return &builderResponse{}
}

func (r *builderResponse) Message(message string) *builderResponse {
	r.messageX = message
	return r
}

func (r *builderResponse) Code(code int) *builderResponse {
	r.codeX = code
	return r
}

func (r *builderResponse) Data(data interface{}) *builderResponse {
	r.dataX = data
	return r
}

func (r *builderResponse) RequestID(data string) *builderResponse {
	r.reqIDx = data
	return r
}

func (r *builderResponse) Bool(statusbool bool) *builderResponse {
	r.statusBoolx = statusbool
	return r
}

func (r *builderResponse) Total(total int64) *builderResponse {
	r.totalX = total
	return r
}

func (r *builderResponse) Build() *Response {
	return &Response{
		Message:    r.messageX,
		Code:       r.codeX,
		Data:       r.dataX,
		RequestID:  r.reqIDx,
		StatusBool: r.statusBoolx,
		Total:      r.totalX,
		View:       r.viewX,
	}
}
