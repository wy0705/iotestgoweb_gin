package web

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //必须手动导入
	"time"
)
// UserInfo 用户信息
type UserInfo struct {
	ID string `gorm:"primary_key"`
	Name string
	Sex string
	Context string
}
func XIII13() {
	//数据库名添加默认前缀
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "my_" + defaultTableName;
	}

	//CREATE DATABASE db1;
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	db.SingularTable(true);//设置表名和结构体名相同，不加s

	if err != nil{
		panic(err.Error())
	}
	defer db.Close()

	// 自动迁移
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8;").AutoMigrate(&UserInfo{})

	u1 := UserInfo{"1", "小明", "男", "学生"}
	u2 := UserInfo{"2", "李老师", "女", "老师"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "context=?", "老师")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&uu).Update("context", "职员")
	// 删除
	db.Delete(&u)

	//模型
	type User struct {
		gorm.Model
		Name         string
		Age          sql.NullInt64
		Birthday     *time.Time
		Email        string  `gorm:"type:varchar(100);unique_index"`
		Role         string  `gorm:"size:255"` // 设置字段大小为255
		MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
		Num          int     `gorm:"column:mynum;AUTO_INCREMENT"` // 设置 num 为自增类型
		Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
		IgnoreMe     int     `gorm:"-"` // 忽略本字段
	}
	user:=User{}
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8;").AutoMigrate(&user)

	//默认前缀对指定名字的创建无效
	db.Table("xxx").CreateTable(&User{});//手动创建表
	var deleted_users []User
	db.Table("xxx").Find(&deleted_users)// SELECT * FROM xxx;
	fmt.Printf("%#v\n", deleted_users)
	db.Table("xxx").Where("name = ?", "xxx").Delete(&deleted_users)// DELETE FROM xxx WHERE name = 'xxx';

}
