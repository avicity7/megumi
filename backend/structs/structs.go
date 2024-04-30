package structs

type License struct {
	Name string `json:"name"`
	Uri  string `json:"url"`
}

type AnimeSeason struct {
	Season string `json:"season"`
	Year   int    `json:"year"`
}

type Anime struct {
	Sources      []string    `json:"sources"`
	Title        string      `json:"title"`
	Type         []string    `json:"type"`
	Episodes     string      `json:"episodes"`
	Status       string      `json:"status"`
	AnimeSeason  AnimeSeason `json:"animeSeason"`
	Picture      string      `json:"picture"`
	Thumbnail    string      `json:"thumbnail"`
	Synonyms     []string    `json:"synonyms"`
	RelatedAnime []string    `json:"relatedAnime"`
	Tags         []string
}

type Root struct {
	License    License `json:"license"`
	Repository string  `json:"repository"`
	LastUpdate string  `json:"lastUpdate"`
	Data       []Anime `json:"data"`
}
