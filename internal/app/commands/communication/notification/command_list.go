package notification

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

const defaultOffset, defaultLimit uint64 = 0, 10

func (c *CommunicationNotificationCommander) List(inputMessage *tgbotapi.Message) {
	args, err := parseCommandArgs(inputMessage.CommandArguments())
	if err != nil {
		log.Printf("CommunicationNotificationCommander.List wrong args %v", args)
		c.SendErrorMessage(inputMessage.Chat.ID, "Wrong command format. Use /list__communication__notification <offset> <limit>")
		return
	}

	var outputMsgText string

	notifications, err := c.notificationService.List(args.offset, args.limit)
	if err != nil {
		outputMsgText = fmt.Sprintf("No notifications (offset = %d, limit = %d)", args.offset, args.limit)
	} else {
		outputMsgText = fmt.Sprintf("Notifications (offset = %d, limit = %d)\n\n", args.offset, args.limit)
		for _, n := range notifications {
			outputMsgText += n.String() + "\n"
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.List: error sending reply message to chat - %v", err)
	}
}

type ListCommandArgs struct {
	limit  uint64
	offset uint64
}

func parseCommandArgs(inputMessage string) (ListCommandArgs, error) {
	if len(inputMessage) == 0 {
		return ListCommandArgs{limit: defaultLimit, offset: defaultOffset}, nil
	}

	inputParts := strings.SplitN(inputMessage, " ", 2)
	if len(inputParts) != 2 {
		return ListCommandArgs{}, fmt.Errorf("parseCommandArgs: unknown args format")
	} else {
		offset, err := strconv.Atoi(inputParts[0])
		if err != nil || offset < 0 {
			return ListCommandArgs{}, fmt.Errorf("parseCommandArgs: invalid offset %s", inputParts[0])
		}
		limit, err := strconv.Atoi(inputParts[1])
		if err != nil || limit <= 0 {
			return ListCommandArgs{}, fmt.Errorf("parseCommandArgs: invalid limit %s", inputParts[1])
		}
		return ListCommandArgs{limit: uint64(limit), offset: uint64(offset)}, nil
	}
}
