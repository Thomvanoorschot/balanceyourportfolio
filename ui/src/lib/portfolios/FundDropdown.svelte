<script>
import Result from "$lib/search/Result.svelte";
</script>
<ul
        use:clickOutside on:click_outside={handleClickOutside}
        class="absolute top-8 w-3/4 z-10">
    {#each funds as fund}
        <div tabindex="0"
             aria-label=""
             role="button"
             on:click={() => handleFundClicked(fund)}
             on:keydown={() => handleFundClicked(fund)}
        >
            <Result href="#" fund="{fund}"></Result>
        </div>
    {/each}
</ul>