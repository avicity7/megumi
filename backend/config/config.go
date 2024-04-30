package config

import (
	"encoding/json"
	"fmt"
	"io"
	"megumi/structs"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache
var AnimeArray []structs.Anime

func Connect(r *chi.Mux) {
	Cache = cache.New(5*time.Minute, 10*time.Minute)
	jsonFile, err := os.Open("../../anime-offline-database/anime-offline-database.json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error opening file")
	}

	fmt.Println("Opened")
	bytes, _ := io.ReadAll(jsonFile)

	var root structs.Root

	json.Unmarshal(bytes, &root)

	defer jsonFile.Close()

	AnimeArray = root.Data
}
