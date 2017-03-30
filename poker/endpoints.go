package poker

import (
	"github.com/jasonsoft/napnap"
)

func NewPokerRouter() *napnap.Router {
	router := napnap.NewRouter()

	router.Get("/v1/tables", listTablesEndpoint)
	router.Post("/v1/tables/:id/join", joinGameEndpoint)
	router.Post("/v1/tables/:id/leave", leaveGameEndpoint)

	return router
}

func listTablesEndpoint(c *napnap.Context) {

}

func joinGameEndpoint(c *napnap.Context) {

}

func leaveGameEndpoint(c *napnap.Context) {

}
