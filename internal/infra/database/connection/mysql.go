package connection

import (
	"database/sql"
	"fmt"

	"github.com/rafaelsouzaribeiro/uow-golang/configs"
)

func NewConnection(config *configs.Conf) *sql.DB {
	dbConn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err)
	}
	return dbConn
}
