API Documentation
Send OTP
Send a POST request to the /otp endpoint with the following body to send an OTP to a user's phone number

POST /otp

Request Body

{
  "phoneNumber": "<phone-number-with-country-code>"
}
curl -H "Content-Type: application/json" -X POST -d '{"phoneNumber": "+92**********"}' http://localhost:8000/otp
Be sure to include the country code in the phone number

Response

{
  "status": 202,
  "message": "success",
  "data": "OTP sent successfully"
}
Verify OTP
Verify a user's OTP by sending a POST request to the /verify endpoint with the following body that contains the phone number and the OTP code received by the user

POST /verifyOTP

Request Body

{
  "user": {
    "phoneNumber": "<phone-number-with-country-code>"
  },
  "code": "<code here>"
}
curl -H "Content-Type: application/json" -X POST -d '{"user": {"phoneNumber": "+92**********"}, "code":"795279"}' http://localhost:8000/verifyOTP
Response

{
  "status": 202,
  "message": "success",
  "data": "OTP verified successfully"
}
