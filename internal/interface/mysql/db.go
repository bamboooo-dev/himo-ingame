package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	dao "github.com/bamboooo-dev/himo-ingame/internal/interface/dao/himo"
	"github.com/go-gorp/gorp"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

const defaultMySQLPort = 3306

// Config is DB configuration.
type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// NewDB is the sql.DB constructor.
func NewDB(cfg Config) (*gorp.DbMap, error) {
	connStr, err := buildConnectionString(cfg)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	dbmap.AddTableWithName(dao.Room{}, "rooms").SetKeys(true, "ID")
	dbmap.AddTableWithName(dao.Theme{}, "themes").SetKeys(true, "ID")
	dbmap.AddTableWithName(dao.RoomTheme{}, "room_themes")

	if err := dbmap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}
	return dbmap, nil
}

func buildConnectionString(cfg Config) (string, error) {
	mysqlCfg := mysql.NewConfig()

	if cfg.Host == "" {
		return "", xerrors.New("db host is not set")
	}

	if cfg.User == "" {
		return "", xerrors.New("db user is not set")
	}

	mysqlCfg.Net = "tcp"
	port := defaultMySQLPort
	if cfg.Port != 0 {
		port = cfg.Port
	}

	// Host に port 番号を含めた場合は port の設定を無視する
	if strings.Contains(cfg.Host, ":") {
		mysqlCfg.Addr = cfg.Host
	} else {
		mysqlCfg.Addr = fmt.Sprintf("%s:%d", cfg.Host, port)
	}

	mysqlCfg.DBName = cfg.Database
	mysqlCfg.User = cfg.User
	mysqlCfg.Passwd = cfg.Password

	ret := mysqlCfg.FormatDSN()
	return ret, nil
}
