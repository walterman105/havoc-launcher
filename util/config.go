package util

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
)

var WorkingDir string

type Config struct {
	Credentials Credentials `json:"credentials"`
	Options     Options     `json:"options"`
}

type Credentials struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RandomPass bool   `json:"random_pass"`
}

type Options struct {
	//Base      []string `json:"base"`
	Console   bool `json:"console"`
	MemoryFix bool `json:"memory_fix"`
}

var config Config

func GetConfig() *Config {
	return &config
}

func SetWorkingDir() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	WorkingDir = dir
}

func UpdateConfigCredentials(cfg *Config, username, password string, randomPass bool) {
	cfg.Credentials.Username = username
	cfg.Credentials.Password = password
	cfg.Credentials.RandomPass = randomPass

	SaveConfig()
}

func UpdateConfigOptions(cfg *Config, console, memoryFix bool) {
	cfg.Options.Console = console
	cfg.Options.MemoryFix = memoryFix
}

func GetCredentials() (string, string) {
	fmt.Println("Username: ", config.Credentials.Username)

	if config.Credentials.RandomPass {
		randomNum, err := rand.Int(rand.Reader, big.NewInt(9000))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Passowrd: ", randomNum)
		return config.Credentials.Username, fmt.Sprintf("%04d", randomNum)
	} else {
		fmt.Println("Passowrd: ", config.Credentials.Password)
		return config.Credentials.Username, config.Credentials.Password
	}
}

func ReadConfig() {
	// Check to see if the config file exists
	_, err := os.Stat("config.json")
	if os.IsNotExist(err) {
		fmt.Println("Config file does not exist!")
	} else {
		// Open the JSON file
		file, err := os.Open("config.json")
		if err != nil {
			fmt.Println("Error opening file:", err)
		}
		fmt.Println("Reading Config")
		defer file.Close()

		// Decode the JSON data into the Config struct
		err = json.NewDecoder(file).Decode(&config)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
		}
	}
}

func SaveConfig() {
	// Convert config to JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling config:", err)
		return
	}

	// Open the file for writing
	file, err := os.OpenFile("config.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Close the file when done

	// Write JSON data to the file
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Config file created successfully.")
}
