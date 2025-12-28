package track

import (
	"context"

	"github.com/monitoria-saas/tracking-api/internal/database"
)

type IService interface {
	SendMonitoring(ctx context.Context)
}

type service struct {
	repo *database.TrackRepository
}
