// auto generated by https://github.com/huoyijie/GoalGenerator/releases/tag/v0.0.12
// Please do not change anything in this file.
package cdn

type Resource struct { 

    ID uint `gorm:"primaryKey" goal:"<number>primary,sortable,desc,uint"`
    File string `gorm:"unique" binding:"required" goal:"<file>unique,uploadTo=uploads,globalSearch,postonly"`
    Creator uint `gorm:"index" goal:"<number>autowired,uint"`
}
