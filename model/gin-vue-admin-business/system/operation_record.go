package model

import (
	"github.com/flipped-aurora/gva/library/global"
	"time"
)

type OperationRecord struct {
	global.Model
	Ip           string `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`
	Path         string `json:"path" form:"path" gorm:"column:path;comment:请求路径"`
	Agent        string `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`
	Method       string `json:"method" form:"method" gorm:"column:method;comment:请求方法"`
	Request      string `json:"body" form:"body" gorm:"type:text;column:request;comment:请求Body"`
	Response     string `json:"resp" form:"resp" gorm:"type:text;column:response;comment:响应Body"`
	ErrorMessage string `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`

	Status int `json:"status" form:"status" gorm:"column:status;comment:请求状态"`
	UserID int `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`

	Latency time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟"`

	User User `json:"user"`
}

func (o *OperationRecord) TableName() string {
	return "system_operation_records"
}
