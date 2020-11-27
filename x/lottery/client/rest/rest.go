package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers lottery-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/lottery/prizeAnnounce", createPrizeAnnounceHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/lottery/prizeAnnounce", listPrizeAnnounceHandler(cliCtx, "lottery")).Methods("GET")
		r.HandleFunc("/lottery/prizeAnnounce/{key}", getPrizeAnnounceHandler(cliCtx, "lottery")).Methods("GET")
		r.HandleFunc("/lottery/prizeAnnounce", setPrizeAnnounceHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/lottery/prizeAnnounce", deletePrizeAnnounceHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/lottery/ticket", createTicketHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/lottery/ticket", listTicketHandler(cliCtx, "lottery")).Methods("GET")
		r.HandleFunc("/lottery/ticket/{key}", getTicketHandler(cliCtx, "lottery")).Methods("GET")
		r.HandleFunc("/lottery/ticket", setTicketHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/lottery/ticket", deleteTicketHandler(cliCtx)).Methods("DELETE")
		r.HandleFunc("/lottery/ticket/lottery/{key}", listTicketByIdHandler(cliCtx, "lottery")).Methods("GET")
		
		r.HandleFunc("/lottery/lottery", createLotteryHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/lottery/lottery", listLotteryHandler(cliCtx, "lottery")).Methods("GET")
		r.HandleFunc("/lottery/lottery/{key}", getLotteryHandler(cliCtx, "lottery")).Methods("GET")
		r.HandleFunc("/lottery/lottery", setLotteryHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/lottery/lottery", deleteLotteryHandler(cliCtx)).Methods("DELETE")

		
}
