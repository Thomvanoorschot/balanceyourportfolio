<script lang="ts">
    import type {PageData} from './$types';
    import Weightings from "$lib/portfolios/Weightings.svelte";
    import Filter from "$lib/fund-details/Filter.svelte";
    import {page} from "$app/stores";
    import type {PortfolioHoldingsFilter} from "$lib/portfolio";
    import {colors, stringToRandomInteger} from "$lib/utils";
    import FundColors from "$lib/portfolios/FundColors.svelte";

    export let data: PageData;
    const {details} = data;

    let holdingsFilter: PortfolioHoldingsFilter = {
        portfolioId: $page.url.searchParams.get("portfolioId")!,
        sectorName: "Any sector",
        searchTerm: "",
        limit: 20,
        offset: 0,
    }

    const colorMap = new Map<string, string>()
    let colorsCopy = [...colors]
    details.portfolioFundSectorWeightings.forEach((x) => {
        const color = colorMap.get(x.fundName)
        if (!color) {
            const randomColor = colorsCopy[stringToRandomInteger(x.fundName, colorsCopy.length)]
            colorMap.set(x.fundName, randomColor)
            colorsCopy = colorsCopy.filter(x => x != randomColor)
        }
    });

    async function filter() {

    }
</script>
<div class="flex flex-grow items-start justify-between w-full">
    <Filter
            on:filterChanged={filter}
            bind:searchTerm={holdingsFilter.searchTerm}
            bind:sectorName={holdingsFilter.sectorName}
            sectors="{details.sectors}"
    ></Filter>
    <div class="flex flex-col flex-grow">
        <div class="pl-4">
            <FundColors colorMap="{colorMap}"></FundColors>
        </div>
        <div class="flex flex-col p-4">
            <Weightings colorMap="{colorMap}" sectorWeightings="{details.portfolioFundSectorWeightings}"></Weightings>
        </div>
        <!--        <Holdings></Holdings>-->
    </div>
</div>