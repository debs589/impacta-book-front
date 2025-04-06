package routes

import (
	"impacta-book/src/handlers"
	"net/http"
)

var mainPageRoute = Route{
	URI:                    "/home",
	Method:                 http.MethodGet,
	Function:               handlers.LoadMainPage,
	AuthenticationRequired: true,
}
