package main

import (
	"boilerplate/internal/connection"
	"boilerplate/internal/model"
	"boilerplate/internal/route"
	"strings"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	_ "go.uber.org/automaxprocs"
)

func init() {
	viper.SetConfigName("env")
	viper.AddConfigPath("./")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func main() {
	svc := viper.GetString("APP.NAME")
	log := connection.NewLog(svc)

	driver, dsn := viper.GetString("APP.DB.DRIVER"), viper.GetString("APP.DB.DSN")
	db, err := connection.NewDB(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	conf := model.NewConf(db, log)
	app := route.New(conf)

	port := viper.GetString("APP.PORT")
	log.Fatal(app.Listen(port))
}
