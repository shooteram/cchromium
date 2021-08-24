//go:build windows
// +build windows

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	defaultConfiguration = `
chromium:
  path: ""
  settings:
    proxy: ""
    proxy_bypass: []
    host_resolver_rules: ""
    user_agent: ""
    disable_features: []
`

	configDir    string
	configPath   string
	shortcutDir  string
	shortcutPath string
)

type Options struct {
	Chromium struct {
		Path     string `yaml:"path"`
		Settings struct {
			Proxy             string   `yaml:"proxy"`
			ProxyBypass       []string `yaml:"proxy_bypass"`
			HostResolverRules string   `yaml:"host_resolver_rules"`
			UserAgent         string   `yaml:"user_agent"`
			DisableFeatures   []string `yaml:"disable_features"`
		} `yaml:"settings"`
	} `yaml:"chromium"`
}

func main() {
	if runtime.GOOS != "windows" {
		log.Fatalf("%s is not supported by this software", runtime.GOOS)
	}

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	configDir = filepath.Join(userConfigDir, "shooteram")
	configPath = filepath.Join(configDir, "config.yaml")
	shortcutDir = filepath.Join(configDir, "Shortcuts")
	shortcutPath = filepath.Join(shortcutDir, "Chromium.lnk")

	chromium := loadConfiguration()
	chromium.createShortcut()
}

func loadConfiguration() *Options {
	o := Options{}

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.Mkdir(configDir, 0o755)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err := yaml.Unmarshal([]byte(defaultConfiguration), &o)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		config, err := yaml.Marshal(&o)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if err := os.WriteFile(configPath, config, 0655); err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Printf("Your configuration has been written to the following file: %s\n"+
			"Please, specify the location of your Chromium executable file "+
			`(usually "C:\Program Files\Chromium\Application\chrome.exe")`,
			configPath)
		os.Exit(0)
	}

	config, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal(config, &o)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if o.Chromium.Path == "" {
		fmt.Printf("To continue, specify a path to the Chromium executable in your configuration file: %s\n", configPath)
		os.Exit(1)
	}

	return &o
}

func (o *Options) createShortcut() {
	if _, err := os.Stat(shortcutDir); os.IsNotExist(err) {
		os.MkdirAll(shortcutDir, 0o755)
	}
	if _, err := os.Stat(shortcutPath); os.IsNotExist(err) {
		os.Remove(shortcutPath)
	}

	args := []string{}
	if o.Chromium.Settings.Proxy != "" {
		args = append(args, fmt.Sprintf("--proxy-server=`\"%s`\"", o.Chromium.Settings.Proxy))
	}
	if len(o.Chromium.Settings.ProxyBypass) > 0 {
		args = append(args, fmt.Sprintf("--proxy-bypass-list=`\"%s`\"", strings.Join(o.Chromium.Settings.ProxyBypass, ";")))
	}
	if o.Chromium.Settings.HostResolverRules != "" {
		args = append(args, fmt.Sprintf("--host-resolver-rules=`\"%s`\"", o.Chromium.Settings.HostResolverRules))
	}
	if o.Chromium.Settings.UserAgent != "" {
		args = append(args, fmt.Sprintf("--user-agent=`\"%s`\"", o.Chromium.Settings.UserAgent))
	}
	if len(o.Chromium.Settings.DisableFeatures) > 0 {
		args = append(args, fmt.Sprintf("--disable-features=`\"%s`\"", strings.Join(o.Chromium.Settings.DisableFeatures, ",")))
	}

	powershellCommand := fmt.Sprintf(
		"$WshShell = New-Object -comObject WScript.Shell;"+
			`$Shortcut = $WshShell.CreateShortcut("%s");`+
			`$Shortcut.TargetPath = "%s";`+
			`$Shortcut.Arguments = "%s";`+
			`$Shortcut.Save()`,
		shortcutPath,
		o.Chromium.Path,
		strings.Join(args, " "))
	cmd := exec.Command("powershell.exe", powershellCommand)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your configured shortcut to Chromium is available in the following folder: %s\n", shortcutDir)
	fmt.Println("Now you can, for example, use this shortcut to pin it to the taskbar or to the start menu")
	os.Exit(0)
}
