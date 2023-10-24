<script lang="ts">
    import type {FundSectorWeighting} from "$lib/fund";
    import {getContext} from "svelte";
    import type {HoldingsStore} from "$lib/stores/FundFilterStore";

    export let sectorWeightings: FundSectorWeighting[];

    const holdingsStore = getContext<HoldingsStore>("holdingsStore")
    async function filterSector(sectorName: string) {
        await holdingsStore.filter({sectorName: sectorName});
    }

</script>
<div id="sectorWeightings" class="flex flex-col gap-2">
    {#each sectorWeightings as {sectorName, percentage}, _}
        <div id="sector-weightings-SectorName"
             tabindex="0"
             aria-label=""
             role="button"
             on:keydown={async () => { await filterSector(sectorName)}}
             on:click={async () => { await filterSector(sectorName)}}
             class="flex flex-1 justify-start items-center cursor-pointer transition duration-100 ease-in-out transform hover:-translate-y-1">
            <div class="w-1/4 text-xs">{sectorName}</div>
            <div class="text-xs pr-1">{percentage}</div>
            <div class="flex w-full items-center">
                <div class="h-4 bg-blue-300 w-[{percentage}]"></div>
            </div>
        </div>
    {/each}
</div>