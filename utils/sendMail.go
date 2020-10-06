package utils

import (
	"context"
	"strconv"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func SendVefiryMail(verifyCode, reciver string) error {
	config := GetConfig().Mailgun

	if config.SkipSend {
		return nil
	}

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(config.Domain, config.ApiKey)

	message := mg.NewMessage(config.Sender, "TeamWork 이메일주소 인증", "", reciver)
	message.SetTemplate("teamwork_emailverify")
	message.AddTemplateVariable("link", "https://example.com/"+verifyCode)
	message.AddTemplateVariable("year", strconv.Itoa(time.Now().Year()))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mg.Send(ctx, message)

	if err != nil {
		return err
	}

	return nil
}

func SendResetPasswordMail(code, reciver string) error {
	config := GetConfig().Mailgun

	if config.SkipSend {
		return nil
	}

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(config.Domain, config.ApiKey)

	message := mg.NewMessage(config.Sender, "TeamWork 비밀번호 변경", "", reciver)
	message.SetTemplate("teamwork_emailverify")
	message.AddTemplateVariable("link", "https://example.com/"+code)
	message.AddTemplateVariable("year", strconv.Itoa(time.Now().Year()))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mg.Send(ctx, message)

	if err != nil {
		return err
	}

	return nil
}
