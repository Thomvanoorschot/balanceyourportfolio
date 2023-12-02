<script lang="ts">
    import type {PageData} from './$types';
    import Weightings from "$lib/portfolios/Weightings.svelte";
    import Filter from "$lib/fund-details/Filter.svelte";
    import {page} from "$app/stores";
    import type {PortfolioHoldingsFilter} from "$lib/portfolio";
    import {colors, stringToRandomInteger} from "$lib/utils";
    import FundColors from "$lib/portfolios/FundColors.svelte";
    import type {PortfolioDetailsResponse__Output} from "$lib/proto/proto/PortfolioDetailsResponse";

    export let data: PageData;
    let details: PortfolioDetailsResponse__Output | undefined;
    let colorMap: Map<string, string> | undefined
    let error: string | undefined
    $:{
        details = data?.details
        colorMap = data?.colorMap || new Map<string, string>()
        error = ""
    }

    let holdingsFilter: PortfolioHoldingsFilter = {
        portfolioId: $page.url.searchParams.get("portfolioId")!,
        sectorName: "Any sector",
        searchTerm: "",
        limit: 20,
        offset: 0,
    }

    async function filter() {

    }
</script>
{#if (!error && details && colorMap)}
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
                <Weightings colorMap="{colorMap}"
                            sectorWeightings="{details.portfolioFundSectorWeightings}"></Weightings>
            </div>
            <!--        <Holdings></Holdings>-->
        </div>
    </div>
{/if}
