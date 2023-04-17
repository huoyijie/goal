// auto generated by https://github.com/huoyijie/GoalGenerator/releases/tag/v0.0.23
// Please do not change anything in this file.
package auth

import (
    "github.com/huoyijie/GoalGenerator/model"
)

type User struct {

    model.Base

    Username string `gorm:"unique" binding:"required,alphanum,min=3,max=40" goal:"<text>unique,sortable,globalSearch,filter"`
    Email string `gorm:"unique" binding:"required,email" goal:"<text>unique,sortable,filter"`
    Password string `binding:"required,min=8" goal:"<password>hidden,secret"`
    IsSuperuser bool `goal:"<switch>readonly,filter"`
    IsActive bool `goal:"<switch>filter"`
}
func (User) TableName() string {
    return "auth_users"
}

func (*User) Icon() string {
    return "users"
}
func (*User) Lazy() {}

func (*User) TranslatePkg() map[string]string {
    t := map[string]string{}
    t["en"] = "Auth"
    t["zh_CN"] = "认证授权"
    return t
}

func (*User) TranslateName() map[string]string {
    t := map[string]string{}
    t["en"] = "User | users"
    t["zh_CN"] = "用户"
    return t
}

func (*User) TranslateFields() map[string]map[string]string {
    return map[string]map[string]string{
        "en": {

            "ID": "ID",

            "Username": "Username",

            "Email": "Email",

            "Password": "Password",

            "IsSuperuser": "Super User",

            "IsActive": "Active",
        },
        "zh_CN": {

            "ID": "ID",

            "Username": "用户名",

            "Email": "邮件地址",

            "Password": "密码",

            "IsSuperuser": "超级管理员",

            "IsActive": "有效",
        },
    }
}
