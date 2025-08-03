package dto

import (
	"github.com/MaxKudIT/messkudi/internal/domain/chats"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto/chat_message_dto"
	"github.com/google/uuid"
	"time"
)

type ChatDTOClient struct {
	Participant string
}

type ChatDTOClientParsing struct {
	Participant uuid.UUID
}

func (cdc *ChatDTOClient) Parse() (ChatDTOClientParsing, error) {
	uuid, err := uuid.Parse(cdc.Participant)
	if err != nil {
		return ChatDTOClientParsing{}, err
	}
	return ChatDTOClientParsing{Participant: uuid}, nil
}

type ChatDTODetailsServer struct {
	CreatorId   uuid.UUID
	Participant uuid.UUID
	ChatCount   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ChatPreviewDTOServer struct {
	User struct {
		Name   string
		Color  string
		Status bool
	}
	MessageMeta struct {
		Content        string
		IsRead         bool
		CreatedAt      time.Time
		IsMy           bool
		SenderId       uuid.UUID
		UnReadMessages []uuid.UUID
	}
	ChatId        uuid.UUID
	ParticipantId uuid.UUID
}
type ChatHeader struct {
	Id     uuid.UUID
	Name   string
	Color  string
	Status bool
	//LastSeen time.Time
}

type ChatDataDTOServer struct {
	Header   ChatHeader
	Messages []chat_message_dto.ChatMessageDTODetailsServer
}

func ToDomainChat(chatid uuid.UUID, creatorid uuid.UUID, createdat time.Time, updatedAt time.Time, dto ChatDTOClientParsing) chats.Chat {
	return chats.Chat{
		Id:          chatid,
		CreatorId:   creatorid,
		Participant: dto.Participant,
		CreatedAt:   createdat,
		UpdatedAt:   updatedAt,
	}
}
