<script lang="ts">
    import Result from "$lib/search/Result.svelte";
    import {debounce} from "$lib/utils";
    import {clickOutside} from "$lib/custom-svelte-typings";
    import type {Fund} from "$lib/fund";
    import {onMount} from "svelte";

    let showList = false;
    let funds: Fund[] = []
    let searchTerm: string;
    const search = debounce(async function () {
        if (searchTerm === "") {
            funds = [];
            return;
        }
        const f = await fetch(`http://localhost:8080/api/v1/fund/search?searchTerm=${searchTerm}`);
        funds = await f.json();
    }, 500)

    onMount(() => {
        document.addEventListener('keydown', closeOnEscape);
        return () => document.removeEventListener('keydown', closeOnEscape);
    });

    const closeOnEscape = (e: KeyboardEvent) => {
        if (e.key === 'Escape')
            showList = false;
    }

    const handleClickOutside = () => {
        showList = false;
    }
    const handleFundClicked = (f: Fund) => {
        showList = false;
        searchTerm = f.name
    }
</script>
<div class="w-full relative">
    <div class="flex">
        <input
                bind:value={searchTerm}
                on:input={search}
                on:click={() => showList = true}
                class="w-3/4" type="text" placeholder="Ticker or name"
        >
        <input class="w-1/4" type="text" placeholder="Amount">
    </div>
    {#if showList}
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
    {/if}
</div>
