package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"getresponse/internal/datastruct"
	"getresponse/internal/dto"
	"getresponse/internal/repository"
	"time"
	"github.com/spf13/viper"

	kafkago "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type GetResponse interface {
	Process(payload *dto.GetResponseV1Payload) (*int64, error)
}

type savePayloadIds struct {
	contacId, accountId, campaignId, sourceCampaignId, messageId, clickTrackId, smsId *int64
}

type getResponse struct {
	dao      repository.DAO
	messages chan kafkago.Message
}

func (g *getResponse) Process(payload *dto.GetResponseV1Payload) (*int64, error) {
	switch payload.Type {
	case "contact_added":
		return g.contactSubscribed(payload)
	case "contact_removed_link":
		return g.contactUnsubscribed(payload)
	case "contact_moved":
		return g.contactMoved(payload)
	case "contact_copied":
		return g.contactCopied(payload)
	case "contact_opened_message":
		return g.contactOpenedMessage(payload)
	case "contact_clicked_message_link":
		return g.contactClickedMessageLink(payload)
	case "contact_clicked_sms_link":
		return g.contactClickedSmsLink(payload)
	default:
		return nil, fmt.Errorf("invalid message type: %s", payload.Type)
	}
}

func (g *getResponse) contactSubscribed(payload *dto.GetResponseV1Payload) (*int64, error) {
	ids, err := g.savePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("contactSubscribed: %w", err)
	}
	contactSubscribed := datastruct.ContactSubscribed{
		OccurredAt: payload.Event.OccurredAt,
		Contact:    *ids.contacId,
		Account:    *ids.accountId,
	}
	err = g.dao.NewContactSubscribedQuery().CreateContactSubscribed(&contactSubscribed)
	if err != nil {
		return nil, fmt.Errorf("contactSubscribed: %w", err)
	}

	err = g.sendToKIS(payload)
	if err != nil {
		errText := fmt.Sprintf("filed send to KIS, type: %s, error: %s", payload.Type, err)
		log.Debug(errText)
		return nil, errors.New(errText)
	}

	return &contactSubscribed.ID, nil
}

func (g *getResponse) contactUnsubscribed(payload *dto.GetResponseV1Payload) (*int64, error) {
	ids, err := g.savePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("contactUnsubscribed: %w", err)
	}
	contactUnsubscribed := datastruct.ContactUnsubscribed{
		OccurredAt: payload.Event.OccurredAt,
		Contact:    *ids.contacId,
		Account:    *ids.accountId,
	}
	err = g.dao.NewContactUnsubscribedQuery().CreateContactUnsubscribed(&contactUnsubscribed)
	if err != nil {
		return nil, fmt.Errorf("contactUnsubscribed: %w", err)
	}

	err = g.sendToKIS(payload)
	if err != nil {
		errText := fmt.Sprintf("filed send to KIS, type: %s, error: %s", payload.Type, err)
		log.Debug(errText)
		return nil, errors.New(errText)
	}

	return &contactUnsubscribed.ID, nil
}

func (g *getResponse) contactMoved(payload *dto.GetResponseV1Payload) (*int64, error) {
	ids, err := g.savePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("contactMoved: %w", err)
	}
	contactMoved := datastruct.ContactMoved{
		OccurredAt:     payload.Event.OccurredAt,
		Contact:        *ids.contacId,
		Account:        *ids.accountId,
		SourceCampaign: *ids.sourceCampaignId,
		CampaignTarget: *ids.campaignId,
	}
	err = g.dao.NewContactMovedQuery().CreateContactMoved(&contactMoved)
	if err != nil {
		return nil, fmt.Errorf("contactMoved: %w", err)
	}

	err = g.sendToKIS(payload)
	if err != nil {
		errText := fmt.Sprintf("filed send to KIS, type: %s, error: %s", payload.Type, err)
		log.Debug(errText)
		return nil, errors.New(errText)
	}

	return &contactMoved.ID, nil
}

func (g *getResponse) contactCopied(payload *dto.GetResponseV1Payload) (*int64, error) {
	ids, err := g.savePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("contactCopied: %w", err)
	}
	contactCopied := datastruct.ContactCopied{
		OccurredAt:     payload.Event.OccurredAt,
		Contact:        *ids.contacId,
		Account:        *ids.accountId,
		SourceCampaign: *ids.sourceCampaignId,
		CampaignTarget: *ids.campaignId,
	}
	err = g.dao.NewContactCopiedQuery().CreateContactCopied(&contactCopied)
	if err != nil {
		return nil, fmt.Errorf("contactCopied: %w", err)
	}

	err = g.sendToKIS(payload)
	if err != nil {
		errText := fmt.Sprintf("filed send to KIS, type: %s, error: %s", payload.Type, err)
		log.Debug(errText)
		return nil, errors.New(errText)
	}

	return &contactCopied.ID, nil
}

func (g *getResponse) contactOpenedMessage(payload *dto.GetResponseV1Payload) (*int64, error) {
	ids, err := g.savePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("contactOpenedMessage: %w", err)
	}
	contactOpenedMessage := datastruct.ContactOpenedMessage{
		OccurredAt: payload.Event.OccurredAt,
		Contact:    *ids.contacId,
		Account:    *ids.accountId,
		Message:    *ids.messageId,
	}
	err = g.dao.NewContactOpenedMessageQuery().CreateContactOpenedMessage(&contactOpenedMessage)
	if err != nil {
		return nil, fmt.Errorf("contactOpenedMessage: %w", err)
	}

	err = g.sendToKIS(payload)
	if err != nil {
		errText := fmt.Sprintf("filed send to KIS, type: %s, error: %s", payload.Type, err)
		log.Debug(errText)
		return nil, errors.New(errText)
	}

	return &contactOpenedMessage.ID, nil
}

func (g *getResponse) contactClickedMessageLink(payload *dto.GetResponseV1Payload) (*int64, error) {
	ids, err := g.savePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("contactClickedMessageLink: %w", err)
	}
	contactLinkClicked := datastruct.ContactLinkClicked{
		OccurredAt: payload.Event.OccurredAt,
		Contact:    *ids.contacId,
		Account:    *ids.accountId,
		Message:    *ids.messageId,
		ClickTrack: *ids.clickTrackId,
	}
	err = g.dao.NewContactLinkClickedQuery().CreateContactLinkClicked(&contactLinkClicked)
	if err != nil {
		return nil, fmt.Errorf("contactClickedMessageLink: %w", err)
	}

	err = g.sendToKIS(payload)
	if err != nil {
		errText := fmt.Sprintf("filed send to KIS, type: %s, error: %s", payload.Type, err)
		log.Debug(errText)
		return nil, errors.New(errText)
	}

	return &contactLinkClicked.ID, nil
}

func (g *getResponse) contactClickedSmsLink(payload *dto.GetResponseV1Payload) (*int64, error) {
	ids, err := g.savePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("contactClickedSmsLink: %w", err)
	}
	contactSmsLinkClicked := datastruct.ContactSmsLinkClicked{
		OccurredAt: payload.Event.OccurredAt,
		Contact:    *ids.contacId,
		Account:    *ids.accountId,
		SMS:        *ids.smsId,
		ClickTrack: *ids.clickTrackId,
	}
	err = g.dao.NewContactSmsLinkClickedQuery().CreateContactSmsLinkClicked(&contactSmsLinkClicked)
	if err != nil {
		return nil, fmt.Errorf("contactClickedSmsLink: %w", err)
	}

	err = g.sendToKIS(payload)
	if err != nil {
		errText := fmt.Sprintf("filed send to KIS, type: %s, error: %s", payload.Type, err)
		log.Debug(errText)
		return nil, errors.New(errText)
	}

	return &contactSmsLinkClicked.ID, nil
}

func (g *getResponse) sendToKIS(payload *dto.GetResponseV1Payload) error {
	value, _ := json.Marshal(payload)
	key := viper.Get("kafka.producer.key").(string)
	log.WithFields(log.Fields{"payload": payload}).Debug("sendding message to kafka")
	message := kafkago.Message{
		Value: value,
		Key:   []byte(fmt.Sprintf("%v", key)),
	}
	g.messages <- message

	return nil
}

func (g *getResponse) savePayload(payload *dto.GetResponseV1Payload) (*savePayloadIds, error) {
	account, err := g.saveAccount(&payload.Account)
	if err != nil {
		return nil, fmt.Errorf("savePayload: %w", err)
	}
	campaign, err := g.saveCampaign(&payload.Contact.Campaign)
	if err != nil {
		return nil, fmt.Errorf("savePayload: %w", err)
	}
	sourceCampaign, err := g.saveSourceCampaign(&payload.Contact.SourceCampaign)
	if err != nil {
		return nil, fmt.Errorf("savePayload: %w", err)
	}
	customField, err := g.saveCustomField(&payload.Contact.PhoneNumber)
	if err != nil {
		return nil, fmt.Errorf("savePayload: %w", err)
	}
	contact, err := g.saveContact(&payload.Contact, campaign, customField)
	if err != nil {
		return nil, fmt.Errorf("savePayload: %w", err)
	}
	message, err := g.saveMessage(&payload.Message)
	if err != nil {
		return nil, fmt.Errorf("savePayload: %w", err)
	}
	clickTrack, err := g.saveClickTrack(&payload.ClickTrack)
	if err != nil {
		return nil, fmt.Errorf("savePayload: %w", err)
	}
	sms, err := g.saveSMS(&payload.SMS)
	if err != nil {
		return nil, fmt.Errorf("savePayload: %w", err)
	}

	ids := savePayloadIds{
		contacId:         contact,
		accountId:        account,
		campaignId:       campaign,
		sourceCampaignId: sourceCampaign,
		messageId:        message,
		clickTrackId:     clickTrack,
		smsId:            sms,
	}
	return &ids, nil
}

func (g *getResponse) saveAccount(accountIn *dto.Account) (*int64, error) {
	account := datastruct.Account{
		AccountId: accountIn.AccountId,
	}
	err := g.dao.NewAccountQuery().CreateAccount(&account)
	if err != nil {
		return nil, fmt.Errorf("saveAccount: %w", err)
	}
	return &account.ID, nil
}

func (g *getResponse) saveCampaign(campaignIn *dto.Campaign) (*int64, error) {
	campaign := datastruct.Campaign{
		CampaignId: campaignIn.CampaignId,
		Name:       campaignIn.Name,
		Href:       campaignIn.Href,
	}
	err := g.dao.NewCampaignQuery().CreateCampaign(&campaign)
	if err != nil {
		return nil, fmt.Errorf("saveCampaign: %w", err)
	}

	return &campaign.ID, nil
}

func (g *getResponse) saveSourceCampaign(sourceCampaignIn *dto.SourceCampaign) (*int64, error) {
	sourceCampaign := datastruct.Campaign{
		CampaignId: sourceCampaignIn.CampaignId,
		Name:       sourceCampaignIn.CampaignName,
		Href:       sourceCampaignIn.Href,
	}
	if *sourceCampaignIn != (dto.SourceCampaign{}) {
		err := g.dao.NewCampaignQuery().CreateCampaign(&sourceCampaign)
		if err != nil {
			return nil, fmt.Errorf("saveSourceCampaign: %w", err)
		}
	}
	return &sourceCampaign.ID, nil
}

func (g *getResponse) saveCustomField(customFieldIn *dto.CustomField) (*int64, error) {
	customField := datastruct.CustomField{
		FieldId: customFieldIn.FieldId,
		Href:    customFieldIn.Href,
	}
	if *customFieldIn != (dto.CustomField{}) {
		err := g.dao.NewCustomFieldQuery().CreateCustomField(&customField)
		if err != nil {
			return nil, fmt.Errorf("saveCustomField: %w", err)
		}
	}

	return &customField.ID, nil
}

func (g *getResponse) saveContact(
	contactIn *dto.Contact, campaignID *int64, customFieldID *int64) (*int64, error) {
	contact := datastruct.Contact{
		ContactId:   contactIn.ContactId,
		Email:       contactIn.Email,
		Name:        contactIn.Name,
		Ip:          contactIn.Ip,
		Origin:      contactIn.Origin,
		Href:        contactIn.Href,
		Campaign:    *campaignID,
		PhoneNumber: *customFieldID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := g.dao.NewContactQuery().CreateContact(&contact)
	if err != nil {
		return nil, fmt.Errorf("saveContact: %w", err)
	}

	return &contact.ID, nil
}

func (g *getResponse) saveMessage(messageIn *dto.Message) (*int64, error) {
	message := datastruct.Message{
		ResourceId:   messageIn.ResourceId,
		ResourceType: messageIn.ResourceType,
		Name:         messageIn.Name,
		Href:         messageIn.Href,
		Subject:      messageIn.Subject,
	}
	if *messageIn != (dto.Message{}) {
		err := g.dao.NewMessageQuery().CreateMessage(&message)
		if err != nil {
			return nil, fmt.Errorf("saveMessage: %w", err)
		}
	}

	return &message.ID, nil
}

func (g *getResponse) saveClickTrack(clickTrackIn *dto.ClickTrack) (*int64, error) {
	clickTrack := datastruct.ClickTrack{
		ClickTrackId: clickTrackIn.ClickTrackId,
		Name:         clickTrackIn.Name,
		Href:         clickTrackIn.Href,
		Url:          clickTrackIn.Url,
	}
	if *clickTrackIn != (dto.ClickTrack{}) {
		err := g.dao.NewClickTrackQuery().CreateClickTrack(&clickTrack)
		if err != nil {
			return nil, fmt.Errorf("saveClickTrack: %w", err)
		}
	}

	return &clickTrack.ID, nil
}

func (g *getResponse) saveSMS(smsIn *dto.SMS) (*int64, error) {
	sms := datastruct.SMS{
		SmsId: smsIn.SmsId,
		Name:  smsIn.Name,
		Href:  smsIn.Href,
	}
	if *smsIn != (dto.SMS{}) {
		err := g.dao.NewSMSQuery().CreateSMS(&sms)
		if err != nil {
			return nil, fmt.Errorf("saveSMS: %w", err)
		}
	}

	return &sms.ID, nil
}
