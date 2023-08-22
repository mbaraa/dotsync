package actions

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mbaraa/dotsync/config"
	"github.com/mbaraa/dotsync/utils/configfile"
	"github.com/mbaraa/dotsync/utils/json"
)

type DeleteUserAction struct {
	output io.Writer
}

func NewDeleteUserAction() IAction {
	return &DeleteUserAction{}
}

func (d *DeleteUserAction) Exec(output io.Writer, args ...any) error {
	d.output = output

	err := d.deleteUser()
	if err != nil {
		return err
	}

	return nil
}

func (d *DeleteUserAction) NeedsRoot() bool {
	return false
}

func (d *DeleteUserAction) HasArgs() bool {
	return false
}

func (d *DeleteUserAction) deleteUser() error {
	fmt.Fprint(d.output, "Are you sure that you want to delete your user and remote file? [y|N] ")
	var choice string
	fmt.Scanln(&choice)
	switch strings.ToLower(choice) {
	case "Y", "y", "yes":
		break
	case "N", "n", "no":
		return errors.New("user canceled")
	default:
		return errors.New("invalid choice")
	}

	fmt.Fprintln(d.output, "deleting your user...")
	req, err := http.NewRequest("DELETE", config.ServerAddress+"/user", nil)
	token, err := configfile.GetValue("token")
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", token)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		respBody, _ := json.ParseFromReader[json.Json](resp.Body)
		resp.Body.Close()
		return errors.New(respBody["error"].(string))
	}

	fmt.Fprintln(d.output, "Sad to see you go...")
	fmt.Fprintln(d.output, "If there's anything that bothered you in Dotsync feel free to contact me at dotsync@mbaraa.com")

	return nil
}
