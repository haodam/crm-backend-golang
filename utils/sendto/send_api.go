package sendto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// gửi yêu cầu API HTTP POST đến một dịch vụ email, nhằm gửi email với mã OTP

type MailRequest struct {
	ToEmail     string `json:"to_email"`
	Subject     string `json:"subject"`
	MessageBody string `json:"message_body"`
	Attachment  string `json:"attachment"`
}

func SendMailToJavaByAPI(otp string, email string, purpose string) error {

	// URL API
	postURL := "http://localhost:8080/email/send_text"

	mailRequest := MailRequest{
		ToEmail:     email,
		Subject:     "Verify OTP " + purpose,
		MessageBody: "OTP is " + otp,
		Attachment:  "path/to/email",
	}

	// convert struct to json
	requestBody, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}

	// create header
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// PUT header
	req.Header.Set("Content-Type", "application/json")

	// execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)

	return nil
}
