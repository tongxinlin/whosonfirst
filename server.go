package main

import (
    "container/list"
    "fmt"
    "golang.org/x/net/websocket"
    "io"
    "net/http"
)

var conns *list.List
func Chat(ws *websocket.Conn) {
    defer ws.Close()
    item := conns.PushBack(ws)
    defer conns.Remove(item)
    var err error
    for {
        var data string
        if err = websocket.Message.Receive(ws, &data); err != nil {
            context = ""
            fmt.Printf("disconnected\n")
            break
        }
        SendMessage(item, fmt.Sprintf("%s", data))
    }
}

func SendMessage(self *list.Element, data string) {
    for item := conns.Front(); item != nil; item = item.Next() {
        ws, _ := item.Value.(*websocket.Conn)
        var readin string = data
        var response string = WhosOnFirst(context, readin, sL1, wL1, cL1, hL1)
        io.WriteString(ws, "<br/> Abbott: " + response + "<br/><br/>")
    }
}

func main() {
    SetCues();
    conns = list.New()
    http.Handle("/chat", websocket.Handler(Chat))
    http.Handle("/", http.FileServer(http.Dir("./html")))
    http.ListenAndServe(":8080", nil)
}


