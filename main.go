package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
	"github.com/pquerna/otp/totp"
	"github.com/atotto/clipboard"
)

func configPath() string {
	configHome, _ := os.UserHomeDir()
	return filepath.Join(configHome, ".config", "gotp", "secrets.yaml")
}

func mustLoadSecrets() map[string]string {
	path := configPath()
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading %s: %v", path, err)
	}

	secrets := make(map[string]string)

	err = yaml.Unmarshal(data, &secrets)
	if err != nil {
		log.Fatalf("error while parsing YAML: %v", err)
	}

	return secrets
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gotp <key>")
		os.Exit(1)
	}

	key := os.Args[1]

	secrets := mustLoadSecrets();

	secret, found := secrets[key]
	if !found {
		log.Fatalf("%s not found", key)
	}

	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		log.Fatalf("Fehler beim Generieren des TOTP-Codes: %v", err)
	}

	err = clipboard.WriteAll(code)
	if err != nil {
		fmt.Printf(code)
	}

}
