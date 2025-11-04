package api

import (
	"ProjetoGobid/internal/services"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type Api struct {
	Router          *chi.Mux
	UserServices    services.UserServices
	ProductServices services.ProductService
	BidsServices    services.BidsService
	Sessions        *scs.SessionManager
	WeUpgrader      websocket.Upgrader
	AuctionLobby    services.AuctionLobby
}
