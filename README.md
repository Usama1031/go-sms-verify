**📱 API Documentation**

**🔐 Send OTP**

**Endpoint:**

POST /otp

**Description:**

Send a POST request to the /otp endpoint to send an OTP to a user's phone number.

**Request Body:**

{
  "phoneNumber": "<phone-number-with-country-code>"
}

**Example cURL Command:**


curl -H "Content-Type: application/json" -X POST -d '{"phoneNumber": "+92**********"}' http://localhost:8000/otp

⚠️ Make sure to include the country code in the phone number.

**Response:**

{
  "status": 202,
  "message": "success",
  "data": "OTP sent successfully"
}

**✅ Verify OTP**

**Endpoint:**

POST /verifyOTP

**Description:**

Verify a user's OTP by sending a POST request containing the phone number and the received OTP code.

**Request Body:**

{
  "user": {
    "phoneNumber": "<phone-number-with-country-code>"
  },
  "code": "<code here>"
}

**Example cURL Command:**

curl -H "Content-Type: application/json" -X POST -d '{"user": {"phoneNumber": "+92**********"}, "code":"795279"}' http://localhost:8000/verifyOTP

**Response:**

{
  "status": 202,
  "message": "success",
  "data": "OTP verified successfully"
}
