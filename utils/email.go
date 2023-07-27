package utils

import (
	"crypto/rand"
	"io"
	"fmt"

	"github.com/Caknoooo/golang-clean_template/config"
	"github.com/Caknoooo/golang-clean_template/dto"

	"gopkg.in/gomail.v2"
)

func SendMail(toEmail string, subject string, body string) error {
	emailConfig, err := config.NewEmailConfig()
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailConfig.AuthEmail)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		emailConfig.Host, 
		emailConfig.Port, 
		emailConfig.AuthEmail, 
		emailConfig.AuthPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

func MakeVerificationEmail(receiverEmail string) (map[string]string, error) {
	token := EncodeToString(6)
	if token == "" {
		return nil, dto.ErrorFailedGenerateVerificationCode
	}

	draftEmail := map[string]string{}
	draftEmail["subject"] = "Quliku - Email Verification"
	draftEmail["code"] = token
	draftEmail["body"] = fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Email Verification - Quliku</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					line-height: 1.6;
				}
				.code-container {
					text-align: center;
				}
				.code {
					font-size: 30px;
					padding: 10px;
					display: inline-block;
					margin: 0 10px; /* Add margin between each digit */
				}
				.note {
					font-size: 14px;
					color: #888;
				}
			</style>
		</head>
		<body>
			<p>Hi, %s! Thanks for registering an account on Quliku.App.</p>
			<p>Please verify your email address by entering the code below:</p>	
			<div class="code-container">
				<p class="code">%s</p>
			</div>
			<p class="note">Please note that this code will expire after 3 minutes.</p>
			<p>Thanks,<br>Quliku Team</p>
		</body>
		</html>
	`, receiverEmail, token)

	return draftEmail, nil
}

func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, max)
	n, _ := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return dto.ErrorFailedGenerateVerificationCode.Error()
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i]) % len(table)]
	}
 
	return string(b)
}