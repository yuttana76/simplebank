package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuttana76/simbplebank/util"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip() // skip test when run with make test
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A Test Email"
	content := `
	<h1>Hello</h1>
	<p>This is a test email from Simple Bank.</p>
	`
	to := []string{"yuttana76@gmail.com"}
	attachFiles := []string{"../README.md"}
	err = sender.SendEmail(
		subject,
		content,
		config.EmailSenderAddress,
		to,
		nil,
		nil,
		attachFiles,
	)
	require.NoError(t, err)

}
