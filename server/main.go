package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    "strings"
    "sync"
)

var (
    upgrader = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all origins
    }
    messageHistory []string
    mu             sync.Mutex
)

func reverseString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    for {
        _, msg, err := ws.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }
        receivedMessage := string(msg)
        fmt.Println("Received:", receivedMessage)

        if receivedMessage == "history" {
            mu.Lock()
            history := strings.Join(messageHistory, ", ")
            mu.Unlock()
            if err := ws.WriteMessage(websocket.TextMessage, []byte(history)); err != nil {
                log.Println("Error sending message:", err)
            }
            continue
        }

        reversedMsg := reverseString(receivedMessage)

        mu.Lock()
        if len(messageHistory) >= 5 {
            messageHistory = messageHistory[1:] // Remove the oldest message
        }
        messageHistory = append(messageHistory, receivedMessage)
        mu.Unlock()

        if err := ws.WriteMessage(websocket.TextMessage, []byte(reversedMsg)); err != nil {
            log.Println("Error sending message:", err)
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleConnections)

    fmt.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
