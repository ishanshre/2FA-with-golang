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

func (s *ApiServer) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	if err := checkPostMethod(r); err != nil {
		return err
	}
	req := new(models.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	password, err := s.store.UserLogin(req.Username)
	if err != nil {
		return err
	}
	if err := utils.VerifyPassword(password.Password, req.Password); err != nil {
		return err
	}
	res, err := s.store.UserUpdateLogin(req.Username)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, res)
}

func (s *ApiServer) handleGenerateOTP(w http.ResponseWriter, r *http.Request) error {
	if err := checkPostMethod(r); err != nil {
		return err
	}
	req := new(models.Username)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	storeOtp, err := utils.GenerateOtp(req.UserName)
	if err != nil {
		return err
	}
	updateVal := &models.Update{
		Username:     req.UserName,
		Otp_secret:   storeOtp.Otp_secret,
		Otp_auth_url: storeOtp.Otp_auth_url,
	}
	if err := s.store.UserUpdateOtp(updateVal); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, ApiSuccess{Success: "otp sercret and url updated"})
}

func (s *ApiServer) handleVerifyOTP(w http.ResponseWriter, r *http.Request) error {
	if err := checkPostMethod(r); err != nil {
		return err
	}
	req := new(models.Token)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	secret, err := s.store.GetSercret(req.Username)
	if err != nil {
		return err
	}
	if err := utils.ValidateOTP(&models.OTP{Token: req.Token, Otp_secret: secret.Secret}); err != nil {
		return err
	}
	if err := s.store.UserUpdateEnableOtp(req.Username); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, ApiSuccess{Success: "otp verified and enabled"})
}

func (s *ApiServer) handleValidateOtp(w http.ResponseWriter, r *http.Request) error {
	if err := checkPostMethod(r); err != nil {
		return err
	}
	req := new(models.Token)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil
	}
	secret, err := s.store.GetSercret(req.Username)
	if err != nil {
		return err
	}
	if err := utils.ValidateOTP(&models.OTP{Token: req.Token, Otp_secret: secret.Secret}); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, ApiSuccess{Success: "otp is valid"})
}
