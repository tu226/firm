package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type (
	DBPool struct {
		db *gorm.DB
	}
)

var (
	dbpool DBPool
)

func init() {
	var dnsraw = fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8&parseTime=True&loc=Local",
		beego.AppConfig.String("MysqlUsername"), beego.AppConfig.String("MysqlPassword"), beego.AppConfig.String("MysqlAddr"))
	var dns = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		beego.AppConfig.String("MysqlUsername"), beego.AppConfig.String("MysqlPassword"), beego.AppConfig.String("MysqlAddr"), beego.AppConfig.String("DbName"))
	log.Printf("dnsraw:%v, dns:%v\n", dnsraw, dns)
	dbpool.newDB(dnsraw, dns, beego.AppConfig.String("DbName"))
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&FirmProduct{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&DemoappTbl{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&AppDev{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Firmware{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&ChipInfo{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&QueryInfo{})
	dbpool.db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&DownloadInfo{})
}

func (this *DBPool) newDB(dnsraw, dns, dbname string) {
	db, err := sql.Open("mysql", dnsraw)
	beego.Debug("newDB:", db)
	if err != nil {
		log.Printf("sql.Open:%v\n", err)
		os.Exit(-1)
	}
	defer db.Close()

	var sqlStr = fmt.Sprintf("create database if not exists %s default character set utf8", dbname)
	if _, err = db.Exec(sqlStr); err != nil {
		log.Printf("db.Exec:%v\n", err)
		os.Exit(-1)
	}

	this.db, err = gorm.Open("mysql", dns)
	if err != nil {
		log.Printf("gorm.Open:%v\n", err)
		os.Exit(-1)
	}

	this.db.DB().SetMaxOpenConns(500)
	this.db.DB().SetMaxIdleConns(100)
}
