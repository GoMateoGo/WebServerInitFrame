package core

import (
	"fmt"
	"official_site/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func InitDB(c *config.Mysql) *xorm.Engine {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
	db, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("database connect failed", err)
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Test link failed", err)
		panic(err.Error())
	}
	db.SetTableMapper(core.SameMapper{})   //core.SameMapper{}
	db.SetColumnMapper(core.SnakeMapper{}) // SnakeMapper  SameMapper
	db.ShowSQL(c.Show)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)
	//tn, _ := dbHand.DBMetas()
	fmt.Println("====== Mode [Mysql] Loading success ======")
	return db
}
