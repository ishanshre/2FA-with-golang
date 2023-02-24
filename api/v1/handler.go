package v1

import (
	"fmt"
	"net/http"
)

func (s *ApiServer) handleSignUpUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return nil
	}
	return fmt.Errorf("%s method not allowed", r.Method)
}
