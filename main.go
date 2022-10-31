package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Time time.Time
}

func main() {

	// Example d'utilisation de local
	//voir dans /usr/local/go/lib/time/zoneinfo.zip pour les timezones dispo
	location, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return
	}
	time1 := time.Now().In(location)
	fmt.Println(time1)

	event1 := Event{
		Time: time.Now().UTC(),
	}
	bytes, err := json.Marshal(event1)
	if err != nil {
		return
	}
	var event2 Event
	err = json.Unmarshal(bytes, &event2)
	if err != nil {
		return
	}
	fmt.Println(event1.Time.Equal(event2.Time))
	fmt.Println(event1)
	fmt.Println(event2)

}
