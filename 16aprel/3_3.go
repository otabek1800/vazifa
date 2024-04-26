package main

import "fmt"

type Event struct {
    name string
    date string
}

type User struct {
    name string
    events []Event
}

func (u *User) addEvent(event Event) {
    u.events = append(u.events, event)
}

func (u *User) removeEvent(event_name string) {
    for i, event := range u.events {
        if event.name == event_name {
            u.events = append(u.events[:i], u.events[i+1:]...)
            break
        }
    }
}

func (u *User) updateEvent(event_name string, new_date string) {
    for i, event := range u.events {
        if event.name == event_name {
            u.events[i].date = new_date
            break
        }
    }
}

func main() {
    user := User{name: "Otabek: ", events: []Event{}}

    event1 := Event{name: "Tadbir1: ", date: "2022-03-22"}
    event2 := Event{name: "Tadbir2: ", date: "2022-03-23"}
    event3 := Event{name: "Dars1: ", date: "2022-03-24"}

    user.addEvent(event1)
    user.addEvent(event2)
    user.addEvent(event3)

    fmt.Println(user)

    user.removeEvent("Tadbir1: ")

    fmt.Println(user)

    user.updateEvent("Dars1: ", "2022-03-25")

    fmt.Println(user)
}