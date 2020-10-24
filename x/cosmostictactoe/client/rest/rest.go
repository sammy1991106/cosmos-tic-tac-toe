package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers cosmostictactoe-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/cosmostictactoe/game", createGameHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/cosmostictactoe/game", listGameHandler(cliCtx, "cosmostictactoe")).Methods("GET")
	r.HandleFunc("/cosmostictactoe/game/{key}", getGameHandler(cliCtx, "cosmostictactoe")).Methods("GET")
	r.HandleFunc("/cosmostictactoe/game", joinGameHandler(cliCtx)).Methods("PUT")

}
