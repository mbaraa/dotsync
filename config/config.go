package config

import "os/user"

func init() {
	u, _ := user.Current()
	ConfigFilePath = u.HomeDir + "/.dotsyncrc"
}

var (
	ServerAddress  = "https://api.dotsync.org"
	ConfigFilePath = "~/.dotsyncrc"
)
