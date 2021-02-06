package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/TechMinerApps/upmaster/modules/database"
	"github.com/TechMinerApps/upmaster/modules/utils"
	"github.com/TechMinerApps/upmaster/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// UpMaster is a object of entire app
type UpMaster struct {
	Config Config

	DB         *gorm.DB
	HTTPServer *gin.Engine

	viper  *viper.Viper
	logger *logrus.Logger
	wg     *sync.WaitGroup
}

// Config has all the configuration required by UpMaster
type Config struct {
	Port            int
	RDBMSConfig     database.RDBMSConfig
	InfluxDBConfig  database.InfluxDBConfig
	OAuthGCInterval int
}

// NewUpMaster is used to generate a UpMaster object
// no error is returned, so error must be handled within NewUpMaster
func NewUpMaster() *UpMaster {
	var app UpMaster

	app.setupLogger()
	app.setupViper()
	app.setupDB()
	app.setupRouter()

	app.wg = &sync.WaitGroup{}

	return &app
}

func (u *UpMaster) Start() {
	u.wg.Add(1)
	go u.HTTPServer.Run(":" + strconv.Itoa(u.Config.Port))
	return
}

func (u *UpMaster) Stop() {

}

func (u *UpMaster) setupRouter() {

	// Parse Config
	var routerCfg router.Config
	routerCfg.DB = u.DB
	routerCfg.DBName = u.Config.RDBMSConfig.MySQLConfig.DBName
	routerCfg.OAuthGCInterval = u.Config.OAuthGCInterval

	// Create new server
	u.HTTPServer = router.NewRouter(routerCfg)
}

func (u *UpMaster) setupDB() {
	return
}

func (u *UpMaster) setupViper() {
	u.viper = viper.New()
	pflag.String("config", "config", "config file name")
	pflag.Parse()
	u.viper.BindPFlags(pflag.CommandLine)

	if u.viper.IsSet("config") {
		u.viper.SetConfigFile(u.viper.GetString("config"))
	} else {
		u.viper.SetConfigName("config")
		u.viper.SetConfigType("yaml")
		u.viper.AddConfigPath(utils.AbsPath(""))
		u.viper.AddConfigPath("/etc/upmaster")
	}

	u.viper.SetEnvPrefix("UPMASTER")
	u.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	u.viper.AutomaticEnv()

	err := u.viper.Unmarshal(u.Config)
	if err != nil {
		// Used logger here, so setupLogger before setupViper
		u.logger.Fatalf("unable to decode into struct, %v", err)
	}
}

func (u *UpMaster) setupLogger() {
	u.logger = logrus.New()
}
