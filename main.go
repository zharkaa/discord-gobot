package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	// openai"github.com/sashabaranov/go-openai"
)

type Quote struct {
	Quote          string `json:"q"`
	Author         string `json:"a"`
	AuthorImage    string `json:"i"`
	CharacterCount string `json:"c"`
	HTMLFormat     string `json:"h"`
}

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
	// getQuotes()
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
	var resultQuote []Quote
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		fmt.Println("Error getting quote:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading body:", err)
	}

	if err := json.Unmarshal(body, &resultQuote); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	// fmt.Println(result)

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

	if m.Content == "!randomquote" {
		// _, err := s.ChannelMessageSend(m.ChannelID, resultQuote[0].Quote+" - "+resultQuote[0].Author)
		_, err := s.ChannelMessageSend(m.ChannelID, resultQuote[0].Quote+" - "+resultQuote[0].Author)
		
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}

// func getQuotes(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	var result []Quote
// 	resp, err := http.Get("https://zenquotes.io/api/random")
// 	if err != nil {
// 		fmt.Println("Error getting quote:", err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)

// 	if err != nil {
// 		fmt.Println("Error reading body:", err)
// 	}

// 	if err := json.Unmarshal(body, &result); err != nil {
// 		fmt.Println("Can not unmarshal JSON")
// 	}
// 	fmt.Println(result)

// }
