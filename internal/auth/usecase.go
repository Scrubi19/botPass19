package auth

import "context"

type UseCase interface {
	AuthUseCase()
	SignUp(ctx context.Context, username, chatId, password string) error
	SignIn(ctx context.Context, username, chatId, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, error)
}
