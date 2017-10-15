package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxb-255678085104-xWWKlbAK3V7A8wr7MFQQwMjO")





	params := slack.PostMessageParameters{
		Text:            "Hello World",
		Username:        "",
		AsUser:          false,
		Parse:           "",
		ThreadTimestamp: "",
		LinkNames:       0,
		Attachments:     nil,
		UnfurlLinks:     false,
		UnfurlMedia:     false,
		IconURL:         "",
		IconEmoji:       "",
		Markdown:        false,
		EscapeText:      false,
	}
	conditionField := slack.AttachmentField{
		Title: "Selfies took",
		Value: "0/10",
		Short: true,
	}
	progressionField := slack.AttachmentField{
		Title: "Progression",
		Value: "0%",
		Short: true,
	}

	acceptAction := slack.AttachmentAction{
		Name:            "Accept",
		Text:            "Challenge accepted !",
		Style:           "",
		Type:            "button",
		Value:           "accept",
		DataSource:      "",
		MinQueryLength:  0,
		Options:         nil,
		SelectedOptions: nil,
		OptionGroups:    nil,
		Confirm:         nil,
	}

	refuseAction := slack.AttachmentAction{
		Name:            "Refuse",
		Text:            "No way !",
		Style:           "",
		Type:            "button",
		Value:           "refuse",
		DataSource:      "",
		MinQueryLength:  0,
		Options:         nil,
		SelectedOptions: nil,
		OptionGroups:    nil,
		Confirm:         nil,
	}

	actions := []slack.AttachmentAction{acceptAction, refuseAction}

	fields := []slack.AttachmentField{conditionField, progressionField}
	attachment := slack.Attachment{
		Color:         "439FE0",
		Fallback:      "",
		CallbackID:    "",
		AuthorName:    "Alexandre Mai",
		AuthorSubname: "",
		AuthorLink:    "",
		AuthorIcon:    "",
		Title:         "Take a photo of the founders of FromHexagon !",
		TitleLink:     "https://api.slack.com/",
		Pretext:       "",
		Text:          "Our founders Alex & Rachid are hidden in the StationF, Found them and *take a selfie* with them ! " +
			"If you are lucky, you'll win a coffee ! :heart_eyes:",
		ImageURL:      "https://cdn.smartrecruiters.com/blog/wp-content/uploads/2014/01/tkolind.jpeg",
		ThumbURL:      "",
		Fields:        fields,
		Actions:       actions,
		MarkdownIn:    nil,
		Footer:        "",
		FooterIcon:    "",
		Ts:            "42",
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage("test", "Would you like to play a game ?", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	replyParams := slack.PostMessageParameters{
		Text:            "",
		Username:        "",
		AsUser:          false,
		Parse:           "",
		ThreadTimestamp: timestamp,
		LinkNames:       0,
		Attachments:     nil,
		UnfurlLinks:     false,
		UnfurlMedia:     false,
		IconURL:         "",
		IconEmoji:       "",
		Markdown:        false,
		EscapeText:      false,
	}


	validateAction := slack.AttachmentAction{
		Name:            "validate",
		Text:            "Validé !",
		Style:           "",
		Type:            "button",
		Value:           "validate",
		DataSource:      "",
		MinQueryLength:  0,
		Options:         nil,
		SelectedOptions: nil,
		OptionGroups:    nil,
		Confirm:         nil,
	}

	rejectAction := slack.AttachmentAction{
		Name:            "reject",
		Text:            "Essaye encore !",
		Style:           "",
		Type:            "button",
		Value:           "reject",
		DataSource:      "",
		MinQueryLength:  0,
		Options:         nil,
		SelectedOptions: nil,
		OptionGroups:    nil,
		Confirm:         nil,
	}

	replyActions := []slack.AttachmentAction{validateAction, rejectAction}

	//replyFields := []slack.AttachmentField{conditionField, progressionField}
	replyAttachment := slack.Attachment{
		Color:         "439FE0",
		Fallback:      "",
		CallbackID:    "",
		AuthorName:    "Rachid Berkane",
		AuthorSubname: "",
		AuthorLink:    "",
		AuthorIcon:    "",
		Title:         "",
		TitleLink:     "",
		Pretext:       "",
		Text:          "Je crois que j'ai trouvé !",
		ImageURL:      "http://nextviewventures.com/wp-content/uploads/2016/01/eric-berry.jpg",
		ThumbURL:      "",
		Fields:        nil,
		Actions:       replyActions,
		MarkdownIn:    nil,
		Footer:        "",
		FooterIcon:    "",
		Ts:            "42",
	}
	replyParams.Attachments = []slack.Attachment{replyAttachment}

	replychannelID, replyTimestamp, err := api.PostMessage("test", "Je crois que j'ai trouvé", replyParams)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("reply successfully sent to channel %s at %s", replychannelID, replyTimestamp)
}