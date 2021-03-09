package websocket

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func ClientStart(max int, channelName string) {

	// path := "/sub/" + channelName + "/" + strconv.Itoa(max)

	u := url.URL{Scheme: "ws", Host: "nchan:80", Path: "/sub"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go receive(c, done)
	waitloop(c, done)
}

func receive(c *websocket.Conn, done chan struct{}) {
	defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

func waitloop(c *websocket.Conn, done chan struct{}) {
	//停止用
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-done:
			log.Println("done")
			return
			/*
				case <-interrupt:
					log.Println("interrupt")
					err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					if err != nil {
						log.Println("write close:", err)
						return
					}
					select {
					case <-done:
					case <-time.After(time.Second):
					}
					return
			*/
		}
	}
}
