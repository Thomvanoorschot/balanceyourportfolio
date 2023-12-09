<script lang="ts">
    import Result from "$lib/search/Result.svelte";
    import {debounce} from "$lib/utils";
    import {clickOutside} from "$lib/custom-svelte-typings";
    import type {Fund} from "$lib/fund";
    import {onMount} from "svelte";
    import type {PortfolioListItem__Output} from "$lib/proto/proto/PortfolioListItem";
    import {enhance} from '$app/forms';
    import type {SearchFundsEntry__Output} from "$lib/proto/proto/SearchFundsEntry";
    import type {ActionResult} from "@sveltejs/kit";

    export let listItem: PortfolioListItem__Output;

    let searchForm: HTMLFormElement;
    let showList = false;
    let funds: SearchFundsEntry__Output[]
    $: {
        funds = []
    }
    const search = debounce(async function () {
        if (listItem.name === "") {
            funds = [];
            return;
        }
        searchForm.requestSubmit()
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
        method="POST"
        action="?/searchFunds"
        bind:this={searchForm}
        use:enhance={({formData}) => {
                formData.set("searchTerm", listItem.name)
                return updateFunds();
        }}
        class="w-full relative"
>
    <div class="flex">
        <input
                bind:value={listItem.fundId}
                class="hidden" type="text"
        >
        <input
                bind:value={listItem.name}
                on:input={search}
                on:click={() => showList = true}
                class="w-3/4" type="text" placeholder="Ticker or name"
        >
        <input
                bind:value={listItem.amount}
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
</form>
