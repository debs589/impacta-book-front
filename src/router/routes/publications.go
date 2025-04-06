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
}
