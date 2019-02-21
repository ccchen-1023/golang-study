package main

import (
	"net/http"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://google.com",
	}

	c := make(chan string)
	// TIP: AKKA Router (Actor tell a message to another ActorRef via router)
	// https://doc.akka.io/docs/akka/2.5/routing.html

	for _, link := range links {
		go checkLink(link, c)

		// TIP: AKKA: tell
	}

	// fmt.Println(<-c)
	// TIP: C#, Java, Node.js: await
	// TIP: AKKA: ask

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	result := "up"
	if err != nil {
		result = "down"
	}
	c <- link + ": " + result
	// TIP: AKKA: tell
}
