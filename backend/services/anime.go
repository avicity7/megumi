package services

import (
	"fmt"
	"megumi/config"
	"megumi/structs"
	"strings"
	"sync"

	"github.com/IGLOU-EU/go-wildcard/v2"
)

func GetAnime(title string) ([]structs.Anime, error) {
	outputArray := make([]structs.Anime, 0)
	var (
		mu = &sync.Mutex{}
		wg sync.WaitGroup
	)
	wg.Add(len(config.AnimeArray))
	for _, anime := range config.AnimeArray {
		go func(anime structs.Anime) {
			defer wg.Done()
			if wildcard.Match("*"+strings.ToLower(title)+"*", strings.ToLower(anime.Title)) {
				mu.Lock()
				outputArray = append(outputArray, anime)
				mu.Unlock()
			} else {
				for _, synonym := range anime.Synonyms {
					if wildcard.Match("*"+strings.ToLower(title)+"*", strings.ToLower(synonym)) {
						mu.Lock()
						outputArray = append(outputArray, anime)
						mu.Unlock()
						break
					}
				}
			}
		}(anime)
	}
	wg.Wait()
	if len(outputArray) > 0 {
		return outputArray, nil
	}
	return []structs.Anime{}, fmt.Errorf("no anime found")
}
