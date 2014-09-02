package sms

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"net/smtp"
	"strings"
)

type SMSClient struct {
	Address    string
	Port       int
	Username   string
	Password   string
	config     *Config
	smtpClient *smtp.Client
}

type Config struct {
	FromAddress string            `yaml:"from_address"`
	Carriers    map[string]string `yaml:"carriers"`
}

func createClient(address string, port int, username string, password string) (SMSClient, error) {
	client := SMSClient{Address: address, Port: port, Username: username, Password: password}
	config := Config{Carriers: make(map[string]string)}
	buffer, err := ioutil.ReadFile("sms.yml")

	if err != nil {
		return SMSClient{}, err
	}

	if err := yaml.Unmarshal(buffer, &config); err != nil {
		return SMSClient{}, err
	}

	client.config = &config

	return client, nil
}

func (client *SMSClient) Deliver(number, carrier, message string) error {
	c := client.config.Carriers[strings.ToLower(carrier)]

	if c == "" {
		return errors.New("Unsupported carrier. Please check sms.yml for supported carriers.")
	}

	to := []string{fmt.Sprintf("%s%s", number, c)}
	auth := smtp.PlainAuth("", client.Username, client.Password, client.Address)
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", client.Address, client.Port), auth, client.config.FromAddress,
		to, []byte(message)); err != nil {
		return err
	}

	return nil
}
