package sendto

type MailRequest struct {
	ToEmail     string `json:"to_email"`
	Subject     string `json:"subject"`
	MessageBody string `json:"message_body"`
	Attachment  string `json:"attachment"`
}

func SendMailToJavaByAPI(otp string, email string, purpose string) error {

	// URL API
	postURL, err := "http://localhost:8080/email/send_text"
	if err != nil {
		return err
	}

	return nil
}
