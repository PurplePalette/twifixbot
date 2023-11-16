package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token           string
	TwitterURLRegex *regexp.Regexp
	FxTwitterApiURL string
)

func init() {

	Token = os.Getenv("DISCORD_TOKEN")
	if Token == "" {
		fmt.Println("No token provided. Please set DISCORD_TOKEN environment variable.")
		os.Exit(1)
	}

	TwitterURLRegex = regexp.MustCompile(`https:\/\/(twitter\.com|x\.com)\/([a-zA-Z0-9_]+)\/status\/(\d+)`)
	FxTwitterApiURL = os.Getenv("FXTWITTER_API_URL")
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Get the all joined guilds.
	guilds, err := dg.UserGuilds(100, "", "")
	fmt.Println("Joined guilds:", len(guilds))
	if err != nil {
		fmt.Println("error getting guilds,", err)
		return
	}

	// Update the game status.
	err = dg.UpdateGameStatus(0, fmt.Sprintf("with %d guilds", len(guilds)))
	if err != nil {
		fmt.Println("error updating status,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
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

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if TwitterURLRegex.MatchString(m.Content) {
		for _, match := range TwitterURLRegex.FindAllStringSubmatch(m.Content, -1) {
			user_id := match[2]
			tweet_id := match[3]
			url := fmt.Sprintf("%s/%s/status/%s", FxTwitterApiURL, user_id, tweet_id)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("error getting response,", err)
				return
			}

			defer resp.Body.Close()
			byteArray, _ := io.ReadAll(resp.Body)

			jsonBytes := ([]byte)(byteArray)
			data := new(Response)

			if err := json.Unmarshal(jsonBytes, data); err != nil {
				fmt.Println("JSON Unmarshal error:", err)
				return
			}
			fmt.Println(data.Tweet.Text)

			embed := discordgo.MessageEmbed{
				Title:       fmt.Sprintf("%s(@%s)", data.Tweet.Author.Name, data.Tweet.Author.ScreenName),
				URL:         data.Tweet.URL,
				Description: data.Tweet.Text,
				Color:       0x1DA1F2,
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf("❤️: %d | 🔁: %d | 💬: %d | 👁️: %d", data.Tweet.Likes, data.Tweet.Retweets, data.Tweet.Replies, data.Tweet.Views),
				},
				Author: &discordgo.MessageEmbedAuthor{
					Name:    data.Tweet.Author.Name,
					URL:     data.Tweet.Author.URL,
					IconURL: data.Tweet.Author.AvatarURL,
				},
			}

			if data.Tweet.Media.Mosaic.Type == "mosaic_photo" {
				embed.Image = &discordgo.MessageEmbedImage{
					URL: data.Tweet.Media.Mosaic.Formats.JPEG,
				}
			} else if len(data.Tweet.Media.Photos) > 0 {
				embed.Image = &discordgo.MessageEmbedImage{
					URL:    data.Tweet.Media.Photos[0].URL,
					Height: data.Tweet.Media.Photos[0].Height,
					Width:  data.Tweet.Media.Photos[0].Width,
				}
			} else if len(data.Tweet.Media.Videos) > 0 {
				embed.Image = &discordgo.MessageEmbedImage{
					URL:    data.Tweet.Media.Videos[0].ThumbnailURL,
					Height: data.Tweet.Media.Videos[0].Height,
					Width:  data.Tweet.Media.Videos[0].Width,
				}
			}
			if data.Tweet.Quote != nil {
				embed.Description = fmt.Sprintf("%s\n\n\n↘️ Quoting %s(@%s)\n%s", data.Tweet.Text, data.Tweet.Quote.Author.Name, data.Tweet.Quote.Author.ScreenName, data.Tweet.Quote.Text)
				if embed.Image == nil && len(data.Tweet.Quote.Media.All) > 0 {
					if data.Tweet.Quote.Media.Mosaic.Type == "mosaic_photo" {
						embed.Image = &discordgo.MessageEmbedImage{
							URL: data.Tweet.Quote.Media.Mosaic.Formats.JPEG,
						}
					} else if len(data.Tweet.Quote.Media.Photos) > 0 {
						embed.Image = &discordgo.MessageEmbedImage{
							URL:    data.Tweet.Quote.Media.Photos[0].URL,
							Height: data.Tweet.Quote.Media.Photos[0].Height,
							Width:  data.Tweet.Quote.Media.Photos[0].Width,
						}
					} else if len(data.Tweet.Quote.Media.Videos) > 0 {
						embed.Image = &discordgo.MessageEmbedImage{
							URL:    data.Tweet.Quote.Media.Videos[0].ThumbnailURL,
							Height: data.Tweet.Quote.Media.Videos[0].Height,
							Width:  data.Tweet.Quote.Media.Videos[0].Width,
						}
					}
				}
			}

			// s.ChannelMessageSendEmbeds(m.ChannelID, embeds)
			// reply
			// s.ChannelMessageSendEmbedsReply(m.ChannelID, []*discordgo.MessageEmbed{&embed}, m.Reference())
			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
				AllowedMentions: &discordgo.MessageAllowedMentions{
					RepliedUser: false,
				},
				Embeds:    []*discordgo.MessageEmbed{&embed},
				Reference: m.Reference(),
			})
		}
	}
}
