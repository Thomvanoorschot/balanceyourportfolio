<script>
    import {theme} from "$lib/stores/theme-store.ts";
    import {toggleTheme} from "$lib/theme/theme.ts";
    import SunIcon from "$lib/icons/SunIcon.svelte";
    import MoonIcon from "$lib/icons/MoonIcon.svelte";
</script>
<div class="flex">
    <button
            class="rounded-md focus:outline-none focus:shadow-outline-violet"
            aria-label="Toggle color mode"
            on:click|preventDefault={() => {
                                toggleTheme(theme, $theme);
                            }}
    >
        {#if ($theme.mode === 'light')}
            <SunIcon></SunIcon>
        {:else}
            <MoonIcon></MoonIcon>
        {/if}
    </button>
</div>