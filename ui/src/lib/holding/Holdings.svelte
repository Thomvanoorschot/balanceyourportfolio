<script lang="ts">
    import {afterUpdate, createEventDispatcher, onDestroy, onMount} from "svelte";
    import type {Holding} from "$lib/holding";

    let observer: IntersectionObserver;
    let root: HTMLElement;

    export let holdings: Holding[] = []
    export let colorMap: Map<string, {
        fundName: string,
        color: string
    }> | undefined = undefined
    const dispatch = createEventDispatcher()

    function endOfPageReached() {
        dispatch('endOfPageReached')
    }

    onMount(() => {
        observer = new IntersectionObserver(async (entries) => {
            if (entries[0].isIntersecting && holdings.length >= 20) {
                endOfPageReached()
            }
        });
        observer.observe(root);
    })
    onDestroy(() => {
        if (observer) {
            observer.disconnect();
        }
    });

    afterUpdate(() => {
        // if (observer) {
        //     observer.disconnect();
        // }
        // observer = new IntersectionObserver(async (entries) => {
        //     if (entries[0].isIntersecting) {
        //         endOfPageReached()
        //     }
        // });
        // observer.observe(root);
    })
</script>
<div class="bg-gray-300 p-4">
    {#each holdings as {name, ticker, percentage, funds}, index}
        <!--{#if (holdings.length < 200 ? index === Math.round(holdings.length * 0.7) : index === Math.round(holdings.length * 0.95))}-->
        <!--    <div class="ABC" bind:this={root}></div>-->
        <!--{/if}-->

        <div class="border-gray-400 flex flex-col mb-2">
            <div class="cursor-pointer bg-gray-200 rounded-md items-center p-4  transition duration-500 ease-in-out transform hover:-translate-y-1 hover:shadow-lg">
                <div class=" flex flex-1">
                    <div class="flex flex-col rounded-md w-10 h-10 bg-gray-300 justify-center items-center mr-4">
                        <img src="/company-logos/{ticker}.png"
                             onerror="this.onerror=null; this.src='/company-logos/UNKNOWN.png'" alt="">
                    </div>
                    <div class="flex-1 pl-1 mr-16">
                        <div class="font-medium">{ticker}</div>
                        <div class="text-gray-600 text-sm">{name}</div>
                    </div>
                    <div class="text-gray-600 text-xs">{Math.round(percentage * 100) / 100}%</div>
                </div>
                {#if (funds?.length > 0)}
                    <div class="flex w-full h-2 mt-3">
                        {#each funds as fund}
                            <div class="h-full" style="
                         background-color: {colorMap?.get(fund.fundId)?.color};
                         width: { `${Math.round(fund.ratiodPercentage / percentage * 100)}%`}"></div>
                        {/each}
                    </div>
                {/if}
            </div>
        </div>
    {/each}
    <div class="ABC" bind:this={root}></div>

</div>
