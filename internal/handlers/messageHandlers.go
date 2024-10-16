package handlers

import (
	"context"

	"github.com/a-shlyapnikov/simple-httprestapi/internal/messagesService"
	"github.com/a-shlyapnikov/simple-httprestapi/internal/web/messages"
)

type MessageHandler struct {
	Service *messagesService.MessageService
}

func NewMessageHandler(service *messagesService.MessageService) *MessageHandler {
	return &MessageHandler{Service: service}
}

// DeleteMessagesId implements messages.StrictServerInterface.
func (h *MessageHandler) DeleteMessagesId(ctx context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	err := h.Service.DeleteMessage(request.Id)
	if err != nil {
		return messages.DeleteMessagesId404Response{}, err
	}
	return messages.DeleteMessagesId204Response{}, nil
}

// PatchMessagesId implements messages.StrictServerInterface.
func (h *MessageHandler) PatchMessagesId(ctx context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	messageRequest := request.Body
	messageToUpdate := messagesService.Message{Text: *messageRequest.Message}

	updatedMessage, err := h.Service.UpdateMessage(request.Id, messageToUpdate)
	if err != nil {
		return messages.PatchMessagesId404Response{}, err
	}
	response := messages.PatchMessagesId200JSONResponse{
		Id:      &updatedMessage.ID,
		Message: &updatedMessage.Text,
	}
	return response, err
}

// GetMessages implements messages.StrictServerInterface.
func (h *MessageHandler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	return response, nil
}

// PostMessages implements messages.StrictServerInterface.
func (h *MessageHandler) PostMessages(ctx context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}
	return response, nil
}
