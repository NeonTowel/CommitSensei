<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<div class="min-h-screen flex flex-col items-center bg-white dark:bg-stone-950 font-sans">
  <!-- Header -->
  <header class="w-full flex flex-col items-center py-4 px-6 bg-amber-300">
    <h1 class="text-3xl font-bold text-gray-800 text-center mt-20 mb-15">
      <span class="mr-2 text-9xl">gitm<span class="text-pink-600">😊</span>ji</span>
    </h1>
    <p class="text-4xl font-sans text-slate-600 text-center mb-13">An emoji guide for your commit messages</p>
    <div class="flex space-x-4 mb-4">
      <button class="bg-gray-800 text-white px-4 py-2 rounded hover:bg-gray-700">GitHub</button>
    </div>
  </header>

  {@render children()}
</div>
