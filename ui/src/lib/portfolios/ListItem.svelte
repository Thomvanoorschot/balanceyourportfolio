<script lang="ts">
    import Result from "$lib/search/Result.svelte";
    import {debounce} from "$lib/utils";
    import {clickOutside} from "$lib/custom-svelte-typings";
    import type {Fund} from "$lib/fund";
    import {onMount} from "svelte";
    import type {PortfolioListItem} from "$lib/portfolio";

    export let listItem: PortfolioListItem;

    let showList = false;
    let funds: Fund[] = []
    const search = debounce(async function () {
        if (listItem.name === "") {
            funds = [];
            return;
        }
        const f = await fetch(`http://localhost:8080/api/v1/fund/search?searchTerm=${listItem.name}`);
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
        listItem.fundId = f.id;
        listItem.name = f.name;
    }
    import { createEventDispatcher } from 'svelte'
    const dispatch = createEventDispatcher()

    function blurField() {
        dispatch('blurField')
    }
</script>
<div class="w-full relative">
    <div class="flex">
        <input
                bind:value={listItem.fundId}
                class="hidden" type="text"
        >
        <input
                bind:value={listItem.name}
                on:input={search}
                on:click={() => showList = true}
                on:blur={blurField}
                class="w-3/4" type="text" placeholder="Ticker or name"
        >
        <input
                bind:value={listItem.amount}
                on:blur={blurField}
                class="w-1/4" type="number" placeholder="Amount"
        >
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
