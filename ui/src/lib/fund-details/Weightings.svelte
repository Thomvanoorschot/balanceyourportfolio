<script lang="ts">
    import type {FundSectorWeighting} from "$lib/fund";
    import {createEventDispatcher} from "svelte";

    export let sectorWeightings: FundSectorWeighting[];
    const relativeHoldingsList = new Map();
    sectorWeightings.map((x) => relativeHoldingsList.set(x.sectorName, `${Math.round(x.percentage / sectorWeightings[0].percentage * 100)}%`))
    const dispatch = createEventDispatcher()
    function sectorClicked(sw: FundSectorWeighting) {
        dispatch('sectorClicked',sw)
    }
</script>
<div id="sectorWeightings" class="flex flex-col gap-2">
    {#each sectorWeightings as sectorWeighting}
        <div
                tabindex="0"
                aria-label=""
                role="button"
                on:keydown={() => { sectorClicked(sectorWeighting)}}
                on:click={() => { sectorClicked(sectorWeighting)}}
                class="flex flex-1 justify-start items-center cursor-pointer transition duration-100 ease-in-out transform hover:-translate-y-1">
            <div class="w-1/4 text-xs">{sectorWeighting.sectorName}</div>
            <div class="text-xs w-16">{Math.round(sectorWeighting.percentage * 100) / 100}%</div>
            <div class="flex w-full items-center">
                <div class="h-4 bg-blue-300" style="width: {relativeHoldingsList.get(sectorWeighting.sectorName)}"></div>
            </div>
        </div>
    {/each}
</div>