package actions

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mbaraa/dotsync/config"
	"github.com/mbaraa/dotsync/utils/configfile"
	"github.com/mbaraa/dotsync/utils/json"
)

type RemoveFileAction struct {
	output io.Writer
}

func NewRemoveFileAction() IAction {
	return &RemoveFileAction{}
}

func (r *RemoveFileAction) Exec(output io.Writer, args ...any) error {
	r.output = output

	err := r.removeFile(args[0].(string))
	if err != nil {
		return err
	}

	return nil
}

func (r *RemoveFileAction) NeedsRoot() bool {
	return false
}

func (r *RemoveFileAction) HasArgs() bool {
	return true
}

func (r *RemoveFileAction) removeFile(filePath string) error {
	_, err := os.Open(filePath)
	if err != nil {
		return err
	}

	fmt.Fprintf(r.output, "Remove '%s' from your synced files? [y|N] ", filePath)
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

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/file?path=%s", config.ServerAddress, filePath), nil)
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

	fmt.Fprintln(r.output, "Done, you can check the synced files list by using `dotsync -list`")

	return nil
}
