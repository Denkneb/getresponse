package handlers

import (
	"encoding/json"
	"getresponse/internal/dto"
	"getresponse/internal/repository"
	"getresponse/internal/rest"
	kafkago "github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

type webhookGetResponse struct {
	dao      repository.DAO
	messages chan kafkago.Message
}

type WebhookGetResponse interface {
	Webhook() http.HandlerFunc
}

func NewWebhook(dao repository.DAO, messages chan kafkago.Message) WebhookGetResponse {
	return &webhookGetResponse{dao: dao, messages: messages}
}

func (w *webhookGetResponse) NewGetResponse() GetResponse {
	return &getResponse{dao: w.dao, messages: w.messages}
}

func (webhook *webhookGetResponse) Webhook() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var getresponsePayload dto.GetResponseV1Payload
		err := json.NewDecoder(r.Body).Decode(&getresponsePayload)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}
		log.Printf("getting payload -> %+v\n", getresponsePayload)
		id, err := webhook.NewGetResponse().Process(&getresponsePayload)
		if err != nil {
			rest.WriteError(w, http.StatusBadRequest, err)
			return
		}
		rest.WriteJSON(w, http.StatusCreated, rest.Response{
			Ok:     true,
			Result: id,
		})
	})
}
