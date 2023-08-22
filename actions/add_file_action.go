package actions

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	if stat.Size() > 256*1024 {
		return errors.New("file is larger than 256KiB...")
	}

	fmt.Fprintf(a.output, "Adding %s to your synced files...\n", filePath)

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	filePath, err = filepath.Abs(filePath)
	if err != nil {
		return err
	}

	encContent := base64.StdEncoding.EncodeToString(content)

	reqBody := bytes.NewBuffer([]byte{})
	_ = json.StringifyToWriter(reqBody, map[string]string{
		"path":    filePath,
		"content": encContent,
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

	if resp.StatusCode != 200 {
		respBody, _ := json.ParseFromReader[json.Json](resp.Body)
		resp.Body.Close()
		return errors.New(respBody["error"].(string))
	}

	fmt.Fprintln(a.output, "Done, you can check the synced files list by using `dotsync -list`")

	return nil
}
