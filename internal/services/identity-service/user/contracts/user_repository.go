package contracts

import (
	"context"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/identity-service/user/models"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, product *models.User) (*models.User, error)
}
