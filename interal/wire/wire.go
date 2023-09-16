//go:build wireinject
// +build wireinject

package wire

import (
	"url-shortener/interal/controller"
	"url-shortener/interal/service"

	"github.com/google/wire"
)

func InitializeUrlController() controller.UrlController {
	wire.Build(
		service.NewUrlService,
		controller.NewUrlController,
	)
	return controller.UrlController{}
}
