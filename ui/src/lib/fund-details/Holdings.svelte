<script lang="ts">
    import {afterUpdate, createEventDispatcher, onDestroy, onMount} from "svelte";
    import type {HoldingsResponse} from "$lib/proto/proto/HoldingsResponse";

    let observer: IntersectionObserver;
    let root: HTMLElement;

    export let holdings: HoldingsResponse[] = []

    const dispatch = createEventDispatcher()

    function endOfPageReached() {
        dispatch('endOfPageReached')
    }

    onDestroy(() => {
        if (observer) {
            observer.disconnect();
        }
    });

    afterUpdate(() => {
        if (observer) {
            observer.disconnect();
        }
        observer = new IntersectionObserver(async (entries) => {
            if (entries[0].isIntersecting) {
                endOfPageReached()
            }
        });
        observer.observe(root);
    })

</script>
<ul class="bg-gray-300 p-4">
    {#each holdings as {name, ticker, percentageOfTotal, type}, index}
        {#if (holdings.length < 200 ? index === Math.round(holdings.length * 0.7) : index === Math.round(holdings.length * 0.95))}
            <div class="ABC" bind:this={root}></div>
        {/if}
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
</ul>
