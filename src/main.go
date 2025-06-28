package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type Message struct {
    ID        interface{} `json:"ID"`
    Timestamp string      `json:"Timestamp"`
    Contents  string      `json:"Contents"`
    Attachments string    `json:"Attachments"`
}

func main() {
    root := "./package/Messages"
    totalMessages := 0
    totalChannels := 0

    files, err := ioutil.ReadDir(root)
    if err != nil {
        fmt.Println("Error reading root directory:", err)
        os.Exit(1)
    }

    for _, f := range files {
        if f.IsDir() && strings.HasPrefix(f.Name(), "c") {
            channelPath := filepath.Join(root, f.Name())
            msgCount := countMessagesInChannel(channelPath)
            totalChannels++
            totalMessages += msgCount
        }
    }

    fmt.Printf("You have sent %d messages across %d channels\n", totalMessages, totalChannels)
}

func countMessagesInChannel(channelPath string) int {
    count := 0
    files, err := ioutil.ReadDir(channelPath)
    if err != nil {
        fmt.Println("Error reading channel:", err)
        return 0
    }
    for _, f := range files {
        if f.Name() == "messages.json" {
            filePath := filepath.Join(channelPath, f.Name())
            data, err := ioutil.ReadFile(filePath)
            if err != nil {
                fmt.Println("Error reading file:", filePath, err)
                continue
            }
            var messages []Message
            if err := json.Unmarshal(data, &messages); err != nil {
                fmt.Println("Error parsing JSON in file:", filePath, err)
                continue
            }
            count += len(messages)
        }
    }
    return count
}