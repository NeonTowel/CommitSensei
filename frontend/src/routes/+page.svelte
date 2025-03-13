<script lang="ts">
  import EmojiCard from '$lib/EmojiCard.svelte';
  import data from '$lib/gitmojis.json';

  let searchQuery = ''; // Track search input

  // Filter emojis based on search query
  const filteredEmojis = () => {
    return data.gitmojis.filter(({ description, code }) =>
      description.toLowerCase().includes(searchQuery.toLowerCase()) ||
      code.toLowerCase().includes(searchQuery.toLowerCase())
    );
  };
</script>

<div class="flex flex-col items-center justify-center w-full max-w-6xl mx-auto p-6">
  <input
    type="text"
    placeholder="Search your gitmoji..."
    class="border border-gray-300 rounded px-4 py-2 w-full max-w-md focus:outline-none focus:border-pink-400 mb-6 text-stone-700 dark:text-stone-300"
    bind:value={searchQuery}
  />
  {#if filteredEmojis().length > 0}
    <main class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-x-25 gap-y-10">
      {#each filteredEmojis() as { emoji, description, code }}
        <EmojiCard emoji={emoji} code={code} description={description}/>
      {/each}
    </main>
  {:else}
    <p class="text-center text-gray-600 mt-6">No results found</p>
  {/if}
</div>

<style>
  main {
    text-align: center;
    padding: 1em;
  }
</style>
