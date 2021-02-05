package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/TechMinerApps/upmaster/modules/utils"
)

// DBType is just a renamed int
type DBType int

// DBType constants
const (
	SQLITE DBType = 0
	MYSQL  DBType = 1
)

// Config is the config used to start a DB connection
type Config struct {
	Type         DBType
	SQLiteConfig sqliteConfig
	MySQLConfig  mysqlConfig
}
type sqliteConfig struct {
	Path string
}

type mysqlConfig struct {
	Username string
	Password string
	Host     string
	DBName   string
}

// NewDBConnection returns a DB object based on config provided
func NewDBConnection(c Config) (*gorm.DB, error) {

	var err error
	var DB *gorm.DB

	switch c.Type {
	case SQLITE:
		DB, err = gorm.Open(sqlite.Open(utils.AbsPath(c.SQLiteConfig.Path)), &gorm.Config{})
	case MYSQL:
		dsn := c.MySQLConfig.Username +
			":" +
			c.MySQLConfig.Password +
			"@tcp(" +
			c.MySQLConfig.Host +
			")/" +
			c.MySQLConfig.DBName +
			"?charset=utf8mb4"
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	// Handle errors
	if err != nil {
		return nil, err
	}
	if DB.Error != nil {
		return nil, DB.Error
	}

	return DB, nil
}
