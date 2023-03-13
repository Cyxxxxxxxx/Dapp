package mysql

import (
	"Dapp/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init MySQL connet
func Init(cfg *setting.MysqlConf) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.Max_conns)      //设置最大连接数量
	db.SetMaxIdleConns(cfg.Max_idle_conns) //设置最大空闲连接数
	return
}

func GetDb() *sqlx.DB {
	return db
}

// Close mysql connet
func Close() {
	_ = db.Close()
}
