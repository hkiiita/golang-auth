package utils

import (
	"github.com/hkbiet/golang-basic-auth/constants"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func readYaml() (constants.CredentialsList, error) {
	var credList constants.CredentialsList
	file, err := os.Open(constants.CredFile)
	if err != nil {
		return credList, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return credList, err
	}
	err = yaml.Unmarshal(data, &credList)
	if err != nil {
		return credList, err
	}
	return credList, nil
}

func GetCreds() map[string]string {
	credentials := make(map[string]string)
	credList, err := readYaml()
	if err != nil {
		return nil
	}
	for _, item := range credList.Users {
		credentials[item.Username] = item.Password
	}
	return credentials
}
