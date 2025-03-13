<script lang="ts">
  export let emoji: string;
  export let code: string;
  export let description = 'Improve structure / format of the code.';
</script>

<div class={`rounded-2xl bg-white dark:bg-zinc-800 flex flex-col justify-center items-center w-[250px] h-[300px] m-4 p-4 pb-5`}>
  <div class="flex items-center justify-center h-[200px]">
    <span class="text-7xl">{emoji}</span>
  </div>
  <div class="text-stone-700 dark:text-stone-300 h-[200px] items-center justify-center w-full overflow-hidden">
    <p class="text-xl font-black break-words">{code}</p>
    <p class="mt-3 pb-3 text-l font-normal break-words">{description}</p>
  </div>
</div> 