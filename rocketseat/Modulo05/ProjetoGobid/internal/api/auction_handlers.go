package api

import (
	"errors"
	"net/http"

	"ProjetoGobid/internal/jsonutils"
	"ProjetoGobid/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api *Api) handleSubscribeUserToAuction(w http.ResponseWriter, r *http.Request) {
	rawProductId := chi.URLParam(r, "product_id")

	productId, err := uuid.Parse(rawProductId)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
			"message": "invalid product id - must be a valid uuid",
		})
		return
	}

	_, err = api.ProductServices.GetProductById(r.Context(), productId)
	if err != nil {
		if errors.Is(err, services.ErrProductNotFound) {
			jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{
				"message": "no product with given id",
			})
			return
		}
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "unexpected error, try again later",
		})
		return
	}

	userid, ok := api.Sessions.Get(r.Context(), "AuthenticateUserId").(uuid.UUID)
	if !ok {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "unexpected error, try again later",
		})
		return
	}

	api.AuctionLobby.Lock()
	room, ok := api.AuctionLobby.Rooms[productId]
	if !ok {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
			"message": "the auction has ended",
		})
	}
	api.AuctionLobby.Unlock()

	conn, err := api.WeUpgrader.Upgrade(w, r, nil)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "could not upgrade connection websocket protocol",
		})
		return
	}

	client := services.NewClient(room, conn, userid)

	room.Register <- client

	go client.ReadEventLoop()
	go client.WriteEventLoop()
	for {
	}
}
