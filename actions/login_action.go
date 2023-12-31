package actions

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/mbaraa/dotsync/config"
	"github.com/mbaraa/dotsync/utils/configfile"
	"github.com/mbaraa/dotsync/utils/json"
)

type LoginAction struct {
	output io.Writer
}

func NewLoginAction() IAction {
	return &LoginAction{}
}

func (l *LoginAction) Exec(output io.Writer, args ...any) error {
	l.output = output

	err := l.login(args[0].(string))
	if err != nil {
		return err
	}

	return nil
}

func (l *LoginAction) NeedsRoot() bool {
	return false
}

func (l *LoginAction) HasArgs() bool {
	return true
}

func (l *LoginAction) login(email string) error {
	fmt.Fprintf(l.output, "Sending an email with the login token to %s.\n", email)

	reqBody := bytes.NewBuffer([]byte{})
	_ = json.StringifyToWriter(reqBody, map[string]string{
		"email": email,
	})
	resp, err := http.Post(config.ServerAddress+"/user/login", "application/json", reqBody)
	if err != nil {
		return err
	}

	if resp.StatusCode == 400 {
		return errors.New("invalid email address")
	}
	if resp.StatusCode != 200 {
		return errors.New("something went wrong...")
	}

	respBody, err := json.ParseFromReader[json.Json](resp.Body)
	if err != nil {
		return err
	}

	var token string
	if tok, ok := respBody["token"].(string); ok {
		token = tok
	}

	fmt.Fprintf(l.output, "An email was sent to %s, with the login token.\n", email)
	fmt.Fprintln(l.output, "Copy the token and paste it here to complete the login process.")

	return l.storeToken(token)
}

func (l *LoginAction) storeToken(midPart string) error {
	fmt.Fprint(l.output, "\nEnter the login token: ")
	var loginToken string
	fmt.Scanln(&loginToken)

	tokenParts := strings.Split(loginToken, "🔒")
	if len(tokenParts) != 2 {
		return errors.New("invalid token")
	}
	loginToken = fmt.Sprintf("%s.%s.%s", tokenParts[0], midPart, tokenParts[1])

	fmt.Fprintln(l.output, "\nChecking your token, hope you're using your token")
	reqBody := bytes.NewBuffer([]byte{})
	_ = json.StringifyToWriter(reqBody, map[string]string{
		"token": loginToken,
	})
	resp, err := http.Post(config.ServerAddress+"/user/login/verify", "application/json", reqBody)
	if err != nil {
		return err
	}

	time.Sleep(time.Second)

	if resp.StatusCode == 401 {
		return errors.New("invalid or expired token")
	}
	if resp.StatusCode != 200 {
		return errors.New("something went wrong...")
	}

	fmt.Fprintln(l.output, "seems legit, carry on...")
	fmt.Fprintln(l.output, "check -help, or the official docs at https://dotsync.org/docs, for a detailed usage!")

	respBody, err := json.ParseFromReader[json.Json](resp.Body)
	if err != nil {
		return err
	}

	var token string
	if tok, ok := respBody["token"].(string); ok {
		token = tok
	}

	err = configfile.SetValue("token", token)
	if err != nil {
		return err
	}

	return nil
}
