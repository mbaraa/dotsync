package actions

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mbaraa/dotsync/config"
	"github.com/mbaraa/dotsync/utils/configfile"
	"github.com/mbaraa/dotsync/utils/json"
)

type DownloadFilesAction struct {
	output io.Writer
}

func NewDownloadFilesAction() IAction {
	return &DownloadFilesAction{}
}

func (d *DownloadFilesAction) Exec(output io.Writer, args ...any) error {
	d.output = output

	err := d.downloadFiles()
	if err != nil {
		return err
	}

	return nil
}

func (d *DownloadFilesAction) NeedsRoot() bool {
	return false
}

func (d *DownloadFilesAction) HasArgs() bool {
	return false
}

func (d *DownloadFilesAction) downloadFiles() error {
	fmt.Fprintln(d.output, "downloading your dotfiles...")
	req, err := http.NewRequest("GET", config.ServerAddress+"/file/download", nil)
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

	files, err := json.ParseFromReader[[]struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}](resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	fmt.Fprintln(d.output, "done 👍")

	fmt.Fprintln(d.output, "\nupdating your local dotfiles...")
	for _, file := range files {
		fmt.Fprintf(d.output, "saving file %s...", file.Path)

		err = os.Truncate(file.Path, 0)
		if err != nil {
			fmt.Fprintln(d.output, err.Error())
			continue
		}

		f, err := os.OpenFile(file.Path, os.O_RDWR, 0755)
		if err != nil {
			fmt.Fprintf(d.output, "opening file %s failed, reason: %s", file.Path, err.Error())
			continue
		}

		decContent, _ := base64.StdEncoding.DecodeString(file.Content)
		_, err = f.Write(decContent)
		if err != nil {
			fmt.Fprintf(d.output, "saving file %s failed, reason: %s", file.Path, err.Error())
			continue
		}
		fmt.Fprintln(d.output, "Done 👍")
	}
	fmt.Fprintln(d.output, "All is done 👍")
	fmt.Fprintln(d.output, "Some of your programs needs a restart after the download, just saying 😁")

	return nil
}
