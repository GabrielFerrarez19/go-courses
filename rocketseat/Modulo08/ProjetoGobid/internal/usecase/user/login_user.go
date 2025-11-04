package user

import (
	"context"

	"ProjetoGobid/internal/validador"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserReq) Valid(ctx context.Context) validador.Evaluator {
	var eval validador.Evaluator

	eval.CheckField(validador.Matches(req.Email, validador.EmailRX), "email", "must be a valid email")
	eval.CheckField(validador.NotBlank(req.Password), "password", "this field cannot be black")

	return eval
}
