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
			{
				Method:  http.MethodGet,
				Path:    "/menu/detail",
				Handler: MenuDetail(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu/allmenu",
				Handler: AllMenu(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/menu/menubuttonlist",
				Handler: MenuButtonList(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/delete",
				Handler: MenuDelete(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/update",
				Handler: MenuUpdate(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/menu/create",
				Handler: MenuCreate(serverCtx),
			},

			// User routes
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: UserInfo(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserLogin(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/logout",
				Handler: UserLogout(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/editpwd",
				Handler: UserEditPwd(serverCtx),
			},

			// // Admins routes
			// {
			// 	Method:  http.MethodGet,
			// 	Path:    "/admins/list",
			// 	Handler: AdminsList(serverCtx),
			// },
			// {
			// 	Method:  http.MethodGet,
			// 	Path:    "/admins/detail",
			// 	Handler: AdminsDetail(serverCtx),
			// },
			// {
			// 	Method:  http.MethodGet,
			// 	Path:    "/admins/adminsroleidlist",
			// 	Handler: AdminsRoleIDList(serverCtx),
			// },
			// {
			// 	Method:  http.MethodPost,
			// 	Path:    "/admins/delete",
			// 	Handler: AdminsDelete(serverCtx),
			// },
			// {
			// 	Method:  http.MethodPost,
			// 	Path:    "/admins/update",
			// 	Handler: AdminsUpdate(serverCtx),
			// },
			// {
			// 	Method:  http.MethodPost,
			// 	Path:    "/admins/create",
			// 	Handler: AdminsCreate(serverCtx),
			// },
			// {
			// 	Method:  http.MethodPost,
			// 	Path:    "/admins/setrole",
			// 	Handler: AdminsSetRole(serverCtx),
			// },

			// // Role routes
			// {
			// 	Method:  http.MethodGet,
			// 	Path:    "/role/list",
			// 	Handler: RoleList(serverCtx),
			// },
			// {
			// 	Method:  http.MethodGet,
			// 	Path:    "/role/detail",
			// 	Handler: RoleDetail(serverCtx),
			// },
			// {
			// 	Method:  http.MethodGet,
			// 	Path:    "/role/rolemenuidlist",
			// 	Handler: RoleMenuIDList(serverCtx),
			// },
			// {
			// 	Method:  http.MethodGet,
			// 	Path:    "/role/allrole",
			// 	Handler: AllRole(serverCtx),
			// },
			// {
			// 	Method:  http.MethodPost,
			// 	Path:    "/role/delete",
			// 	Handler: RoleDelete(serverCtx),
			// },
			// {
			// 	Method:  http.MethodPost,
			// 	Path:    "/role/update",
			// 	Handler: RoleUpdate(serverCtx),
			// },
			// {
			// 	Method:  http.MethodPost,
			// 	Path:    "/role/create",
			// 	Handler: RoleCreate(serverCtx),
			// },
			// {
			// 	Method:  http.MethodPost,
			// 	Path:    "/role/setrole",
			// 	Handler: RoleSetRole(serverCtx),
			// },
		},
	)
}
