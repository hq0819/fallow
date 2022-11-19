package configuration

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var Db *gorm.DB

func InitConfiguration() {
	v := viper.New()
	v.AddConfigPath("configuration")
	v.SetConfigName("configuration")
	v.SetConfigType("yml")
	ds := new(Datasource)
	_ = v.ReadInConfig()
	err := v.UnmarshalKey("datasource", ds, func(s *mapstructure.DecoderConfig) { s.TagName = "yaml" })
	if err != nil {
		log.Print("datasouce加载异常")
	}
	dns := ds.Username + `:` + ds.Password +
		`@tcp` + `(` + ds.IP + `:` + ds.Port + `)` + `/` + ds.Database + `?parseTime=True`
	open := mysql.Open(dns)
	db, err := gorm.Open(open, &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: false}})
	if err != nil {
		log.Print("数据库初始化失败")
	}

	Db = db
}
