package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
	"github.com/spf13/cobra"
)

type Account struct {
	Name        string `ini:"name"`
	Email       string `ini:"email"`
	SSHCommand  string `ini:"sshCommand"`
	Description string `ini:"description,optional"`
}

type GitConfig struct {
	Accounts map[string]Account
}

func getTutConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err) // Handle the error appropriately
	}
	// if the .tut/.ini file does not exist, create it
	if _, err := os.Stat(filepath.Join(homeDir, ".tut", ".ini")); os.IsNotExist(err) {
		fmt.Println("Creating .tut directory and .ini file under $HOME...")
		os.MkdirAll(filepath.Join(homeDir, ".tut"), 0755)
		os.Create(filepath.Join(homeDir, ".tut", ".ini"))
	}
	return filepath.Join(homeDir, ".tut", ".ini")
}

func getLocalGitConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err) // Handle the error appropriately
	}
	return filepath.Join(homeDir, ".tut", ".gitconfig")
}

// Function to load and parse the INI config
func loadConfig() (GitConfig, error) {
	cfg, err := ini.Load(getTutConfigPath())
	if err != nil {
		return GitConfig{}, err
	}
	return parseConfig(cfg)
}

// Function to parse the loaded INI data into the data structure
func parseConfig(cfg *ini.File) (GitConfig, error) {
	config := GitConfig{Accounts: make(map[string]Account)}

	for _, section := range cfg.Sections() {
		sectionName := section.Name()
		if !strings.HasPrefix(sectionName, "account.") {
			continue
		}

		Shortcut := strings.TrimPrefix(sectionName, "account.")
		config.Accounts[Shortcut] = Account{
			Name:        section.Key("name").String(),
			Email:       section.Key("email").String(),
			SSHCommand:  section.Key("sshCommand").String(),
			Description: section.Key("description").String(),
		}
	}

	return config, nil
}

func createLocalConfig(account Account) error {
	content := []byte(fmt.Sprintf(
		"[user]\n\tname = %s\n\temail = %s\n[core]\n\tsshCommand = %s\n",
		account.Name, account.Email, account.SSHCommand))

	// Create the .gitconfig file (overwrites existing ones)
	configPath := getLocalGitConfigPath()
	err := os.WriteFile(configPath, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func configureGit(selectedAccount Account) {
	// Assuming the .gitconfig file should be in the current directory
	err := createLocalConfig(selectedAccount)
	if err != nil {
		// Handle the error
		fmt.Println("Error creating local configuration:", err)
	}
}

func loadLiveConfig() (Account, error) {
	cfg, err := ini.Load(getLocalGitConfigPath())
	if err != nil {
		return Account{}, err
	}
	// Create Account struct from cfg
	account := Account{
		Name:       cfg.Section("user").Key("name").String(),
		Email:      cfg.Section("user").Key("email").String(),
		SSHCommand: cfg.Section("core").Key("sshCommand").String(),
	}
	return account, nil
}

func interactive(config GitConfig) {
	liveConfig, err := loadLiveConfig()
	if err != nil {
		fmt.Println("Error loading live config:", err)
	}
	for shortcut, account := range config.Accounts {
		if account == liveConfig {
			fmt.Printf("\033[32m[%s] %s (%s)\033[0m\n", shortcut, account.Name, account.Email) // Yellow text
		} else {
			fmt.Printf("[%s] %s (%s)\n", shortcut, account.Name, account.Email)
		}
	}
	// if user input q or quit, exit. Else check shortcut and configure git
	fmt.Println("--------------------")
	fmt.Print("Select an account: ")
	userInput := ""
	fmt.Scanln(&userInput)
	if userInput == "q" || userInput == "quit" {
		os.Exit(0)
	}
	if selectedAccount, ok := config.Accounts[userInput]; ok {
		configureGit(selectedAccount)
	} else {
		fmt.Println("Invalid account selected")
	}
}

func list(config GitConfig) {
	for shortcut, account := range config.Accounts {
		fmt.Printf("[%s] %s (%s)\n", shortcut, account.Name, account.Email)
	}
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	rootCmd := &cobra.Command{
		Use:   "tut",
		Short: "A tool to manage multiple Github accounts",
		Run: func(cmd *cobra.Command, args []string) {
			interactive(config)
		},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List configured accounts",
		Run: func(cmd *cobra.Command, args []string) {
			list(config)
		},
	}

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new account",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Adding account...") // Replace with interactive input
		},
	}

	editCmd := &cobra.Command{
		Use:   "edit <account_name>",
		Short: "Edit an account",
		Args:  cobra.ExactArgs(1), // Require an account name argument
		Run: func(cmd *cobra.Command, args []string) {
			accountName := args[0]
			fmt.Printf("Editing account: %s...\n", accountName) // Replace with editing logic
		},
	}

	// Add commands
	rootCmd.AddCommand(listCmd, addCmd, editCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
