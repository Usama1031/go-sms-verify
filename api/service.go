package api

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envAccountSID(),
	Password: envAuthToken(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(envServicesSID(), params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	res, err := client.VerifyV2.CreateVerificationCheck(envServicesSID(), params)

	if err != nil {
		return err
	}

	if res.Status == nil || *res.Status != "approved" {
		return errors.New("not a valid code")
	}
	return nil
}
