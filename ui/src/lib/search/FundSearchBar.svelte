<script lang="ts">
    import Result from "$lib/search/Result.svelte";
    import SearchBar from "$lib/search/SearchBar.svelte";
    import {debounce} from "$lib/utils.ts";
    import type {FilterFundsResponseEntry__Output} from "$lib/proto/proto/FilterFundsResponseEntry.ts";

    let funds: FilterFundsResponseEntry__Output[] = []

    let searchForm: HTMLFormElement;
    export let value: string | undefined = ""

    const search = debounce(async function () {
        const resp = await fetch("/api/search-funds", {
            method: 'POST',
            body: JSON.stringify({
                value
            })
        })
        if (resp.ok) {
            funds = await resp.json() as FilterFundsResponseEntry__Output[]
        } else if (!resp.ok) {
            // error = result.data?.error
        }
    }, 200)

</script>

<SearchBar
        placeholder="Search for funds"
        on:inputChanged={search}
        bind:value={value}
>
    <ul id="searchResults" class="absolute top-12 left-0 right-0 w-full">
        {#each funds || [] as fund}
            <Result href="/fund/{fund.id}" fund="{fund}"></Result>
        {/each}
    </ul>
</SearchBar>