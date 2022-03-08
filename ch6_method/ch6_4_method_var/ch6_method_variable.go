package main

import (
	"fmt"
	"time"
)

type Rocket struct {
	Name string
}

func (r *Rocket) Launch() {
	fmt.Println(r.Name, " launched")
}

func main() {
	r := Rocket{Name: "r"}
	c := make(chan bool)
	go func() {
		time.AfterFunc(time.Second*3, func() {
			r.Launch()
			c <- true
		})
	}()
	fmt.Printf("Is %s launched? %v\n", r.Name, <-c)
}
