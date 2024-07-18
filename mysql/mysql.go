package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func BeginClient() {
	//dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
	//	user, pass, hort, port, dbname,
	//)
	//db.Begin()
}

//func ()  {
//
//}
