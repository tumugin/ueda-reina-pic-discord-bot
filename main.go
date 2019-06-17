package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
	"syscall"
	"ueda-reina-pic-discord-bot/reinalibs"
)

var authToken string
var reinaJSONPath string
var reinaPics reinalibs.UedaReinaPics

func main() {
	app := cli.NewApp()
	app.Name = "UedaReinaPicDiscordBot"
	app.Description = "UESHAMAAAAAAAAAAAAAA\n" +
		"You can add this BOT to your guild with this URL: https://discordapp.com/oauth2/authorize?client_id=<YOUR_CLIENT_ID_HERE>&scope=bot&permissions=522304"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token",
			Usage:       "Discord authnication token.",
			Destination: &authToken,
		},
		cli.StringFlag{
			Name:        "json",
			Usage:       "JSON file contains Ueda Reina image urls.",
			Value:       "reina.json",
			Destination: &reinaJSONPath,
		},
	}
	app.Action = func(c *cli.Context) error {
		if authToken == "" {
			return fmt.Errorf("please specify Discord authnication token with --token option")
		}
		initBot()
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initBot() {
	reinaPics = reinalibs.NewUedaReinaPics(reinaJSONPath)
	discord, err := discordgo.New("Bot " + authToken)
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(onMessage)
	if err := discord.Open(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func onMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if !reinalibs.IsReinaCalling(message.Content) {
		return
	}
	reinaPicUrl := reinaPics.GetRandomReinaPic()
	if _, err := session.ChannelMessageSend(message.ChannelID, reinaPicUrl); err != nil {
		log.Print(err)
	}
}
