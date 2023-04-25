package fileutils

import (
    "errors"
    "io"
    "net/http"
    "os"
)

func DownloadFileFromUrl(url string, filepath string) error {
    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return errors.New("ERROR: Unable to create copy of file")
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return errors.New("ERROR: Unable to fetch data from URL")
    }
    defer resp.Body.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return errors.New("ERROR: File download Error")
    }

    return nil
}