package routes

import (
	"impacta-book/src/handlers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:                    "/publication",
		Method:                 http.MethodPost,
		Function:               handlers.CreatePublication,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/publication/{publicationId}/like",
		Method:                 http.MethodPost,
		Function:               handlers.LikePublication,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/publication/{publicationId}/unlike",
		Method:                 http.MethodPost,
		Function:               handlers.UnlikePublication,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/publication/{publicationId}/edit",
		Method:                 http.MethodGet,
		Function:               handlers.LoadPageEditPublication,
		AuthenticationRequired: true,
	},
	{
		URI:                    "/publication/{publicationId}",
		Method:                 http.MethodPut,
		Function:               handlers.EditPublication,
		AuthenticationRequired: true,
	},
}
