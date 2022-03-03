package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
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
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Printf("Got message: %s \n From: %s\n", m.Content, m.Author.Username)

	if strings.Contains(m.Content, "pixie") && strings.Contains(m.Content, "stick") {
		s.ChannelMessageSend(m.ChannelID, "https://giphy.com/gifs/range-of-emotions-cotton-candy-girl-safeco-field-foodie-3o6Ztg2MgUkcXyCgtG")
	}

	spellCmd := fmt.Sprintf("<@!%s> spell ", s.State.User.ID)

	if strings.HasPrefix(m.Content, spellCmd) {
		spell := strings.ReplaceAll(m.Content, spellCmd, "")

		files, err := ioutil.ReadDir("./spells")
		if err != nil {
			log.Fatal(err)
		}

		var spellFile fs.FileInfo

	FileLoop:
		for _, f := range files {
			spellFileName := strings.ToLower(f.Name())
			for _, t := range strings.Split(spell, " ") {
				if !strings.Contains(spellFileName, strings.ToLower(t)) {
					continue FileLoop
				}
			}
			spellFile = f
			break
		}

		if spellFile == nil {
			s.ChannelMessageSend(m.ChannelID, "Spell not found")
		} else {
			fmt.Println(spellFile.Name())
			f, err := os.Open("./spells/" + spellFile.Name())
			if err != nil {
				panic(err)
			}
			defer f.Close()

			s.ChannelFileSend(m.ChannelID, spellFile.Name(), f)
		}
	}

	reactCmd := fmt.Sprintf("<@!%s> react ", s.State.User.ID)
	if strings.HasPrefix(m.Content, reactCmd) {
		fmt.Printf("Got reaction command from %s", m.Author.Username)
		react := strings.ReplaceAll(m.Content, reactCmd, "")
		characters := map[string]string{
			"Kenzie":          "Adalia",
			"twylawolf":       "Balladonna",
			"silvershoes":     "Carindrina",
			"furus":           "Furus",
			"kittythewildcat": "Galia",
			"Alicia":          "Midna",
			"wolfswing":       "Nira",
			"gettingvetted":   "Venus",
		}

		emotion := map[string]string{
			"angry":      "Angry",
			"confused":   "Confused",
			"dead":       "Dead",
			"got it":     "Figured it out",
			"initiative": "Game on",
			"happy":      "Happy",
			"innocent":   "Innocent",
			"inspired":   "Inspired",
			"love":       "Love",
			"ugh":        "Nauseous",
			"nope":       "Oh fuck",
			"sad":        "Sad",
			"star":       "Star",
		}

		if react == "help" {
			keys := make([]string, len(emotion))

			i := 0
			for k := range emotion {
				keys[i] = k
				i++
			}
			s.ChannelMessageSend(m.ChannelID, "React with one of the following: "+strings.Join(keys, ", "))
			return
		}

		c := characters[m.Author.Username]
		e := emotion[react]
		if c == "" || e == "" {
			fmt.Printf("Did not find character emotion for %s %s \n", m.Author.Username, react)
		} else {
			fileName := c + " " + e + ".PNG"
			f, err := os.Open("./avatars/Reactions/" + fileName)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			s.ChannelFileSend(m.ChannelID, fileName, f)
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s %s", c, e))
		}

	}
}
