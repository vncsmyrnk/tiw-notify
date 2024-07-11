package notification

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"log"
)

func Notify(title, message string) {
	iconPath := ""

	err := beeep.Notify(title, message, iconPath)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	} else {
		fmt.Println("Notification sent successfully!")
	}
}
