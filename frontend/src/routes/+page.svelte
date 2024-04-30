<script lang="ts">
  import type { Anime } from '../types/global.js';
  let search: string,
    results: Array<Anime>
  export let data
  const { backend_uri } = data
  const searchAnime = async () => {
    const response = await fetch(`${backend_uri}:3000/anime/get-anime/${search}`)
    const parsed = await response.json()
    results = parsed
  }
</script>

<div class="flex flex-col items-center p-8">
  <form on:submit|preventDefault={searchAnime} class="flex flex-col">
    <input class="outline outline-stone-300 rounded p-1" type="text" bind:value={search}>
    <button name="submit" class="mt-4 bg-pink-500 hover:bg-pink-600 rounded text-white">Search</button>
  </form>

  {#if results}
    <div class="grid grid-cols-3">
      {#each results as anime}
        <a href={anime.sources[0]} class="mt-8">
          <h1 class="font-semibold text-lg">{anime.title}</h1>
          <img src={anime.picture} alt="headerImage" class="h-[20vh] object-fit"/>
          <h2>{anime.animeSeason.year} {anime.animeSeason.season}</h2>
        </a>
      {/each}
    </div>
  {/if}
</div>

