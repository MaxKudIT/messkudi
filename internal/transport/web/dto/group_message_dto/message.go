package group_message_dto

import (
	"github.com/MaxKudIT/messkudi/internal/domain/messages"
	"github.com/google/uuid"
	"time"
)

type GroupMessageDTOClient struct {
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //group
	GroupId            uuid.UUID                //айди группы
}

type GroupMessageDTODetailsServer struct {
	Type               messages.MessageT
	Content            string
	CorrespondenceType messages.CorrespondenceT //group
	GroupId            uuid.UUID                //айди группы
	CreatedAt          time.Time
	UpdatedAt          *time.Time
	ReadAt             *time.Time
	SenderId           uuid.UUID
}

func ToDomain(messageId uuid.UUID, senderid uuid.UUID, createdat time.Time, updatedat time.Time, gmDTO GroupMessageDTOClient) messages.GroupMessage {
	return messages.GroupMessage{
		Id:                 messageId,
		SenderId:           senderid,
		CreatedAt:          createdat,
		UpdatedAt:          updatedat,
		GroupId:            gmDTO.GroupId,
		CorrespondenceType: gmDTO.CorrespondenceType,
		Content:            gmDTO.Content,
		Type:               gmDTO.Type,
	}
}
