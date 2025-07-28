package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
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

	fmt.Printf(secret)
}
