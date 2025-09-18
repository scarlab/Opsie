package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
func ParseBody(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // reject unknown JSON fields
	return decoder.Decode(payload)
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
//   {
//     "error": {
//       "code": <status>,
//       "error": "<message>"
//     }
//   }
//
// Example:
//   utils.HandleErrorResponse(w, 404, fmt.Errorf("user not found"))
func HandleErrorResponse(w http.ResponseWriter, status int, err error) {
	resp := map[string]any{
		"error": map[string]any{
			"code":  status,
			"error": err.Error(),
		},
	}
	_ = HandleResponse(w, status, resp)
}


// HandleBusinessError processes an error returned from the service layer.
// If the error is a *CsError, it responds with its specific HTTP code and message.
// Otherwise, it logs the internal error and responds with a generic 500 error.
//
// Returns true if a response was written (so the handler should return immediately).
//
// Example:
//   user, err := service.RegisterUser(req)
//   if utils.HandleBusinessError(w, err) {
//       return
//   }
func HandleBusinessError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	if cerr, ok := err.(*CsError); ok {
		HandleErrorResponse(w, cerr.Code, fmt.Errorf("%s", cerr.Message))
		return true
	}

	// unexpected error
	fmt.Println("internal error:", err)
	HandleErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
	return true
}
