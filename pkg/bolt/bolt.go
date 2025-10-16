package bolt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	PATCH   HTTPMethod = "PATCH"
	DELETE  HTTPMethod = "DELETE"
	OPTIONS HTTPMethod = "OPTIONS"
	HEAD    HTTPMethod = "HEAD"
)




func  Api(router *mux.Router, method HTTPMethod, path string, handler THandlerFunc, middlewares ...TMiddleware) *mux.Router {
	// Apply user middlewares + bolt logger
	final := Middleware(handler, middlewares...)
	
	if path == "/" {
    	path = ""
	}

	router.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		final(w, req)
	}).Methods(string(method))

	return router
}




// ParseBody reads the JSON body of an HTTP request into the given payload.
// It returns an error if the body is missing, contains invalid JSON,
// or includes unexpected fields (due to DisallowUnknownFields).
//
// Example:
//   var req CreateUserRequest
//   if err := utils.ParseBody(r, &req); err != nil {
//       utils.HandleErrorResponse(w, 400, err)
//       return
//   }
func ParseBody(w http.ResponseWriter, r *http.Request, payload any) {
	if r.Body == nil {
		HandleErrorResponse(w, http.StatusBadRequest, fmt.Errorf("missing request body"))
		return 
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(payload)
}


// HandleResponse writes a JSON response with the given status code and payload.
// It automatically sets the "Content-Type" header to "application/json".
// Returns an error if JSON encoding fails (e.g., unsupported types).
//
// Example: 
//   utils.HandleResponse(w, 200, map[string]string{"status": "ok"})
func HandleResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		fmt.Println("failed to encode response:", err)
		return err
	}
	return nil
}


// HandleErrorResponse writes a JSON error response with the given HTTP status code.
// The response body has a consistent structure:
//
// {
// 	  "code": <status>,
// 	  "error": "<message>"
// }
//
// Example:
//   utils.HandleErrorResponse(w, 404, fmt.Errorf("user not found"))
func HandleErrorResponse(w http.ResponseWriter, status int, err error) {
	resp := map[string]any{
		"code":  status,
		"error": err.Error(),
	}
	HandleResponse(w, status, resp)
}


// HandleServiceError processes an error returned from the service layer.
// If the error is a *CsError, it responds with its specific HTTP code and message.
// Otherwise, it logs the internal error and responds with a generic 500 error.
//
// Returns true if a response was written (so the handler should return immediately).
//
// Example:
//   user, err := service.RegisterUser(req)
//   if utils.HandleServiceError(w, err) {
//       return
//   }
func HandleServiceError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	if cerr, ok := err.(*Error); ok {
		HandleErrorResponse(w, cerr.Code, fmt.Errorf("%s", cerr.Message))
		return true
	}

	// unexpected error
	HandleErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
	log.Panic("internal error:", err)
	return true
}
