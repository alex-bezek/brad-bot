package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

var (
	// conf         = new(config)
	dg         *discordgo.Session
	lastReboot string
	// log          = newLog()
	// status       = map[discordgo.Status]string{"dnd": "busy", "online": "online", "idle": "idle", "offline": "offline"}
	footer = new(discordgo.MessageEmbedFooter)
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	footer.Text = "Last Bot reboot: " + time.Now().Format("Mon, 02-Jan-06 15:04:05 MST")
}

func main() {

	// runtime.GOMAXPROCS(conf.MaxProc)

	// log.Info("/*********BOT RESTARTING*********\\")

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)
	// MessageReactionAdd
	// MessageReactionRemove

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// TODO: Add prefix check

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Contains(m.Content, "pixie") && strings.Contains(m.Content, "stick") {
		s.ChannelMessageSend(m.ChannelID, "https://giphy.com/gifs/range-of-emotions-cotton-candy-girl-safeco-field-foodie-3o6Ztg2MgUkcXyCgtG")
	}

	if m.Content == "battle" {
		// squad := [5]string{"Furus", "Galia", "Adalia", "Carindrina", "Belladona"}
		message, err := s.ChannelMessageSendEmbed(m.ChannelID, embed())
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "🤠")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "🦊")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "🐈‍⬛")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "🐉")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "🏹")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "🌊")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "⛏️")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "👼")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "😻")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "🐆")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "😇")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "🦡")
		if err != nil {
			fmt.Println(err)
		}
		err = s.MessageReactionAdd(m.ChannelID, message.ID, "😺")
		if err != nil {
			fmt.Println(err)
		}
		// Clear current state for new battle
		// for _, p := range squad {
		// 	s.ChannelMessageSend(m.ChannelID, p)
		// 	// Add each number emoji to player
		// 	// "⏺1️⃣1️⃣2️⃣2️⃣3️⃣3️⃣4️⃣4️⃣"
		// }
		// func (s *Session) MessageReactionAdd(channelID, messageID, emojiID string) error {

		// s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

// :cowboy: Furus
// :fox: Galia
// :black_cat:  venus
// :dragon_face:  midna
// :bow_and_arrow:  belladonna's
// :ocean:  adalia
// :axe:  Carindrina
// :angel:  nira
// :heart_eyes_cat:  juniper
// :leopard:  aquilla
// :innocent: nira
// :badger:   scamander
// :smiley_cat:  Perry

// Options:
// * 1 post per person which then has all the numbers as reactions below each person. Clicking a reaction number registers them as that number. Ugly, spammy, might be harder to code since they are all separate messages
// * 1 embed, have a specific icon reaction for each player. Clicking each one places them in that order. Might be confusing/hard to know which icon is for which person
// * 1 embed, uses the next/previous paradigm to cycle though each player. Clicking the check reaction then puts that player next in the order
// Uses arrows https://www.reddit.com/r/Discord_Bots/comments/ah4rqr/help_with_an_embed_reaction_menu/

// https://github.com/jmsheff/discord-checkers
// https://github.com/bwmarrin/discordgo
// https://github.com/Necroforger/dgwidgets (might be useful for the third option)

func embed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "",
		Type:        "",
		Title:       "Lets Fight!",
		Description: "Roll initiatives!",
		Timestamp:   "",
		Color:       3447003,
		// Image:       &discordgo.MessageEmbedImage{},
		// Thumbnail:   &discordgo.MessageEmbedThumbnail{},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Players",
				Value: "The player names \n more names? :one: \\:one:",
			},
			{
				Name:  "Next",
				Value: "Next icon",
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "test footer text",
			// IconURL:      "",
		},
	}
}
