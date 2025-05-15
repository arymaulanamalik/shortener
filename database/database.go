package database

import (
	"database/sql"

	"github.com/arymaulanamalik/shortener/config"
	"github.com/arymaulanamalik/shortener/database/mysql"
)

type DatabaseConnect struct {
	Mysql *sql.DB
}

func NewDatabase(cfg config.Config) (res DatabaseConnect, err error) {
	if cfg.Database.MysqlEnabled {
		res.Mysql, err = mysql.Connect()
		if err != nil {
			return res, err
		}
		defer res.Mysql.Close()
	}
	return res, nil
}
