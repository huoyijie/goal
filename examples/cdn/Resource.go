// auto generated by https://github.com/huoyijie/GoalGenerator/releases/tag/v0.0.23
// Please do not change anything in this file.
package cdn

import (
    "github.com/huoyijie/GoalGenerator/model"
)

type Resource struct {

    model.Base

    File string `gorm:"unique" binding:"required" goal:"<file>unique,postonly,globalSearch,filter,uploadTo=uploads"`
    Status string `binding:"required" goal:"<dropdown>filter,strings"`
    Level uint `binding:"required" goal:"<dropdown>filter,uints"`
}
func (*Resource) StatusStrings() []string {
    return []string{"tbd", "on", "off"}
}
func (*Resource) LevelUints() []uint {
    return []uint{1, 2, 3}
}

func (*Resource) Icon() string {
    return "cloud-upload"
}
func (*Resource) Purge() {}

func (*Resource) Lazy() {}

func (*Resource) TranslatePkg() map[string]string {
    t := map[string]string{}
    t["en"] = "CDN"
    t["zh_CN"] = "CDN"
    return t
}

func (*Resource) TranslateName() map[string]string {
    t := map[string]string{}
    t["en"] = "Resource | resources"
    t["zh_CN"] = "资源"
    return t
}

func (*Resource) TranslateFields() map[string]map[string]string {
    return map[string]map[string]string{
        "en": {

            "ID": "ID",

            "File": "File",

            "Status": "Status",

            "Level": "Level",
        },
        "zh_CN": {

            "ID": "ID",

            "File": "文件",

            "Status": "状态",

            "Level": "等级",
        },
    }
}
