package components

import (
    "sync"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "wego3/app/models"
)

var DB *gorm.DB

func init() {
    var once sync.Once

    // 初始化DB客户端
    once.Do(func() {
        var err error

        // 表前缀
        // gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
        //     return "prefix_" + defaultTableName;
        // }

        db, err := gorm.Open("mysql", "root:root@tcp(192.168.1.155:3306)/tosneaker.dev?charset=utf8&parseTime=True&loc=Local")
        if err != nil {
          panic("failed to connect database")
        }

        DB = db

        // 全局禁用表名复数
        // User模型默认的就是user而不是users
        DB.SingularTable(true)

        // 自动迁移，创建表时添加表后缀
        DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User2{})
    })
}
