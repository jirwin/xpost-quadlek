package xpost

import (
	"context"
	"fmt"
	"strings"

	"github.com/jirwin/quadlek/quadlek"
)

const XpostPrefix = "xpost-"

func xpostReaction(ctx context.Context, reactionChannel <-chan *quadlek.ReactionHookMsg) {
	for {
		select {
		case rh := <-reactionChannel:
			if strings.HasPrefix(rh.Reaction.Reaction, XpostPrefix) {
				dstChan := strings.TrimPrefix(rh.Reaction.Reaction, XpostPrefix)
				if dstChan == rh.Reaction.Item.Channel {
					continue
				}
				dstChanId, err := rh.Bot.GetChannelId(dstChan)
				if err != nil {
					fmt.Println("error getting channel id", err.Error())
					continue
				}
				rh.Bot.Say(dstChanId, "xpost inc!")
			}

		case <-ctx.Done():
			fmt.Println("Shutting down xpost")
			return
		}
	}
}

func Register() quadlek.Plugin {
	return quadlek.MakePlugin(
		"xpost",
		nil,
		nil,
		[]quadlek.ReactionHook{
			quadlek.MakeReactionHook(xpostReaction),
		},
		nil,
		nil,
	)
}
