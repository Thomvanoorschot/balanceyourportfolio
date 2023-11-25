<script lang="ts">
    import {debounce} from "$lib/utils.js";
    import {createEventDispatcher} from 'svelte'

    export let sectors: string[];
    export let searchTerm: string;
    export let sectorName: string;

    const dispatch = createEventDispatcher()
    function filterChanged() {
        dispatch('filterChanged')
    }

    const filterHoldings = debounce(async function () {
        filterChanged()
    }, 500)
    filterChanged()
</script>

<div class="flex flex-col items-start justify-start p-10">
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
                bind:value={sectorName}
                on:change={filterChanged}
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
</div>