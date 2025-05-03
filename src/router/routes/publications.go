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
}
