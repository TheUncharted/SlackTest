package main

import (
	"github.com/nlopes/slack"
	"fmt"
	"strings"
	"strconv"
	"math"
)

func main() {
	api := slack.New("xoxb-255678085104-xWWKlbAK3V7A8wr7MFQQwMjO")

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Println("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace #general with your Channel ID
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "#general"))

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
		//
		//case *slack.PresenceChangeEvent:
		//	fmt.Printf("Presence Change: %v\n", ev)
		//
		//case *slack.LatencyReport:
		//	fmt.Printf("Current latency: %v\n", ev.Value)
		//
		//case *slack.RTMError:
		//	fmt.Printf("Error: %s\n", ev.Error())
		//
		//case *slack.InvalidAuthEvent:
		//	fmt.Printf("Invalid credentials")
		//	return
		case *slack.ReactionAddedEvent:
			fmt.Printf("Reaction added: %+v\n", ev)
			messageWithEventTimestamp := ev.Item.Timestamp
			channel := ev.Item.Channel
			channelHistory := slack.HistoryParameters{
				Latest:    messageWithEventTimestamp,
				Oldest:    "",
				Count:     1,
				Inclusive: true,
				Unreads:   false,
			}
			history, err:= api.GetChannelHistory(channel, channelHistory)
			if err == nil &&  len(history.Messages) > 0{
				message := history.Messages[0]
				threadTimestamp := message.ThreadTimestamp

				channelHistory = slack.HistoryParameters{
					Latest:    threadTimestamp,
					Oldest:    "",
					Count:     1,
					Inclusive: true,
					Unreads:   false,
				}
				history, err:= api.GetChannelHistory(channel, channelHistory)
				if err == nil && len(history.Messages) > 0 {
					threadMessage := history.Messages[0]
					fmt.Printf("thread mesage : %+v\n", threadMessage)
					attachments := threadMessage.Msg.Attachments
					fields := attachments[0].Fields
					split := strings.Split(fields[0].Value, "/")
					currentValidated, _ := strconv.Atoi(split[0])
					currentValidated++
					fields[0].Value = strconv.Itoa(currentValidated) + "/10"

					currentPercent := (currentValidated *100) / 10


					var bar string
					nbcursor := currentPercent / 10
					leftcursor := 10 - nbcursor
					for i := 0; i < nbcursor; i++ {
						bar += ":black_medium_square:"
					}
					for i := 0; i < leftcursor; i++ {
						bar += ":white_medium_square:"
					}
					fields[1].Value = bar + strconv.Itoa(currentPercent) + "%"
					//endi

					params := slack.PostMessageParameters{
						Text:            "Hello World",
						Username:        "",
						AsUser:          false,
						Parse:           "",
						ThreadTimestamp: threadTimestamp,
						LinkNames:       0,
						Attachments:     nil,
						UnfurlLinks:     false,
						UnfurlMedia:     false,
						IconURL:         "",
						IconEmoji:       "",
						Markdown:        false,
						EscapeText:      false,

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

					//fields := []slack.AttachmentField{conditionField, progressionField}
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
					//channelID, timestamp, err := api.PostMessage("test", "Would you like to play a game ?", params)
					//api.UpdateMessage()

					channelID, timestamp, text, err := api.SendMessage(channel, slack.MsgOptionUpdate(threadTimestamp), slack.MsgOptionAttachments(attachment))

					if err != nil {
						fmt.Printf("%s\n", err)
						return
					}
					fmt.Printf("Message updated successfully sent to channel %s at %s, %s", channelID, timestamp, text)
				}


			}


			//ev.
			//api.UpdateMessage("test", )
		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}



}

func Round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
