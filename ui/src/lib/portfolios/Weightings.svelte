<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import type {PortfolioFundSectorWeightings} from "$lib/portfolio";

    export let sectorWeightings: PortfolioFundSectorWeightings[];
    export let colorMap: Map<string, string>;
    interface Sector {
        cumulativePercentage: number;
        fundWeightings: FundWeighting[]
    }

    interface FundWeighting {
        fundName: string;
        percentage: number;
    }

    const sectorMap = new Map<string, Sector>()
    sectorWeightings.forEach((x) => {
        x.fundSectorWeighting.forEach((y) => {
            const sector = sectorMap.get(y.sectorName)
            if (sector) {
                sector.cumulativePercentage += y.percentage * x.percentageOfTotal
                sector.fundWeightings.push({
                    fundName: x.fundName,
                    percentage: y.percentage * x.percentageOfTotal,
                });
                return;
            }
            sectorMap.set(y.sectorName, {
                cumulativePercentage: y.percentage * x.percentageOfTotal,
                fundWeightings: [{
                    fundName: x.fundName,
                    percentage: y.percentage * x.percentageOfTotal,
                }]
            })
        })
    })
    const dispatch = createEventDispatcher()

    function sectorClicked(test: string) {
        console.log(test)
        // dispatch('sectorClicked',sw)
    }
</script>
<div class="flex flex-col gap-2">
    {#each [...sectorMap] as [key, value]}
        <div
                tabindex="0"
                aria-label=""
                role="button"
                on:keydown={() => { sectorClicked(key)}}
                on:click={() => { sectorClicked(key)}}
                class="flex flex-1 justify-start items-center cursor-pointer">
            <div class="w-1/4 text-xs">{key}</div>
            <div class="text-xs w-16">{Math.round(value.cumulativePercentage * 100) / 100}%</div>
            <div class="flex w-full">
                {#each value.fundWeightings as fundSector}
                    <div class="h-4 bg-blue-300"
                         style="
                         background-color: {colorMap.get(fundSector.fundName)};
                         width: { `${Math.round(fundSector.percentage / sectorMap.entries().next().value[1].cumulativePercentage * 100)}%`}">
                    </div>
                {/each}
            </div>
        </div>
    {/each}
</div>