package env

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	color "github.com/logrusorgru/aurora"
)

type Conf struct {
	DB            DataBase
	WebServerPort string
	GinMode       string
	LogLevel      string
}

type DataBase struct {
	Host     string
	Port     int64
	Driver   string
	User     string
	Name     string
	Password string
	TimeZone string
}

var Config Conf

func Load() {

	var err error
	var now = time.Now().Format("02/01/2006 15:04:05")

	if err := godotenv.Load(); err != nil {
		fmt.Println(color.Yellow(fmt.Sprintf("%s -> Running application without .env file", now)))
	}

	Config.GinMode = os.Getenv("GIN_MODE")
	Config.LogLevel = os.Getenv("LOG_LEVEL")
	Config.WebServerPort = os.Getenv("WEB_SERVER_PORT")

	Config.DB = DataBase{}
	Config.DB.Driver = os.Getenv("DB_DRIVER")
	Config.DB.Host = os.Getenv("DB_HOST")
	Config.DB.User = os.Getenv("DB_USER")
	Config.DB.Name = os.Getenv("DB_NAME")
	Config.DB.Password = os.Getenv("DB_PASSWORD")
	Config.DB.TimeZone = os.Getenv("DB_TIMEZONE")

	Config.DB.Port, err = strconv.ParseInt(os.Getenv("DB_PORT"), 0, 64)
	if err != nil {
		fmt.Println(color.Yellow(fmt.Sprintf("%s -> Running application with the default value for the variable DB_SKYLER_PORT, because: %s", now, err.Error())))
	}

}
