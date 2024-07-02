package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

/*
	d973711 -->{
				ID: "d973711",
				OriginalURL : "https://github.com/bhaveshranjan",
				ShortURL : "d973711"
				CreationDate : time.Now()
	}
*/

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) //It converts the originalURL string to a byte slice
	fmt.Println("hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println("hasher data: ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("hasher EncodedTOSTRING data: ", hash)
	fmt.Println("Final String is: ", hash[:8])
	return hash[:8]
}

func main() {
	fmt.Println("Starting URL shortner...")
	OriginalURL := "https://github.com/bhaveshranjan"
	generateShortURL(OriginalURL)
}
