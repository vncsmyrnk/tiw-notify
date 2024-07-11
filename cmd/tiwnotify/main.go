// vim: noexpandtab

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/vncsmyrnk/tiwnotify/internal/schedule"
)

type appointment struct {
	time        time.Time
	description string
}

func parseAppointmentFromString(str string) (appointment, error) {
	parts := strings.SplitN(str, " ", 2)
	now := time.Now()
	time, err := time.Parse(time.RFC3339, fmt.Sprintf("%4d-%02d-%02dT%v:00Z", now.Year(), int(now.Month()), now.Day(), parts[0]))
	if err != nil {
		return appointment{}, err
	}
	return appointment{
		time:        time,
		description: parts[1],
	}, nil
}

func readAppointmentsFromFile(fileName string) ([]appointment, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []appointment{}, err
	}

	var appointments []appointment
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		a, err := parseAppointmentFromString(line)
		if err != nil {
			continue
		}
		appointments = append(appointments, a)
	}
	return appointments, nil
}

func main() {
	t, err := time.Parse(time.RFC3339, "2024-07-11T21:50:00Z")
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	job := schedule.Job{
		Time: t,
		Task: func() { fmt.Println("Job executed at", time.Now().Format(time.Stamp)) },
	}

	schedule.AddJob(job)
	time.Sleep(4 * time.Minute)
}

func main_2() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
					appointments, err := readAppointmentsFromFile(event.Name)
					if err != nil {
						log.Println("An error occurred while reading appointments.", err)
					}
					log.Println(appointments)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("/home/dev/.local/share/todayiwill")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
