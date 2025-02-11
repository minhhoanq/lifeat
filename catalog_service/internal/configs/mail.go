package configs

type Mail struct {
	EmailSenderName     string `yaml:"email_sender_name"`
	EmailSenderAddress  string `yaml:"email_sender_address"`
	EmailSenderPassword string `yaml:"email_sender_password"`
}
