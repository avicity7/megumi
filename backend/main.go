package main

import (
	"encoding/json"
	"fmt"
	"io"
	"megumi/structs"
	"os"
	"slices"
	"strings"
)

func main() {
	jsonFile, err := os.Open("../../anime-offline-database/anime-offline-database.json")
	if err != nil {
		fmt.Println("Error opening file")
	}

	fmt.Println("Opened")
	bytes, _ := io.ReadAll(jsonFile)

	var root structs.Root

	json.Unmarshal(bytes, &root)

	animeArray := root.Data
	idx := slices.IndexFunc(animeArray, func(anime structs.Anime) bool { return strings.Contains(anime.Title, "Re:LIFE") })
	fmt.Println(idx)

	defer jsonFile.Close()
}
