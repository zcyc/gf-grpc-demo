// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int         `json:"id"       orm:"id"        description:""`
	Passport string      `json:"passport" orm:"passport"  description:""`
	Password string      `json:"password" orm:"password"  description:""`
	Nickname string      `json:"nickname" orm:"nickname"  description:""`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""`
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" description:""`
}
