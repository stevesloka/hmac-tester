package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	requestline := "url.foo.com"
	date := strings.Replace(time.Now().UTC().Format(time.RFC1123), "UTC", "GMT", 1)
	stringToSign := fmt.Sprintf("date: %s", date)

	key := []byte("8RSmn6QdOqXnOofofJ3i")
	h := hmac.New(sha1.New, key)
	h.Write([]byte(stringToSign))
	sigString := base64.StdEncoding.EncodeToString(h.Sum(nil))

	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://"+requestline, nil)
	hmacHeader := fmt.Sprintf("hmac username=\"slokas\", algorithm=\"hmac-sha1\", headers=\"date\", signature=\"%s\"", sigString)
	req.Header.Add("Authorization", hmacHeader)
	req.Header.Add("date", date)
	req.Header.Add("Host", requestline)

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
