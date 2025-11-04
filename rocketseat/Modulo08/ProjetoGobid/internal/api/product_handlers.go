package api

import (
	"context"
	"net/http"

	"ProjetoGobid/internal/jsonutils"
	"ProjetoGobid/internal/services"
	"ProjetoGobid/internal/usecase/product"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api *Api) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[product.CreateProductReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}
	userID, ok := api.Sessions.Get(r.Context(), "AuthenticateUserId").(uuid.UUID)
	if !ok {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "unexpected error, try again laster",
		})
		return
	}

	productId, err := api.ProductServices.CreateProduct(r.Context(), userID, data.ProductName, data.Description, data.Baseprice, data.AuctionEnd)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "failed to created auction try again laster",
		})
		return
	}

	ctx, _ := context.WithDeadline(context.Background(), data.AuctionEnd)

	auctionRoom := services.NewAuctionRoom(ctx, productId, api.BidsServices)

	go auctionRoom.Run()

	api.AuctionLobby.Lock()
	api.AuctionLobby.Rooms[productId] = auctionRoom
	api.AuctionLobby.Unlock()

	jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"message":    "Auction has started with success",
		"product_id": productId,
	})
}

func (api *Api) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	data, err := api.ProductServices.GetProduct(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "failed to get product try again laster",
		})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message":  "get to product is success",
		"products": data,
	})
}

func (api *Api) HandleGetProductById(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "id")
	productIdUUID, err := uuid.Parse(product_id)
	data, err := api.ProductServices.GetProductById(r.Context(), productIdUUID)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "failed to get product try again laster",
		})
		return
	}
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message":  "get to product is success",
		"products": data,
	})
}
