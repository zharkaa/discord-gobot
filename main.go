package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	// openai"github.com/sashabaranov/go-openai"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token := os.Getenv("TOKEN")
	// Create a new Discord session with a bot token
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Add an event handler for the ready event
	dg.AddHandler(onReady)
	dg.AddHandler(onMessageSend)

	// Open a websocket connection to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl-C to exit.")

	// Wait for a signal to interrupt the bot (Ctrl-C)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close the Discord session before exiting
	dg.Close()
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	// Print a message to indicate that the bot is ready
	fmt.Println("Bot is ready.")
}

func onMessageSend(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	// Check if the message content is "ping"
	if m.Content == "ping" {
		// Send a "pong" message to the same channel
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}

func get_Quotes() {
	response, err := http.NewRequest("GET", "https://zenquotes.io/api/random", nil)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
	fmt.Println(response)

}