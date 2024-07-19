package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysql的链接
func MysqlClient(handler func(db *gorm.DB) error) error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		//user, pass, hort, port, dbname,
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.hort"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db, err := cli.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("mysqlOk")
	return handler(cli)
}

// 开启事务
func BeginClient(handle func(db *gorm.DB) error) error {
	return MysqlClient(func(db *gorm.DB) error {
		var err error
		tx := db.Begin()
		defer func() {
			if err == nil {
				fmt.Println("提交")
				tx.Commit()
			} else {
				fmt.Println("回滚")
				tx.Rollback()
			}
		}()

		return handle(tx)
	})

}
