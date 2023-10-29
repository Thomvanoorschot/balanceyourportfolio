<script lang="ts">
    import Filter from "$lib/fund-details/Filter.svelte";
    import Information from "$lib/fund-details/Information.svelte";
    import Weightings from "$lib/fund-details/Weightings.svelte";
    import Holdings from "$lib/fund-details/Holdings.svelte";

    import type {PageData} from './$types';
    import {createHoldingsStore} from "$lib/stores/fund-filter-store";
    import {setContext} from "svelte";
    import {page} from "$app/stores";

    export let data: PageData;
    $: ({details} = data);

    const holdingsStore = createHoldingsStore($page.url.searchParams.get("fundId")!)
    setContext("holdingsStore", holdingsStore)
</script>

<div id="fundDetails" class="flex flex-grow items-start justify-between w-full">
    <Filter sectors="{details.sectors}"></Filter>
    <div class="flex flex-col flex-grow">
        <div class="flex flex-col p-4">
            <Information fundInformation="{details.information}"></Information>
        </div>
        <div class="flex flex-col p-4">
            <Weightings sectorWeightings="{details.sectorWeightings}"></Weightings>
        </div>
        <Holdings></Holdings>
    </div>
</div>