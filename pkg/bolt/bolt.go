package bolt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPMethod string

const (
	MethodGet     HTTPMethod = "GET"
	MethodPost    HTTPMethod = "POST"
	MethodPut     HTTPMethod = "PUT"
	MethodPatch   HTTPMethod = "PATCH"
	MethodDelete  HTTPMethod = "DELETE"
	MethodOptions HTTPMethod = "OPTIONS"
	MethodHead    HTTPMethod = "HEAD"
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
	var errMsg string
	if len(err) > 0 && err[0] != nil {
		errMsg = err[0].Error()
	}

	resp := map[string]any{
		"code":    status,
		"message": message,
	}
	if errMsg != "" {
		resp["error"] = errMsg
	}

	WriteResponse(w, status, resp)
}



// ErrorHandler processes an error returned from the service layer.
// If the error is a *CsError, it responds with its specific HTTP code and message.
// Otherwise, it logs the internal error and responds with a generic 500 error.
//
// Returns true if a response was written (so the handler should return immediately).
//
// func ErrorHandler(w http.ResponseWriter, err error) bool {
// 	// Skip if no error or typed nil error
// 	if err == nil || (reflect.ValueOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil()) {
// 		return false
// 	}

// 	// Handle our custom errors
// 	if cerr, ok := err.(*errors.Error); ok {
// 		msg := cerr.Message
// 		if msg == "" {
// 			msg = http.StatusText(cerr.Code)
// 		}

// 		// Include underlying error only if available
// 		if cerr.Err != nil {
// 			WriteErrorResponse(w, cerr.Code, msg, cerr.Err)
// 		} else {
// 			WriteErrorResponse(w, cerr.Code, msg)
// 		}
// 		return true
// 	}

// 	// Handle unexpected errors
// 	WriteErrorResponse(w, http.StatusInternalServerError, "internal server error", err)
// 	log.Printf("⚠️ Internal error: %+v\n", err)

// 	return true
// }

