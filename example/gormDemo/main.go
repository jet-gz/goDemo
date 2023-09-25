package main

import (
	"gormDemo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "root:123456@tcp(121.4.181.166:3306)/gormTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 这里可以设置 别的配置
		//DryRun: false,//生成 SQL 但不执行，可以用于准备或测试生成的 SQL
		Logger: logger.Default.LogMode(logger.Info),
	})

	// err := db.AutoMigrate(&model.User{}) // 创建表  db.AutoMigrate(&User{}, &Product{}, &Order{}) 多个
	// err := db.AutoMigrate(&model.Student{})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	//user := model.User{Age: 35, Author: model.Author{Name: "Jet1", Email: "Jet123@outlook.com"}}
	//user := model.User{Age: 35, Name: "Jet", Email: "Jet123@outlook.com"}

	//批量插入
	// user := []*model.User{
	// 	{Age: 35, Author: model.Author{Name: "Jet1", Email: "Jet123@outlook.com"}},
	// 	{Age: 36, Author: model.Author{Name: "Jet334", Email: "Jet4556@outlook.com"}},
	// }
	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	//result := db.Create(&user)
	//fmt.Println("插入Id为：", user.ID)
	//fmt.Println("受影响行数", result.RowsAffected)

	// 给指定字段赋值
	// user := model.User{Age: 35, Author: model.Author{Name: "Jet1", Email: "Jet123@outlook.com"}}
	// db.Select("Name", "Email").Create(&user)

	// // 忽略字段赋值
	// user2 := model.User{Age: 35, Author: model.Author{Name: "Jetok", Email: "Jet123@outlook.com"}}
	// db.Omit("Age", "Email").Create(&user2)

	//测试Before
	// s := model.Student{Name: "Jet", Age: 36}
	// fmt.Println("创建之前...")
	// db.Create(&s)
	// fmt.Println("创建之后...")

	// // 添加map
	// db.Model(&model.Student{}).Create(map[string]interface{}{
	// 	"Name": "jet4",
	// 	"Age":  34,
	// })
	// //map 批量
	// db.Model(&model.Student{}).Create([]map[string]interface{}{
	// 	{"Name": "jet4", "Age": 34},
	// 	{"Name": "jet5", "Age": 56},
	// })

	// db.AutoMigrate(&model.Order{}, &model.Commodity{})
	// o := model.Order{
	// 	OrderNo: string(time.Now().Nanosecond()), Commoditys: []model.Commodity{
	// 		//model.Commodity{ }
	// 		{Name: "铅笔1"}, {Name: "钢笔2"}, {Name: "本子3"},
	// 	},
	// }
	// db.Create(&o)

	// o := model.Order{}
	// db.Where("id=?", 1).First(&o)
	// fmt.Println(o)

	// fmt.Println("===============")
	// db.Model(&o).Association("Commoditys").Find(&o.Commoditys)
	// data, _ := json.Marshal(&o)
	// fmt.Println(string(data))

	//===========多对多=============

	// db.AutoMigrate(&model.Role{}, &model.Permission{}, &model.UserPermission{})

	// r1 := model.Role{
	// 	RoleName: "修理工",
	// 	Permission: []model.Permission{
	// 		{Title: "修改"}, {Title: "删除"},
	// 	},
	// }

	// p1 := model.Permission{
	// 	Title: "上传权限",
	// 	Roles: []model.Role{
	// 		{RoleName: "客户部"}, {RoleName: "设计部"},
	// 	},
	// 	UserPermission: model.UserPermission{
	// 		C1: "c1",
	// 		C2: "c2",
	// 	},
	// }

	// db.Create(&r1)
	// db.Create(&p1)

	//u := model.User{}

	//db.Row("select * from users where id = @id", sql.Named("id", 1)).Find(&u)

	// db.Raw("select * from users where id = @id", sql.Named("id", 9)).Find(&u)
	// data, _ := json.Marshal(&u)
	// fmt.Println(string(data))

	// us := []model.User{}
	// db.Raw("select * from users").Find(&us)
	// data, _ := json.Marshal(&us)
	// fmt.Println(string(data))

	//db.Exec("UPDATE users SET name = ? WHERE id = ?", "王五", 9)

	// s := model.Student{Name: "33333"}
	// r := db.Create(&s)

	// u := model.User{Author: model.Author{Name: "jetrrrrrrrrr"}}
	// r1 := db.Create(&u)
	// fmt.Println("行数r", r.RowsAffected)
	// fmt.Println("行数r1", r1.RowsAffected)

	//r 插入成功，r1 插入失败
	// 自动事务  只要报错就会回滚
	// db.Transaction(func(tx *gorm.DB) error {

	// 	s := model.Student{Name: "33333"}
	// 	r := tx.Create(&s)

	// 	u := model.User{Author: model.Author{Name: "jetrrrrrrrrr"}}
	// 	r1 := tx.Create(&u)
	// 	if r.RowsAffected <= 0 || r1.RowsAffected <= 0 {
	// 		tx.Rollback()
	// 	}
	// 	return nil // 返回nil 就提交事务
	// })

	// s := model.Student{Name: "33333"}
	// u := model.User{Author: model.Author{Name: "jetrrrrrrrrr"}}

	// tx := db.Begin()
	// r := tx.Create(&s)
	// r1 := tx.Create(&u)
	// if r.RowsAffected <= 0 || r1.RowsAffected <= 0 {
	// 	tx.Rollback()
	// } else {
	// 	tx.Commit()
	// }

	//db.Model(&model.User{}).Where("id", "11").Update("Name", "荷塘飞鱼")

	// u := model.User{}
	// db.Where("id", "11").Find(&u)
	// db.Model(&u).Update("Name", "荷塘飞鱼2")

	// u := model.User{}
	// db.Where("id", "11").Find(&u)
	// db.Model(&u).Updates(model.User{Age: 43, Author: model.Author{Name: "Jet"}})

	//根据id删除
	// db.Delete(&model.User{}, 15)
	// 批量删除
	//db.Delete(&model.User{}, []int{13, 14})

	// 软删除
	//db.Delete(&model.Order{}, 5)
	// u := model.Order{}
	// db.Where("id = 5").Find(&u)

	// data, _ := json.Marshal(&u)
	// fmt.Println(string(data))
	// fmt.Println("查询没被软删除的数据")
	// // 查询被软删除的数据
	// db.Unscoped().Where("id=5").Find(&u)
	// data, _ = json.Marshal(&u)
	// fmt.Println(string(data))

	u := model.User{}

	rows, _ := db.Model(&model.User{}).Where("id>?", 10).Select("id", "age").Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &u)
	}

}
