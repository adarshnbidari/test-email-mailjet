package cmd

import (
	"fmt"

	"email/tasks"

	"github.com/spf13/cobra"
)

var fromEmail string = ""
var toEmail string = ""
var public string = ""
var private string = ""
var body string = ""
var subject string = ""

var sendEmailCmd = &cobra.Command{

	Use:   "Send email",
	Short: "Send email",
	Run: func(cmd *cobra.Command, args []string) {

		if fromEmail == "" {

			fmt.Println("Email cannot be empty")

		}

		if toEmail == "" {

			fmt.Println("toEmail cannot be empty")

		}

		if body == "" {
			fmt.Println("email body cannot be empty")

		}

		if subject == "" {
			fmt.Println("email subject cannot be empty")

		}

		if public == "" || private == "" {
			public = ""  // specify default public key here
			private = "" // specify default private key here
		}

		keys := map[string]string{
			"public":  public,
			"private": private,
		}

		tasks.SendEmail(fromEmail, toEmail, subject, body, keys)

	},
}

func init() {

	rootCmd.AddCommand(sendEmailCmd)

	sendEmailCmd.Flags().StringVarP(&fromEmail, "fromEmail", "f", "", "fromEmail")

	sendEmailCmd.Flags().StringVarP(&toEmail, "toEmail", "t", "", "toEmail")

	sendEmailCmd.Flags().StringVarP(&public, "public", "p", "", "public")

	sendEmailCmd.Flags().StringVarP(&private, "private", "s", "", "private")

	sendEmailCmd.Flags().StringVarP(&body, "body", "b", "", "body")

	sendEmailCmd.Flags().StringVarP(&subject, "subject", "h", "", "subject")

	sendEmailCmd.MarkFlagRequired("fromEmail")

	sendEmailCmd.MarkFlagRequired("toEmail")

	sendEmailCmd.MarkFlagRequired("body")

	sendEmailCmd.MarkFlagRequired("subject")

}
