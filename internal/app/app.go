package app

import (
	"github.com/webbsalad/storya-content-backend/internal/api/content"
	"github.com/webbsalad/storya-content-backend/internal/api/user_content"
	"github.com/webbsalad/storya-content-backend/internal/config"
	pb "github.com/webbsalad/storya-content-backend/internal/pb/github.com/webbsalad/storya-content-backend/content"
	"github.com/webbsalad/storya-content-backend/internal/repository/content/pg"

	content_service "github.com/webbsalad/storya-content-backend/internal/service/content/v1"
	user_content_service "github.com/webbsalad/storya-content-backend/internal/service/user_content/v1"
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
			content.NewContentImplementation,
			user_content.NewUserContentImplementation,
			pg.NewRepository,
			content_service.NewService,
			user_content_service.NewService,
		),
		fx.Invoke(
			pb.RegisterContentServiceServer,
			pb.RegisterUserContentServiceServer,
		),
	)
}
