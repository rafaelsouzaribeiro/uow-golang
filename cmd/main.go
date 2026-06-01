package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rafaelsouzaribeiro/uow-golang/configs"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/infra/database/mysql/connection"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/infra/di"
	"github.com/rafaelsouzaribeiro/uow-golang/internal/infra/web/server"
)

func main() {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	con := connection.NewConnection(conf)
	defer con.Close()
	useCase := di.NewDI(con)
	webServer := server.NewServer()
	webServer.SetRoutes(useCase)
	err = webServer.Start(conf.WebPort)

	if err != nil {
		panic(err)
	}
}
