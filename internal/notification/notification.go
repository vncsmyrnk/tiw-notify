package notification

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"log"
)

func Notify() {
	title := "Example notification"
	message := "This is an example"
	iconPath := ""

	err := beeep.Notify(title, message, iconPath)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	} else {
		fmt.Println("Notification sent successfully!")
	}
}
