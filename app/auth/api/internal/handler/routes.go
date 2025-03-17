package handler

import (
	"net/http"
	"zero-app/app/auth/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			// Menu routes
			{
				Method:  http.MethodGet,
				Path:    "/menu/list",
				Handler: MenuList(serverCtx),
			},
		},
		//使用jwt
		// rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
