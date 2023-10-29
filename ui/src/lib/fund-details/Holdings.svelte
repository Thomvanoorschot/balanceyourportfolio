<script lang="ts">
    import {getContext, onDestroy, onMount} from "svelte";
    import type {HoldingsStore} from "$lib/stores/fund-filter-store";

    let observer: IntersectionObserver;
    let root: HTMLElement;

    const holdingsStore = getContext<HoldingsStore>("holdingsStore")
    onMount(() => {
        observer = new IntersectionObserver(async (entries) => {
            if (entries[0].isIntersecting) {
                await holdingsStore.nextPage();
            }
        });
        observer.observe(root);
    });

    onDestroy(() => {
        if (observer) {
            observer.disconnect();
        }
    });

</script>
<ul class="bg-gray-300 p-4">
    {#each $holdingsStore.holdings as {name, ticker, percentageOfTotal, type}, _}
        <li class="border-gray-400 flex flex-row mb-2">
            <div class="cursor-pointer bg-gray-200 rounded-md flex flex-1 items-center p-4  transition duration-500 ease-in-out transform hover:-translate-y-1 hover:shadow-lg">
                <div class="flex flex-col rounded-md w-10 h-10 bg-gray-300 justify-center items-center mr-4">
                    {#if type === "Currency"}
                        $
                    {:else}
                        <img src="/company-logos/{ticker}.png"
                             onerror="this.onerror=null; this.src='/company-logos/UNKNOWN.png'" alt="">
                    {/if}
                </div>
                <div class="flex-1 pl-1 mr-16">
                    <div class="font-medium">{ticker}</div>
                    <div class="text-gray-600 text-sm">{name}</div>
                </div>
                <div class="text-gray-600 text-xs">{percentageOfTotal}</div>
            </div>
        </li>
    {/each}
    <div bind:this={root}></div>
</ul>
