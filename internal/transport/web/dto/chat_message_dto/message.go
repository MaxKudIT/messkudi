package chat_message_dto

import (
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/google/uuid"
	"time"
)

type ChatMessageDTOClient struct {
	Id                 string
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //chat
	ChatId             string                   //айди чата
	SenderId           string
	RecieverId         string
}

type ChatMessageDTOClientParsing struct {
	Id                 uuid.UUID
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //chat
	ChatId             uuid.UUID                //айди чата
	SenderId           uuid.UUID
	RecieverId         uuid.UUID
}

func (cmdto *ChatMessageDTOClient) UuidParse() (ChatMessageDTOClientParsing, error) {
	uuidI, err := uuid.Parse(cmdto.Id)
	if err != nil {
		return ChatMessageDTOClientParsing{}, err
	}

	uuidR, err := uuid.Parse(cmdto.RecieverId)
	if err != nil {
		return ChatMessageDTOClientParsing{}, err
	}
	uuidS, err := uuid.Parse(cmdto.SenderId)
	if err != nil {
		return ChatMessageDTOClientParsing{}, err
	}
	uuidC, err := uuid.Parse(cmdto.ChatId)
	if err != nil {
		return ChatMessageDTOClientParsing{}, err
	}
	return ChatMessageDTOClientParsing{Id: uuidI, Type: cmdto.Type, Content: cmdto.Content, CorrespondenceType: cmdto.CorrespondenceType, ChatId: uuidC, SenderId: uuidS, RecieverId: uuidR}, nil
}

type ChatMessageDTODetailsServer struct {
	Id                 uuid.UUID
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //chat
	ChatId             uuid.UUID                //айди чата
	CreatedAt          time.Time
	UpdatedAt          time.Time
	ReadAt             *time.Time
	SenderId           uuid.UUID
}

func ToDomain(createdat time.Time, updatedat time.Time, readAt *time.Time, cmDTO ChatMessageDTOClientParsing) messages.ChatMessage {
	return messages.ChatMessage{
		Id:                 cmDTO.Id,
		SenderId:           cmDTO.SenderId,
		CreatedAt:          createdat,
		UpdatedAt:          updatedat,
		ChatId:             cmDTO.ChatId,
		CorrespondenceType: cmDTO.CorrespondenceType,
		Content:            cmDTO.Content,
		Type:               cmDTO.Type,
		ReadAt:             readAt,
	}
}
