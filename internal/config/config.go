package config

import (
	"flag"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port     int
	ChatSize int
	ChatsNum int
	Env      string
}

func Parse() *Config {
	maxChatSizeDefault := os.Getenv("MAX_CHAT_SIZE")
	maxChatsNumDefault := os.Getenv("MAX_CHATS_NUM")
	portDefault := os.Getenv("PORT")
	env := os.Getenv("ENV")

	chatSize := flag.Int("chat-size", 100, "max number of messages in a chat")
	chatsNum := flag.Int("chats-num", 10, "max number of chats")
	port := flag.Int("port", 50051, "gRPC server port")

	flag.Parse()

	var isSetChatSize, isSetChatsNum, isSetPort bool

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "chat-size":
			isSetChatSize = true
		case "chats-num":
			isSetChatsNum = true
		case "port":
			isSetPort = true
		}
	})

	if !isSetChatSize && maxChatSizeDefault != "" {
		val, err := strconv.Atoi(maxChatSizeDefault)
		if err != nil {
			log.Fatalf("invalid MAX_CHAT_SIZE value: %s", maxChatSizeDefault)
		}
		*chatSize = val
	}

	if !isSetChatsNum && maxChatsNumDefault != "" {
		val, err := strconv.Atoi(maxChatsNumDefault)
		if err != nil {
			log.Fatalf("invalid MAX_CHATS_NUM value: %s", maxChatsNumDefault)
		}
		*chatsNum = val
	}

	if !isSetPort && portDefault != "" {
		val, err := strconv.Atoi(portDefault)
		if err != nil {
			log.Fatalf("invalid PORT value: %s", portDefault)
		}
		*port = val
	}

	return &Config{
		Port:     *port,
		ChatSize: *chatSize,
		ChatsNum: *chatsNum,
		Env:      env,
	}
}
