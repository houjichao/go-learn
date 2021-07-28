package myOrm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
  UserInfo 用户信息
*/
type UserInfo struct {
	ID      uint
	NAME    string
	AGE     uint
	EMAIL   string
	DELETED uint
}

func OperationDb() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/learn?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userTable := db.Table("user_info")
	userTable.CreateTable(&UserInfo{})
	// 自动迁移
	//这里会自动创建表，总结,使用gorm自动对应实体创建表结构时,字段一定要大写
	//db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "JACK", 13, "123@qq.com", 0}
	u2 := UserInfo{2, "MIKE", 18, "456@qq.com", 0}
	// 创建记录
	userTable.Create(&u1)
	userTable.Create(&u2)
	// 查询
	/*var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)
	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)
	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)*/
}
