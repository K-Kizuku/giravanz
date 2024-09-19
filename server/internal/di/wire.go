//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/giravanz/server/internal/app/handler"
	"github.com/K-Kizuku/giravanz/server/internal/app/repository"
	"github.com/K-Kizuku/giravanz/server/internal/app/service"
	"github.com/K-Kizuku/giravanz/server/pkg/redis"
	"github.com/google/wire"
)

func InitHandler() *handler.Root {
	wire.Build(
		redis.New,
		repository.NewMessageRepository,
		service.NewMessageService,
		handler.NewMessageHandler,
		handler.New,
	)
	return &handler.Root{}
}
