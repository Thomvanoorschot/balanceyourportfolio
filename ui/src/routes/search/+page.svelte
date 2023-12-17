<script lang="ts">
    import type {PageData} from './$types';

    import {debounce} from "$lib/utils";
    import Result from "$lib/search/Result.svelte";

    export let data: PageData;
    $: ({funds} = data);

    let fElement: HTMLFormElement;

    const search = debounce(async function () {
        fElement.requestSubmit();
    }, 200)
</script>
<div class="flex flex-grow items-center justify-center w-full">
    <div class="bg-white shadow-md rounded-lg p-3 w-[50vw]">
        <div class="flex relative flex-col items-center bg-gray-200 rounded-md w-full">
            <div class="flex w-full items-center">
                <div class="pl-2">
                    <svg class="fill-current text-gray-500 w-6 h-6" xmlns="http://www.w3.org/2000/svg"
                         viewBox="0 0 24 24">
                        <path class="heroicon-ui"
                              d="M16.32 14.9l5.39 5.4a1 1 0 0 1-1.42 1.4l-5.38-5.38a8 8 0 1 1 1.41-1.41zM10 16a6 6 0 1 0 0-12 6 6 0 0 0 0 12z"/>
                    </svg>
                </div>
                <form
                        bind:this={fElement}
                        method="GET"
                        data-sveltekit-keepfocus
                        class="w-full"
                >
                    <input
                            name="searchTerm"
                            on:input={search}
                            class="w-full rounded-md bg-gray-200 text-gray-700 leading-tight focus:outline-none py-2 px-2"
                            type="search"
                            placeholder="ISIN, Name or Ticker"
                    >
                </form>
            </div>
            <ul id="searchResults" class="absolute top-12 w-full">
                {#each funds || [] as fund}
                    <Result href="/fund-details?fundId={fund.id}" fund="{fund}"></Result>
                {/each}
            </ul>
        </div>
    </div>
</div>
