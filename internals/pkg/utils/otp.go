package utils

import (
	"fmt"

	"github.com/ishanshre/2FA-with-golang/api/v1/models"
	"github.com/pquerna/otp/totp"
)

func GenerateOtp(username string) (*models.GenerateOTP, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "ishan.com",
		AccountName: username,
		SecretSize:  15,
	})
	if err != nil {
		return nil, err
	}
	storeOTP := &models.GenerateOTP{
		Otp_secret:   key.Secret(),
		Otp_auth_url: key.URL(),
	}
	return storeOTP, nil
}

func ValidateOTP(otp *models.OTP) error {
	result := totp.Validate(otp.Token, otp.Otp_secret)
	if !result {
		return fmt.Errorf("token is invalid")
	}
	return nil
}
