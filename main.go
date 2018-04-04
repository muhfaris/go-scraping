package main

import (
	"net/http"

	"github.com/muhfaris/go-transaction/infrastructure"
	"github.com/muhfaris/go-transaction/interfaces"
	"github.com/muhfaris/go-transaction/usecases"
)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/production.sqlite")

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler
	handlers["DbItemRepo"] = dbHandler
	handlers["DbOrderRepo"] = dbHandler

	orderInteractor := new(usecases.OrderInteractor)
	orderInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
	orderInteractor.ItemRepository = interfaces.NewDbItemRepo(handlers)
	orderInteractor.OrderRepository = interfaces.NewDbOrderRepo(handlers)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.OrderInteractor = orderInteractor

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowOrder(res, req)
	})
	http.ListenAndServe(":1818", nil)
}
