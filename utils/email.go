package utils

import (
	"crypto/rand"
	"fmt"
	"io"

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

func MakeForgotPasswordEmail(receiverEmail string) (map[string]string, error) {
	token := EncodeToString(6)
	if token == "" {
		return nil, dto.ErrorFailedGenerateVerificationCode
	}

	draftEmail := map[string]string{}
	draftEmail["subject"] = "Quliku - Forgot Password"
	draftEmail["code"] = token
	draftEmail["body"] = fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Password Reset - Quliku</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					line-height: 1.6;
				}
				.instructions {
					margin-top: 20px;
					margin-bottom: 30px;
				}
				.code-container {
					text-align: center;
				}
				.button {
					display: inline-block;
					padding: 10px 20px;
					background-color: #007bff;
					color: white;
					text-decoration: none;
					border-radius: 5px;
				}
				.note {
					font-size: 14px;
					color: #888;
				}
				.code {
					font-size: 30px;
					padding: 10px;
					display: inline-block;
					margin: 0 10px; /* Add margin between each digit */
				}
			</style>
		</head>
		<body>
			<p>Hi, %s!</p>
			<p>We received a request to reset your password for your Quliku.App account.</p>
			<p>If you didn't make this request, you can ignore this email.</p>
			<div class="instructions">
				<p>To reset your password, entering the code below:</p>
			</div>
			<div class="code-container">
				<p class="code">%s</p>
			</div>
			<p class="note">This link is valid for 3 minutes.</p>
			<p>If you have any questions, please contact our support team.</p>
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
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}
