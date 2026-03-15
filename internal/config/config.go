// Package config provides application configuration parsing from flags and environment variables.
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

const (
	maxChatSizeDefault = 100
	maxChatsNumDefault = 10
	portDefault        = 50051
)

// Config holds the application configuration.
type Config struct {
	Port     int
	ChatSize int
	ChatsNum int
	Env      string
	LogLevel string
}

func lookupInt(envKey string, defaultVal int) int {
	if v := os.Getenv(envKey); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("invalid %s value: %s", envKey, v)
		}

		return val
	}

	return defaultVal
}

func lookupString(envKey string, defaultVal string) string {
	if v := os.Getenv(envKey); v != "" {
		return v
	}

	return defaultVal
}

// Parse parses configuration from flags and environment variables.
func Parse() *Config {
	chatSize := pflag.Int("max-chat-size", 0, "max number of messages in a chat")
	chatsNum := pflag.Int("max-chats-num", 0, "max number of chats")
	port := pflag.Int("port", 0, "gRPC server port")

	pflag.Parse()

	cfg := &Config{
		Port:     0,
		ChatSize: 0,
		ChatsNum: 0,
		Env:      lookupString("ENV", "local"),
		LogLevel: lookupString("LOG_LEVEL", "debug"),
	}

	if pflag.Lookup("max-chat-size").Changed {
		cfg.ChatSize = *chatSize
	} else {
		cfg.ChatSize = lookupInt("MAX_CHAT_SIZE", maxChatSizeDefault)
	}

	if pflag.Lookup("max-chats-num").Changed {
		cfg.ChatsNum = *chatsNum
	} else {
		cfg.ChatsNum = lookupInt("MAX_CHATS_NUM", maxChatsNumDefault)
	}

	if pflag.Lookup("port").Changed {
		cfg.Port = *port
	} else {
		cfg.Port = lookupInt("PORT", portDefault)
	}

	return cfg
}
