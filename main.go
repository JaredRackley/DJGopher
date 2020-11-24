package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"./games"
	"./musicplayer"

	"github.com/bwmarrin/discordgo"
)

// Token - tring for the discord bot
const Token string = "NzcwMDAyMzExODc4OTM0NTI4.X5XOiQ.Z9F3_0y55l_VScYv7qx_zbV38rg"

var gameRunning = false
var musicRunning = false
var triviaGameRunning = false

// BotID is the unique id of the bot
var BotID string

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the runProgram func as a callback for MessageCreate events.
	dg.AddHandler(runProgram)

	//// Register ready as a callback for the ready events.
	//dg.AddHandler(ready)
	//
	//// Register guildCreate as a callback for the guildCreate events.
	//dg.AddHandler(guildCreate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	// Open the websocket and begin listening.
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
func runProgram(s *discordgo.Session, m *discordgo.MessageCreate) {
	testing := "765802303978340352"
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "g!ping" {
		s.ChannelMessageSend(testing, "Pong!")
		return
	}

	//TODO: have it play all music in music folder
	if m.Content == "g!playall" && !musicRunning {
		s.UpdateStatus(0, "all music!")
		musicplayer.MusicPlayer(s, m, "music/")
	}

	if m.Content == "m!stop" {
		musicRunning = false
		musicplayer.MusicPlayer(s, m, "")
	}

	if m.Content == "m!rickroll" && !musicRunning {
		musicRunning = true
		s.UpdateStatus(0, "rickroll!")
		musicplayer.MusicPlayer(s, m, "music/rickroll/")
	}

	if m.Content == "m!gag" && !musicRunning {
		musicRunning = true
		s.UpdateStatus(0, "gag music!")
		musicplayer.MusicPlayer(s, m, "music/gag/")
	}

	if m.Content == "m!jazz" && !musicRunning {
		musicRunning = true
		s.UpdateStatus(0, "jazz music!")
		musicplayer.MusicPlayer(s, m, "music/jazz/")
	}

	if m.Content == "m!pop" && !musicRunning {
		musicRunning = true
		s.UpdateStatus(0, "pop music!")
		musicplayer.MusicPlayer(s, m, "music/pop/")
	}

	if m.Content == "m!rock" && !musicRunning {
		musicRunning = true
		s.UpdateStatus(0, "rock music!")
		musicplayer.MusicPlayer(s, m, "music/rock/")
	}

	if m.Content == "m!alternative" && !musicRunning {
		musicRunning = true
		s.UpdateStatus(0, "alternative music!")
		musicplayer.MusicPlayer(s, m, "music/alternative/")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "g!pong" {
		s.UpdateStatus(0, "pong!")
		s.ChannelMessageSend(testing, "Ping!")
		return
	}

	// If the message is "pog" reply with ":gitpog:"
	if m.Content == "pog" {
		s.UpdateStatus(0, "pog!")
		s.ChannelMessageSend(testing, "<:gitpog:770159988915044352>")
		return
	}

	if m.Content == "g!hangman stop" {
		games.Restart(s, m)
		gameRunning = false
		return
	}

	if m.Content == "g!hangman" || gameRunning == true {
		s.UpdateStatus(0, "hangman!")
		if !gameRunning {
			games.Hangman(s, m, gameRunning)
			gameRunning = true
			return
		}
		games.Hangman(s, m, gameRunning)
		return
	}

	if m.Content == "g!trivia" || triviaGameRunning == true {
		s.UpdateStatus(0, "trivia!")
		if triviaGameRunning {
			games.Trivia(s, m, triviaGameRunning)
			triviaGameRunning = true
			return
		}
		triviaGameRunning = games.Trivia(s, m, triviaGameRunning)
		return
	}
}

//// This function will be called (due to AddHandler above) every time a new
//// guild is joined.
//func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
//
//	if event.Guild.Unavailable {
//		return
//	}
//
//	for _, channel := range event.Guild.Channels {
//		if channel.ID == event.Guild.ID {
//			_, _ = s.ChannelMessageSend(channel.ID, "Airhorn is ready! Type !airhorn while in a voice channel to play a sound.")
//			return
//		}
//	}
//}
//
////This function will be called (due to AddHandler above) when the bot receives
//// the "ready" event from Discord.
//func ready(s *discordgo.Session, event *discordgo.Ready) {
//
//	// Set the playing status.
//	s.UpdateStatus(0, "!TayTay")
//}
