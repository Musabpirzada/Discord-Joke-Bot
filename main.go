package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"GoProject.me/ACP Project/x/mux"
	"github.com/bwmarrin/discordgo"
)

func main() {

	var Router = mux.New()
	Session, err := discordgo.New("Bot MTA1NDAwMjM5ODg2MDkzNTE5OA.GMNGkP.miVKomYj4cWU1u-dWUMyIEGN8qr1gTTZNbCcj8") //add your token in it
	Session.AddHandler(Router.OnMessageCreate)
	///Register Routes
	Router.Route("joke", "This is a joke", Router.Joke)

	err = Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()

}
