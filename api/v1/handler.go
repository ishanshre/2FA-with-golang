package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishanshre/2FA-with-golang/api/v1/models"
	"github.com/ishanshre/2FA-with-golang/internals/pkg/utils"
)

func (s *ApiServer) handleSignUpUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		req := new(models.RegisterUser)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		encPass, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user := models.NewRegisterUser(req.Name, req.Username, req.Email, encPass)
		if err := s.store.UserSignUp(user); err != nil {
			return err
		}
		return writeJSON(w, http.StatusCreated, ApiSuccess{Success: "New User Created"})
	}
	return fmt.Errorf("%s method not allowed", r.Method)
}
