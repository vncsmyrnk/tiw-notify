// vim: noexpandtab

package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/vncsmyrnk/tiwnotify/internal/appointment"
	"github.com/vncsmyrnk/tiwnotify/internal/notification"
	"github.com/vncsmyrnk/tiwnotify/internal/schedule"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events
	go func() {
		jobScheduler := schedule.Schedule{}

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
					jobScheduler.StopAllJobs()

					as := appointment.AppointmentSchedule{Scheduler: &jobScheduler, Notifier: notification.BeeepNotifier{}}
					err = as.ScheduleFromFile(event.Name)
					if err != nil {
						log.Println("An error occurred while reading appointments.", err)
					}
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
