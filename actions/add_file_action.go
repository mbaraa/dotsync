package actions

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mbaraa/dotsync/config"
	"github.com/mbaraa/dotsync/utils/configfile"
	"github.com/mbaraa/dotsync/utils/json"
)

type AddFileAction struct {
	output io.Writer
}

func NewAddFileAction() IAction {
	return &AddFileAction{}
}

func (a *AddFileAction) Exec(output io.Writer, args ...any) error {
	a.output = output

	err := a.addFile(args[0].(string))
	if err != nil {
		return err
	}

	return nil
}

func (a *AddFileAction) NeedsRoot() bool {
	return false
}

func (a *AddFileAction) HasArgs() bool {
	return true
}

func (a *AddFileAction) addFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	fmt.Fprintf(a.output, "Adding %s to your synced files...\n", filePath)

	reqBody := bytes.NewBuffer([]byte{})
	_ = json.StringifyToWriter(reqBody, map[string]string{
		"path":    filePath,
		"content": string(content),
	})

	req, err := http.NewRequest("POST", config.ServerAddress+"/file", reqBody)
	req.Header.Add("Content-Type", "application/json")
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

	// TODO: handle different occuring errors
	if resp.StatusCode != 200 {
		return errors.New("something went wrong...")
	}

	fmt.Fprintln(a.output, "Done, you can check the synced files list by using `dotsync -list`")

	return nil
}
