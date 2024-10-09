package mysql

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"note/configs"
	"strconv"
)

var DBGorm *gorm.DB

var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)

func Setup() {
	Host = configs.CfgMysql.GetString("host")
	Port = configs.CfgMysql.GetInt("port")
	Name = configs.CfgMysql.GetString("database")
	Username = configs.CfgMysql.GetString("user")
	Password = configs.CfgMysql.GetString("passwd")

	var err error

	conn := GetMysqlConnect()

	log.Println(conn)

	var db Database
	db = new(Mysql)
	DBGorm, err = db.Open("mysql", conn)
	if err != nil {
		log.Fatalf("%s connect error %v", DbType, err)
	} else {
		log.Printf("%s connect success!", DbType)
	}

	if DBGorm.Error != nil {
		log.Fatalf("database error %v", DBGorm.Error)
	}

	DBGorm.LogMode(true)

	AutoMigrate()

}

func AutoMigrate() {
	//err := DBGorm.AutoMigrate(&model.Note{}, &model.Tag{})
	//if err != nil {
	//	log.Fatalf("auto migration error %v", err)
	//} else {
	//	log.Println("Auto migration completed successfully")
	//}
}

func GetMysqlConnect() string {
	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")
	return conn.String()
}

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}
