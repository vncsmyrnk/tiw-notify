// vim: noexpandtab

package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/vncsmyrnk/tiwnotify/internal/appointment"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events
	go func() {
		var appointments []appointment.Appointment
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)

					// Stops all existent jobs
					for _, a := range appointments {
						a.StopJob()
					}

					appointments, err = appointment.ScheduleAppointmentNotificationsFromFile(event.Name)
					if err != nil {
						log.Println("An error occurred while reading appointments.", err)
					}

					log.Println(len(appointments), "appointments read and scheduled")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/home/dev/.local/share/todayiwill")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever
	<-make(chan struct{})
}
