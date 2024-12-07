package app

import (
	"github.com/webbsalad/storya-content-backend/internal/api/content"
	"github.com/webbsalad/storya-content-backend/internal/config"
	pb "github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"github.com/webbsalad/storya-content-backend/internal/repository/content/pg"

	v1 "github.com/webbsalad/storya-content-backend/internal/service/content/v1"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			initDB,
		),
		grpcOptions(),
		serviceOptions(),
	)
}

func serviceOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			content.NewImplementation,
			pg.NewRepository,
			v1.NewService,
		),
		fx.Invoke(
			pb.RegisterContentServiceServer,
		),
	)
}
