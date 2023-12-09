<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import type {PortfolioSectorWeighting} from "$lib/portfolio";

    export let sectorWeightings:  PortfolioSectorWeighting[]
    export let colorMap: Map<string, {fundName: string, color: string}>;

    interface FundWeighting {
        fundName: string;
        percentage: number;
    }

    const dispatch = createEventDispatcher()

    function sectorClicked(test: string) {
        // dispatch('sectorClicked',sw)
    }
</script>
<div class="flex flex-col gap-2">
    {#each sectorWeightings as sectorWeighting}
        <div
                tabindex="0"
                aria-label=""
                role="button"
                on:keydown={() => { sectorClicked(sectorWeighting.sectorName)}}
                on:click={() => { sectorClicked(sectorWeighting.sectorName)}}
                class="flex flex-1 justify-start items-center cursor-pointer">
            <div class="w-1/4 text-xs">{sectorWeighting.sectorName}</div>
            <div class="text-xs w-16">{Math.round(sectorWeighting.weighting.totalPercentage * 100) / 100}%</div>
            <div class="flex w-full">
                {#each sectorWeighting.weighting.fundSectorWeighting as fundSector}
                    <div class="h-4 bg-blue-300"
                         style="
                         background-color: {colorMap.get(fundSector.fundId)?.color};
                         width: { `${Math.round(fundSector.percentage / sectorWeightings[0].weighting.totalPercentage * 100)}%`}">
                    </div>
                {/each}
            </div>
        </div>
    {/each}
</div>