// auto generated by https://github.com/huoyijie/GoalGenerator/releases/tag/v0.0.31
// Please do not change anything in this file.
package country

import (
	"github.com/huoyijie/GoalGenerator/model"
)

type People struct {
	model.Base
	Name     string   `binding:"required,alpha" goal:"<text>globalSearch,filter"`
	Age      uint     `binding:"required" goal:"<number>filter,uint"`
	Identify Identify `binding:"required" goal:"<dropdown>filter,hasOne=country.Identify.NO"`
}

func (*People) Icon() string {
	return "user"
}

func (*People) TranslatePkg() map[string]string {
	t := map[string]string{}
	t["en"] = "Country"
	t["zh-CN"] = "国家"
	return t
}

func (*People) TranslateName() map[string]string {
	t := map[string]string{}
	t["en"] = "People | people"
	t["zh-CN"] = "公民"
	return t
}

func (*People) TranslateFields() map[string]map[string]string {
	return map[string]map[string]string{
		"en": {
			"ID":       "ID",
			"Name":     "Name",
			"Age":      "Age",
			"Identify": "Identify",
		},
		"zh-CN": {
			"ID":       "ID",
			"Name":     "姓名",
			"Age":      "年龄",
			"Identify": "身份证号码",
		},
	}
}

func (m *People) TranslateOptions() map[string]map[string]map[string]string {
	t := map[string]map[string]map[string]string{"en": {}, "zh-CN": {}}

	return t
}
