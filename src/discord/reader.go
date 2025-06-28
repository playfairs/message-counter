package discord

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

func GetChannelFolders(root string) ([]string, error) {
    var channels []string
    files, err := ioutil.ReadDir(root)
    if err != nil {
        return nil, err
    }
    for _, f := range files {
        if f.IsDir() && strings.HasPrefix(f.Name(), "c") {
            channels = append(channels, filepath.Join(root, f.Name()))
        }
    }
    return channels, nil
}

func GetJSONFiles(channelPath string) ([]string, error) {
    var jsonFiles []string
    files, err := ioutil.ReadDir(channelPath)
    if err != nil {
        return nil, err
    }
    for _, f := range files {
        if strings.HasSuffix(f.Name(), ".json") {
            jsonFiles = append(jsonFiles, filepath.Join(channelPath, f.Name()))
        }
    }
    return jsonFiles, nil
}