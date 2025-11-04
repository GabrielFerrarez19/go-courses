package product

import (
	"context"
	"time"

	"ProjetoGobid/internal/validador"

	"github.com/google/uuid"
)

type CreateProductReq struct {
	SellerId    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
}

type GetProductRes struct {
	ID          uuid.UUID `json:"id"`
	SellerID    uuid.UUID `json:"seller_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Baseprice   float64   `json:"baseprice"`
	AuctionEnd  time.Time `json:"auction_end"`
	IsSold      bool      `json:"is_sold"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const minAuctionDuration = 2 * time.Hour

func (req CreateProductReq) Valid(cts context.Context) validador.Evaluator {
	var eval validador.Evaluator

	eval.CheckField(validador.NotBlank(req.ProductName), "product_name", "this field cannot be blank")
	eval.CheckField(validador.NotBlank(req.Description), "description", "this field cannot be blank")
	eval.CheckField(validador.MinChars(req.Description, 10) && validador.MaxChars(req.Description, 255),
		"description", "this field must have a length between 10 and 255")

	eval.CheckField(req.Baseprice > 0, "baseprice", "this filed must be greater than 0")
	eval.CheckField(req.AuctionEnd.Sub(time.Now()) >= minAuctionDuration, "auctionend", "must be at least two hours duration")

	return eval
}
