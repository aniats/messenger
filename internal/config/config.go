// Package config provides application configuration parsing from flags and environment variables.
package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

const (
	maxChatSizeDefault          = 100
	maxChatsNumDefault          = 10
	portDefault                 = 50051
	sessionIdleTimeoutDefault   = 3600 // seconds
)

// Config holds the application configuration.
type Config struct {
	Port               int
	ChatSize           int
	ChatsNum           int
	Env                string
	LogLevel           string
	SessionIdleTimeout time.Duration
}

func lookupInt(envKey string, defaultVal int) (int, error) {
	if v := os.Getenv(envKey); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("invalid %s value: %q: %w", envKey, v, err)
		}

		return val, nil
	}

	return defaultVal, nil
}

func lookupString(envKey string, defaultVal string) string {
	if v := os.Getenv(envKey); v != "" {
		return v
	}

	return defaultVal
}

func Parse() (*Config, error) {
	chatSize := pflag.Int("max-chat-size", 0, "max number of messages in a chat")
	chatsNum := pflag.Int("max-chats-num", 0, "max number of chats")
	port := pflag.Int("port", 0, "gRPC server port")

	pflag.Parse()

	cfg := &Config{
		Env:      lookupString("ENV", "local"),
		LogLevel: lookupString("LOG_LEVEL", "debug"),
	}

	if pflag.Lookup("max-chat-size").Changed {
		cfg.ChatSize = *chatSize
	} else {
		val, err := lookupInt("MAX_CHAT_SIZE", maxChatSizeDefault)
		if err != nil {
			return nil, err
		}

		cfg.ChatSize = val
	}

	if pflag.Lookup("max-chats-num").Changed {
		cfg.ChatsNum = *chatsNum
	} else {
		val, err := lookupInt("MAX_CHATS_NUM", maxChatsNumDefault)
		if err != nil {
			return nil, err
		}

		cfg.ChatsNum = val
	}

	if pflag.Lookup("port").Changed {
		cfg.Port = *port
	} else {
		val, err := lookupInt("PORT", portDefault)
		if err != nil {
			return nil, err
		}

		cfg.Port = val
	}

	sessionTimeoutSec, err := lookupInt("SESSION_IDLE_TIMEOUT", sessionIdleTimeoutDefault)
	if err != nil {
		return nil, err
	}

	cfg.SessionIdleTimeout = time.Duration(sessionTimeoutSec) * time.Second

	return cfg, nil
}
