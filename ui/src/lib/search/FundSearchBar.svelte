<script lang="ts">
import Result from "$lib/search/Result.svelte";
import SearchBar from "$lib/search/SearchBar.svelte";
import {debounce} from "$lib/utils.ts";
import type {SearchFundsEntry__Output} from "$lib/proto/proto/SearchFundsEntry.ts";
import type {ActionResult} from "@sveltejs/kit";
import {enhance} from '$app/forms';

let funds: SearchFundsEntry__Output[] = []

let searchForm: HTMLFormElement;
const search = debounce(async function () {
    searchForm.requestSubmit();
}, 200)
const updateFunds = () => {
    return ({result}: { result: ActionResult }) => {
        if (result.type === "success" && result?.data?.funds) {
            funds = result?.data?.funds
        } else if (result.type === "failure") {
            // error = result.data?.error
        }
    };
};
</script>
<form
        bind:this={searchForm}
        method="POST"
        action="/search"
        data-sveltekit-keepfocus
        class="w-full"
        use:enhance={updateFunds}
>
    <SearchBar
            placeholder="Search for funds"
            on:inputChanged={search}
    >
        <ul id="searchResults" class="absolute top-12 w-full">
            {#each funds || [] as fund}
                <Result href="/fund-details?fundId={fund.id}" fund="{fund}"></Result>
            {/each}
        </ul>
    </SearchBar>
</form>