package bolt

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"opsie/pkg/errors"

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
	
	router.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		final(w, req)
	}).Methods(string(method))

	return router
}




// ParseBody reads the JSON body of an HTTP request into the given payload.
// It returns an error if the body is missing, contains invalid JSON,
// or includes unexpected fields (due to DisallowUnknownFields).
func ParseBody(w http.ResponseWriter, r *http.Request, payload any) {
	if r.Body == nil {
		WriteErrorResponse(w, http.StatusBadRequest, "missing request body")
		return 
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(payload)
}


// WriteResponse writes a JSON response with the given status code and payload.
// It automatically sets the "Content-Type" header to "application/json".
// Returns an error if JSON encoding fails (e.g., unsupported types).
func WriteResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		fmt.Println("failed to encode response:", err)
		return err
	}
	return nil
}


// WriteErrorResponse writes a JSON error response with the given HTTP status code.
// The response body has a consistent structure:
func WriteErrorResponse(w http.ResponseWriter, status int, message string, err ...error) {
	var underlying error
	if len(err) > 0 {
		underlying = err[0]
	}

	resp := map[string]any{
		"code":    status,
		"message": message,
	}

	if underlying != nil {
		resp["error"] = underlying.Error()
	}

	WriteResponse(w, status, resp)
}



// ErrorHandler processes an error returned from the service layer.
// If the error is a *CsError, it responds with its specific HTTP code and message.
// Otherwise, it logs the internal error and responds with a generic 500 error.
//
// Returns true if a response was written (so the handler should return immediately).
func ErrorHandler(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	if cerr, ok := err.(*errors.Error); ok {
		WriteErrorResponse(w, cerr.Code, cerr.Message, cerr.Err)
		return true
	}

	// unexpected error
	WriteErrorResponse(w, http.StatusInternalServerError, "internal server error")
	log.Panic("internal error:", err)
	return true
}
