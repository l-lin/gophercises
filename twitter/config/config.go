package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/l-lin/gophercises/twitter/auth"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

const (
	twitterAppDashboardURL = "https://developer.twitter.com/en/apps"
	keyKey                 = "key"
	secretKeyKey           = "secret-key"
	tokenKey               = "token"
)

// InitTwitterAPIKeys writes twitter consumer API keys in config file
func InitTwitterAPIKeys(cfgFile string) {
	fmt.Println("Create a new twitter app in your dashboard and get the API keys:", aurora.BrightCyan(twitterAppDashboardURL))
	openBrowser(twitterAppDashboardURL)
	prompt := promptui.Prompt{
		Label:    "Consumer API key: ",
		Validate: validate,
		Mask:     '*',
	}
	key, err := prompt.Run()
	if err != nil {
		log.Fatalln(aurora.BrightRed(err))
	}

	prompt = promptui.Prompt{
		Label:    "Consumer API secret key: ",
		Validate: validate,
		Mask:     '*',
	}
	secretKey, err := prompt.Run()
	if err != nil {
		log.Fatalln(aurora.BrightRed(err))
	}

	a := auth.TwitterAuthenticator{}
	resultCh := make(chan *auth.Result)
	go a.Authenticate(resultCh, &auth.Credentials{UserName: key, Password: secretKey})
	result := <-resultCh
	if result.Error != nil {
		log.Fatalln(aurora.BrightRed(result.Error))
	}
	cfgContent := []byte(fmt.Sprintf("%s: %s\n%s: %s\n%s: %s", keyKey, key, secretKeyKey, secretKey, tokenKey, result.Token))
	err = ioutil.WriteFile(cfgFile, cfgContent, 0644)

	if err != nil {
		log.Fatalln(aurora.BrightRed(err))
	}
}

// GetAPIKey configured from the config file
func GetAPIKey() string {
	return viper.GetString(keyKey)
}

// GetAPISecretKey configured from the config file
func GetAPISecretKey() string {
	return viper.GetString(secretKeyKey)
}

// GetAPIAuthToken configured from the config file
func GetAPIAuthToken() string {
	return viper.GetString(tokenKey)
}

func validate(input string) error {
	if len(input) < 1 {
		return errors.New("The consumer API key cannot be empty")
	}
	return nil
}

func openBrowser(url string) error {
	browser := "xdg-open"

	args := []string{url}
	if runtime.GOOS == "windows" {
		browser = "rundll32.exe"
		args = []string{"url.dll,FileProtocolHandler", url}
	} else if runtime.GOOS == "darwin" {
		browser = "open"
		args = []string{url}
	} else if runtime.GOOS == "plan9" {
		browser = "plumb"
	}
	browser, err := exec.LookPath(browser)
	if err == nil {
		cmd := exec.Command(browser, args...)
		cmd.Stderr = os.Stderr
		err = cmd.Start()
		if err != nil {
			return fmt.Errorf("Cannot start command: %v", err)
		}
	}
	return nil
}
