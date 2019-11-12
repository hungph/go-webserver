package apps

import (
	res "../utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resWriter http.ResponseWriter, req *http.Request) {
		resWriter.WriteHeader(http.StatusNotFound)
		res.Respond(resWriter, res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.NotFound.Code(),
			0, res.ErrorConstants.NotFound.Message()))
		next.ServeHTTP(resWriter, req)
	})
}
