package slack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"

	"github.com/danielqsj/k8s-slackbot/pkg/kubernetes"
)

// SlackBot defines a slack client
type SlackBot struct {
	Client *slack.Client
}

// InitSlackLog inits slack log
func InitSlackLog() {
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
}

// NewSlackBot returns new slackbot with token
func NewSlackBot(token string) SlackBot {
	return SlackBot{
		Client: slack.New(token),
	}
}

// EnableDebug enables debug log
func (bot SlackBot) EnableDebug() {
	bot.Client.SetDebug(true)
}

// RunSlackRTMServer runs rtm server
func (bot SlackBot) RunSlackRTMServer(kubeconfig string) {
	rtm := bot.Client.NewRTM()
	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			input := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")
			if len(input) > 0 {
				command := input[0]
				args := input[1:]
				log.Printf("command: %v, args: %v", command, args)
				switch command {
				case "kubectl":
					if len(args) > 0 {
						output, err := kubernetes.RunKubectl(kubeconfig, args)
						if err != nil {
							rtm.SendMessage(rtm.NewOutgoingMessage(err.Error(), ev.Msg.Channel))
						} else {
							rtm.SendMessage(rtm.NewOutgoingMessage(output, ev.Msg.Channel))
						}
					}
				}
			}

		case *slack.InvalidAuthEvent:
			log.Printf("Invalid credentials")
			return

		case *slack.RTMError:
			log.Printf("Error: %s\n", ev.Error())

		default:
		}
	}
}

// GetUserName returns username by userid
func (bot SlackBot) GetUserName(userId string) string {
	user, err := bot.Client.GetUserInfo(userId)
	if err != nil {
		bot.Client.Debugf("GetUserName err: %s\n", err)
		return ""
	}
	return user.Profile.FirstName + " " + user.Profile.LastName
}

// GetUserId returns userid by username
func (bot SlackBot) GetUserId(userName string) (string, error) {
	users, err := bot.Client.GetUsers()
	if err != nil {
		return "", err
	}
	for _, user := range users {
		if user.Name == userName {
			return user.ID, nil
		}
	}
	return "", fmt.Errorf("Cannot find this user %s", userName)
}

// SendMessage sends message to users
func (bot SlackBot) SendMessage(userNames []string, message string) {
	params := slack.PostMessageParameters{}
	for _, userName := range userNames {
		userId, err := bot.GetUserId(userName)
		if err != nil {
			log.Printf("GetUserId err: %s\n", err)
			return
		}
		channelID, timestamp, err := bot.Client.PostMessage(userId, message, params)
		if err != nil {
			log.Printf("Send slack message err: %s\n", err)
			return
		}
		log.Printf("Send slack message successfully to channel %s at %s", channelID, timestamp)
	}
}

// SendMessages sends messages to a user
func (bot SlackBot) SendMessages(userName string, message []string) {
	params := slack.PostMessageParameters{}
	userId, err := bot.GetUserId(userName)
	if err != nil {
		log.Printf("GetUserId err: %s\n", err)
		return
	}
	channelID, timestamp, err := bot.Client.PostMessage(userId, strings.Join(message, "\n"), params)
	if err != nil {
		log.Printf("Send slack message err: %s\n", err)
		return
	}
	log.Printf("Send slack message successfully to channel %s at %s", channelID, timestamp)
}
