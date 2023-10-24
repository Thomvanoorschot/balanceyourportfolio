<script lang="ts">
    import {getContext} from "svelte";
    import type {HoldingsStore} from "$lib/stores/FundFilterStore";
    import {debounce} from "$lib/utils.js";

    export let sectors: string[];

    const holdingsStore = getContext<HoldingsStore>("holdingsStore")
    let searchTerm: string= ""
    const filterHoldings = debounce(async function() {
        await holdingsStore.filter({searchTerm: searchTerm});
    }, 500)
    async function filterSector() {
        await holdingsStore.filter({sectorName: $holdingsStore.filter.sectorName});
    }

</script>

<form class="flex flex-col items-start justify-start p-10">
    <div class="flex items-center pb-5 w-full">
        <label for="company-search">Company:</label>
        <input
                bind:value={searchTerm}
                on:input={filterHoldings}
                id="company-search"
                class="ml-3 w-full rounded-md bg-gray-200 text-gray-700 leading-tight focus:outline-none py-2 px-2"
                type="search"
                placeholder="Company name or ticker">
    </div>
    <div class="flex items-center">
        <label for="sector-select">Sector:</label>
        <select
                bind:value={$holdingsStore.filter.sectorName}
                on:change={filterSector}
                id="sector-select"
                class=" ml-3 w-full rounded-md bg-gray-200 text-gray-700 leading-tight focus:outline-none py-2 px-2"
        >
            {#each sectors as sector}
                {#if sector === "Any sector"}
                    <option selected>{sector}</option>
                {:else}
                    <option>{sector}</option>
                {/if}
            {/each}
        </select>
    </div>
</form>