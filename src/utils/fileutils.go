package utils

import "os"

func FileExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}