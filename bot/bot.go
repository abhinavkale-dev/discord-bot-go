package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/abhinavkale-dev/go-discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	botSession *discordgo.Session
	botID      string
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Replies with Pong!",
	},
	{
		Name:        "start",
		Description: "Starts the process",
	},
	{
		Name:        "dance",
		Description: "Make the bot dance 💃",
	},
}

func Start() {
	var err error
	botSession, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	user, err := botSession.User("@me")
	if err != nil {
		fmt.Println("Error fetching bot user info:", err)
		return
	}
	botID = user.ID

	botSession.AddHandler(onReady)

	botSession.AddHandler(onInteraction)

	err = botSession.Open()
	if err != nil {
		fmt.Println("Error connecting to Discord:", err)
		return
	}
	fmt.Println("Bot is now running ✅")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	for _, cmd := range commands {
		_ = botSession.ApplicationCommandDelete(botID, "", cmd.ID)
	}

	botSession.Close()
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Println("Registering slash commands...")
	for _, cmd := range commands {
		created, err := s.ApplicationCommandCreate(botID, "", cmd)
		if err != nil {
			fmt.Println("Failed to create command", cmd.Name, ":", err)
		} else {
			cmd.ID = created.ID
			fmt.Println("Registered command:", cmd.Name)
		}
	}
}

func onInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "ping":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})

	case "start":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Starting up… 🚀",
			},
		})

	case "dance":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "🕺💃 Watch me dance!",
			},
		})
	}
}
