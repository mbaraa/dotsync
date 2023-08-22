package actions

import (
	"bytes"
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

type UploadFilesAction struct {
	output io.Writer
}

func NewUploadFilesAction() IAction {
	return &UploadFilesAction{}
}

func (u *UploadFilesAction) Exec(output io.Writer, args ...any) error {
	u.output = output

	err := u.uploadFiles()
	if err != nil {
		return err
	}

	return nil
}

func (u *UploadFilesAction) NeedsRoot() bool {
	return false
}

func (u *UploadFilesAction) HasArgs() bool {
	return false
}

func (u *UploadFilesAction) uploadFiles() error {
	files, err := u.listFiles()
	if err != nil {
		return err
	}

	updateList := make([]struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}, 0)

	for _, filePath := range files {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}

		stat, err := file.Stat()
		if err != nil {
			return err
		}

		if stat.Size() > 256*1024 {
			return errors.New(fmt.Sprintf("file %s is larger than 256KiB...", filePath))
		}
		content, err := io.ReadAll(file)
		if err != nil {
			return err
		}

		encContent := base64.StdEncoding.EncodeToString(content)
		updateList = append(updateList, struct {
			Path    string `json:"path"`
			Content string `json:"content"`
		}{
			Path:    filePath,
			Content: encContent,
		})

	}

	fmt.Fprintln(u.output, "updating your synced files...")
	reqBody := bytes.NewBuffer([]byte{})
	_ = json.StringifyToWriter(reqBody, map[string]any{
		"files": updateList,
	})

	req, err := http.NewRequest("POST", config.ServerAddress+"/file/upload", reqBody)
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
	fmt.Fprintln(u.output, "done 👍")

	return nil
}

func (u *UploadFilesAction) listFiles() ([]string, error) {
	req, err := http.NewRequest("GET", config.ServerAddress+"/file", nil)
	token, err := configfile.GetValue("token")
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", token)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		respBody, _ := json.ParseFromReader[json.Json](resp.Body)
		resp.Body.Close()
		return nil, errors.New(respBody["error"].(string))
	}

	files, err := json.ParseFromReader[[]string](resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return files, nil

}
