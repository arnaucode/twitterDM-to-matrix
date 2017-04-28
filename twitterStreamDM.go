package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
)

func stream(client *twitter.Client) {
	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()

	demux.All = func(message interface{}) {
		//fmt.Println(message)
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		if dm.SenderScreenName != twitterConfig.ScreenName {
			matrixSendMsg(dm.SenderScreenName, dm.Text, dm.CreatedAt)
		}
	}
	demux.Event = func(event *twitter.Event) {
		//fmt.Printf("%#v\n", event)
	}

	fmt.Println("Starting Stream...")

	streamUserParams := &twitter.StreamUserParams{}
	stream, err := client.Streams.User(streamUserParams)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	demux.HandleChan(stream.Messages)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}
