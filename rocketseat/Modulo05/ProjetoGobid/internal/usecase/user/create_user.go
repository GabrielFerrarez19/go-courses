package user

import (
	"context"

	"ProjetoGobid/internal/validador"
)

type CreateUserReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validador.Evaluator {
	var eval validador.Evaluator

	eval.CheckField(validador.NotBlank(req.UserName), "user_name", "this field cannot be empty")
	eval.CheckField(validador.NotBlank(req.Email), "email", "this field cannot be empty")
	eval.CheckField(validador.Matches(req.Email, validador.EmailRX), "email", "must be a valid email")
	eval.CheckField(validador.NotBlank(req.Bio), "bio", "this field cannot be empty")
	eval.CheckField(validador.NotBlank(req.Password), "password", "this field cannot be empty")
	eval.CheckField(validador.MinChars(req.Bio, 10) && validador.MaxChars(req.Bio, 255), "bio", "this field must have length between 10 and 255")

	eval.CheckField(validador.MinChars(req.Password, 8), "password", "password must bigger than 8 chars")

	eval.CheckField(validador.Matches(req.Email, validador.EmailRX), "email", "must be a valid email")
	// validate stuff
	return eval
}
