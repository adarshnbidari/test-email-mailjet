package tasks

import (
	"fmt"

	"github.com/mailjet/mailjet-apiv3-go/v3"
)

// SendEmail for sending email using mailjet provider
func SendEmail(fromEmail string, toEmail string, subject string, body string, keys map[string]string) {

	mj := mailjet.NewMailjetClient(keys["public"], keys["private"])

	message := []mailjet.InfoMessagesV31{

		{
			From: &mailjet.RecipientV31{
				Email: fromEmail,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: toEmail,
				},
			},
			Subject:  subject,
			HTMLPart: body,
		},
	}

	messages := mailjet.MessagesV31{Info: message}

	_, err := mj.SendMailV31(&messages)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("message sent successfully to " + toEmail)

}
