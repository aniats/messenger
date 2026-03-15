package repackers

import (
	"github.com/aniats/messenger/internal/model"
	"github.com/aniats/messenger/internal/service"
	pb "github.com/aniats/messenger/pkg/messenger/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func chatToProto(chat model.Chat) *pb.ChatInfo {
	return &pb.ChatInfo{
		ChatId:     chat.Id.String(),
		Name:       chat.Name,
		IsReadOnly: chat.IsReadOnly,
	}
}

func ChatsToProto(chats []model.Chat) []*pb.ChatInfo {
	result := make([]*pb.ChatInfo, 0, len(chats))
	for _, chat := range chats {
		result = append(result, chatToProto(chat))
	}
	return result
}

func CreateChatRequestToParams(sessionID uuid.UUID, req *pb.CreateChatRequest) service.CreateChatParams {
	params := service.CreateChatParams{
		SessionID:  sessionID,
		Name:       req.GetName(),
		IsReadOnly: req.GetIsReadOnly(),
	}

	if req.ChatTtl != nil {
		params.ChatTtlMs = req.GetChatTtl().AsDuration().Milliseconds()
	}

	if req.MessageTtl != nil {
		params.MessageTtlMs = req.GetMessageTtl().AsDuration().Milliseconds()
	}

	if req.MaxMessages != nil {
		params.MaxMessages = int64(req.GetMaxMessages())
	}

	return params
}

func ListChatsRequestToParams(sessionID uuid.UUID, req *pb.ListChatsRequest) service.ListChatsParams {
	return service.ListChatsParams{
		SessionID: sessionID,
		OnlyMy:    req.GetFilter() == pb.ChatFilter_CHAT_FILTER_MY,
	}
}

func ParseSessionID(raw string) (uuid.UUID, error) {
	if raw == "" {
		return uuid.Nil, status.Error(codes.InvalidArgument, "session_id is required")
	}

	id, err := uuid.Parse(raw)
	if err != nil {
		return uuid.Nil, status.Errorf(codes.InvalidArgument, "invalid session_id: %v", err)
	}

	return id, nil
}
