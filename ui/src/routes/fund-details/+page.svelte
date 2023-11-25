<script lang="ts">
    import Filter from "$lib/fund-details/Filter.svelte";
    import Information from "$lib/fund-details/Information.svelte";
    import Weightings from "$lib/fund-details/Weightings.svelte";
    import Holdings from "$lib/fund-details/Holdings.svelte";

    import type {PageData} from './$types';
    import type {FundHolding, FundHoldingsFilter} from "$lib/fund";
    import {page} from "$app/stores";

    export let data: PageData;
    $: ({details} = data);
    let holdings: FundHolding[] = [];
    let holdingsFilter: FundHoldingsFilter =  {
        fundId: $page.url.searchParams.get("fundId")!,
        sectorName: "Any sector",
        searchTerm: "",
        limit: 20,
        offset: 0,
    }

    async function nextPage(): Promise<void> {
        holdingsFilter.offset = holdings.length
        const holdingsResult = await fetch(`http://localhost:8080/api/v1/fund/holdings/filter`, {
            method: "POST",
            body: JSON.stringify(holdingsFilter)
        });
        holdings = [...holdings,...await holdingsResult.json()];
    }
    async function filter(): Promise<void> {
        const holdingsResult = await fetch(`http://localhost:8080/api/v1/fund/holdings/filter`, {
            method: "POST",
            body: JSON.stringify(holdingsFilter)
        });
        holdings = await holdingsResult.json();
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
        <div class="flex flex-col p-4">
            <Information fundInformation="{details.information}"></Information>
        </div>
        <div class="flex flex-col p-4">
            <Weightings
                    on:sectorClicked={filter}
                    sectorWeightings="{details.sectorWeightings}"
            ></Weightings>
        </div>
        <Holdings
                on:endOfPageReached={nextPage}
                holdings="{holdings}"
        ></Holdings>
    </div>
</div>