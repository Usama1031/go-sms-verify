package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/usama1031/go-sms-verify/data"
)

const appTimeout = time.Second * 10

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		var payload data.OTPData

		app.validateBody(c, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, "OTP Send Successfully!")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		var payload data.VerifyData

		app.validateBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		fmt.Println("err: ", err)

		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully!")
	}

}
