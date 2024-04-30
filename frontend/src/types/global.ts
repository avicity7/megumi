export type AnimeSeason = {
  season: string,
  year: number
}

export type Anime = {
  sources: Array<string>
  title: string,
  type: string,
  episodes: number,
  status: string,
  animeSeason: AnimeSeason,
  picture: string,
  thumbnail: string,
  synonyms: Array<string>,
  relatedAnime: Array<string>,
  tags: Array<string> 
}