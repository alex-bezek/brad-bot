package main

import "github.com/bwmarrin/discordgo"

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
