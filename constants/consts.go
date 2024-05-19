package constants

import "flag"

var (
	CredFile string
)

type Credential struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type CredentialsList struct {
	Users []Credential `yaml:"users"`
}

func init() {
	flag.StringVar(&CredFile, "cred-file-name", "sample-creds.yaml", "Name of the file where credentials are stored.")
	flag.Parse()
}
