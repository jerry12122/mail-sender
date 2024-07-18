//go:build linux || windows

package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile    string
	smtpServer    string
	smtpPort      int
	smtpUser      string
	smtpPass      string
	senderEmail   string
	receiverEmail string
	emailSubject  string
	emailBody     string
	useStartTLS   bool
)

var rootCmd = &cobra.Command{
	Use:   "mail-sender",
	Short: "A CLI tool to send emails using SMTP",
	Run: func(cmd *cobra.Command, args []string) {
		sendEmail()
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display help for the mail-sender command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mail-sender is a CLI tool to send emails using SMTP.")
		fmt.Println("\nUsage:")
		fmt.Println("  mail-sender [flags]")
		fmt.Println("\nFlags:")
		cmd.Flags().PrintDefaults()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Path to the config file")
	rootCmd.Flags().StringVarP(&smtpServer, "host", "H", "", "SMTP host address")
	rootCmd.Flags().IntVarP(&smtpPort, "port", "P", 0, "SMTP port number")
	rootCmd.Flags().StringVarP(&smtpUser, "user", "u", "", "SMTP username")
	rootCmd.Flags().StringVarP(&smtpPass, "password", "p", "", "SMTP password")
	rootCmd.Flags().StringVarP(&senderEmail, "from", "f", "", "Sender's email address")
	rootCmd.Flags().StringVarP(&receiverEmail, "to", "t", "", "Receiver's email address")
	rootCmd.Flags().StringVarP(&emailSubject, "subject", "s", "", "Email subject")
	rootCmd.Flags().StringVarP(&emailBody, "body", "b", "", "Email body")
	rootCmd.Flags().BoolVarP(&useStartTLS, "use-starttls", "", false, "Use STARTTLS encryption")

	rootCmd.MarkFlagRequired("from")
	rootCmd.MarkFlagRequired("to")
	rootCmd.MarkFlagRequired("subject")
	rootCmd.MarkFlagRequired("body")

	rootCmd.AddCommand(helpCmd)
}

func sendEmail() {
	// Load config from file if provided
	if configFile != "" {
		viper.SetConfigFile(configFile)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}

		smtpServer = viper.GetString("smtp.host")
		smtpPort = viper.GetInt("smtp.port")
		smtpUser = viper.GetString("smtp.user")
		smtpPass = viper.GetString("smtp.password")
		useStartTLS = viper.GetBool("smtp.use_starttls")
	}
	// Create the message
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", emailSubject, emailBody))

	// Authenticate with the SMTP server
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpServer)

	// Send the email
	var err error
	if useStartTLS {
		err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, senderEmail, []string{receiverEmail}, message)
	} else {
		err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, senderEmail, []string{receiverEmail}, message)
	}
	if err != nil {
		fmt.Println("Error sending email:", err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully!")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
