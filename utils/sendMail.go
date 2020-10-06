package utils

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func SendVefiryMail(verifyCode, reciver string) error {
	config := GetConfig().Mailgun

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(config.Domain, config.ApiKey)

	message := mg.NewMessage(config.Sender, "TeamWork 이메일주소 인증", "", reciver)
	message.SetTemplate("teamwork_emailverify")
	message.AddTemplateVariable("link", "https://example.com/"+verifyCode)
	message.AddTemplateVariable("year", strconv.Itoa(time.Now().Year()))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
