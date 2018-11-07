package settings

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

//DBType is the structure which define parameters for database connection
type DBType struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

var (
	ListenAddr  string
	VerboseMode bool
	DB          DBType
)

func init() {
	ListenAddr = "localhost:8181"
	VerboseMode = false
	DB.Host = "127.0.0.1"
	DB.Port = 5432
	DB.Database = "test"
	DB.User = "user"
	DB.Password = "password"
}

func FromFile(file string) error {
	log.Printf("Loading configuration from '%s'", file)

	cfg, err := ini.InsensitiveLoad(file)
	if err != nil {
		return fmt.Errorf("error parsing file '%s', message: %s", file, err)
	}

	newListenAddr := cfg.Section("").Key("ListenAddr").String()

	var newDB DBType
	newDB.Host = cfg.Section("DB").Key("Host").String()
	portInt, err := cfg.Section("DB").Key("Port").Int()
	if err != nil {
		return errors.New(fmt.Sprint("configuration DB.Port has wrong value, message: ", err))
	}
	newDB.Port = uint16(portInt)
	newDB.Database = cfg.Section("DB").Key("Database").String()
	newDB.User = cfg.Section("DB").Key("User").String()
	newDB.Password = cfg.Section("DB").Key("Password").String()

	ListenAddr = newListenAddr
	DB = newDB

	return nil
}
