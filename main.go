package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gopkg.in/yaml.v3"
	"github.com/pquerna/otp/totp"
	"github.com/atotto/clipboard"
)

func main() {
	list := flag.Bool("list", false, "list all available keys")
	flag.Parse()

	if !*list && len(flag.Args()) != 1 {
		fmt.Println("Usage:\ngotp --list - lists all available keys\ngotp <key>  - copys a token for <key> to the clipboard\n")
		os.Exit(1)
	}

	if *list {
		printKeys()
		return
	}

	key := flag.Args()[0]
	copyToken(key)
}

func printKeys() {
	secrets := mustLoadSecrets();

	var keys []string
	for k := range secrets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k)
	}
}

func copyToken(key string) {
	secrets := mustLoadSecrets();

	secret, found := secrets[key]
	if !found {
		log.Fatalf("%s not found", key)
	}

	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		log.Fatalf("error while generating token: %v", err)
	}

	err = clipboard.WriteAll(code)
	if err != nil {
		fmt.Printf(code)
	}
}

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
