package main

import (
	"lld/design_patterns/observer/publisher"
	"lld/design_patterns/observer/subscriber"
)

func main() {
    channel := publisher.NewChannel("TechWorld")
    aditya := subscriber.NewUserSubscriber("Aditya")
    rahul := subscriber.NewUserSubscriber("Rahul")
    priya := subscriber.NewUserSubscriber("Priya")
    
    channel.Subscribe(aditya)
    channel.Subscribe(rahul)
    channel.Subscribe(priya)
    
    channel.UploadVideo("LLD Pattern")
}




