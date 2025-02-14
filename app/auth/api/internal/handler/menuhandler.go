package handler

import (
	"net/http"
	"strconv"

	"zero-app/app/auth/api/internal/svc"
	"zero-app/app/auth/api/internal/types"
	"zero-app/app/auth/model"
	"zero-app/app/auth/model/basemodel"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MenuList(c *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")
		sort := r.URL.Query().Get("sort")
		key := r.URL.Query().Get("key")
		menuTypeStr := r.URL.Query().Get("type")
		parentIdStr := r.URL.Query().Get("parent_id")

		var whereOrder []basemodel.PageWhereOrder

		order := "ID DESC"
		if len(sort) >= 2 {
			orderType := sort[0:1]
			order = sort[1:len(sort)]
			if orderType == "+" {
				order += " ASC"
			} else {
				order += " DESC"
			}
		}
		whereOrder = append(whereOrder, basemodel.PageWhereOrder{Order: order})
		if key != "" {
			v := "%" + key + "%"
			var arr []interface{}
			arr = append(arr, v)
			arr = append(arr, v)
			whereOrder = append(whereOrder, basemodel.PageWhereOrder{Where: "name like ? or code like ?", Value: arr})
		}

		menuType, _ := strconv.ParseInt(menuTypeStr, 10, 64)
		if menuType > 0 {
			var arr []interface{}
			arr = append(arr, menuType)
			whereOrder = append(whereOrder, basemodel.PageWhereOrder{Where: "menu_type = ?", Value: arr})
		}

		parent_id, _ := strconv.ParseInt(parentIdStr, 10, 64)
		if parent_id > 0 {
			var arr []interface{}
			arr = append(arr, parent_id)
			whereOrder = append(whereOrder, basemodel.PageWhereOrder{Where: "parent_id = ?", Value: arr})
		}
		var total uint64
		list := []model.Menu{}
		page, _ := strconv.ParseUint(pageStr, 10, 64)
		limit, _ := strconv.ParseUint(limitStr, 10, 64)
		err := basemodel.GetPage(c.DbEngin, &model.Menu{}, &model.Menu{}, &list, page, limit, &total, whereOrder...)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		res := types.ResponsePage{
			Code:    types.SUCCESS_CODE,
			Message: "ok",
			Data: types.ResponsePageData{
				Total: total,
				Items: list,
			},
		}

		httpx.OkJson(w, res)
	}
}

func MenuDetail(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")

		type Res struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		res := Res{
			Id:   id,
			Name: name,
		}

		httpx.OkJson(w, res)
	}
}

func AllMenu(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func MenuButtonList(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func MenuUpdate(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func MenuCreate(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func MenuDelete(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
