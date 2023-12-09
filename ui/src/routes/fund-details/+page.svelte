<script lang="ts">
    import Filter from "$lib/fund-details/Filter.svelte";
    import Information from "$lib/fund-details/Information.svelte";
    import Weightings from "$lib/fund-details/Weightings.svelte";
    import Holdings from "$lib/holding/Holdings.svelte";

    import type {PageData} from './$types';
    import type {FundHoldingsFilter, FundSectorWeighting} from "$lib/fund";
    import {page} from "$app/stores";
    import {enhance} from '$app/forms';
    import type {ActionResult} from "@sveltejs/kit";
    import type {FundDetailsResponse__Output} from "$lib/proto/proto/FundDetailsResponse";
    import type {Holding} from "$lib/holding";

    export let data: PageData;

    let details: FundDetailsResponse__Output | undefined;
    let holdings: Holding[];
    let error: string | undefined
    $:{
        holdings = data?.holdings || []
        details = data?.details
        error = ""
    }

    let holdingsFilter: FundHoldingsFilter = {
        fundId: $page.url.searchParams.get("fundId")!,
        sectorName: "Any sector",
        searchTerm: "",
        limit: 20,
        offset: 0,
    }
    let searchForm: HTMLFormElement;
    let filterForm: HTMLFormElement;

    function submitNextPage(): void {
        searchForm.requestSubmit()
    }

    function filterSector(fsw: CustomEvent<FundSectorWeighting>) {
        holdingsFilter.sectorName = fsw.detail.sectorName
        submitFilter()
    }

    function submitFilter(): void {
        filterForm.requestSubmit()
    }

    const updateNextPage = () => {
        return ({result}: { result: ActionResult }) => {
            if (result.type === "success" && result?.data?.holdings) {
                holdings = [...holdings, ...result?.data?.holdings]
            } else if (result.type === "failure") {
                error = result.data?.error
            }
        };
    };

    const updateFilteredHoldings = () => {
        return ({result}: { result: ActionResult }) => {
            if (result.type === "success" && result?.data?.holdings) {
                holdings = [...result?.data?.holdings.entries]
            } else if (result.type === "failure") {
                error = result.data?.error
            }
        };
    };

    function setFilterForm(formData: FormData) {
        formData.set("fundId", holdingsFilter.fundId);
        formData.set("sectorName", holdingsFilter.sectorName);
        formData.set("searchTerm", holdingsFilter.searchTerm);
    }
</script>

{#if (!error && details && details.information)}
    <div class="flex flex-grow items-start justify-between w-full">
        <form
                method="POST"
                action="?/filterHoldings"
                bind:this={filterForm}
                use:enhance={({formData}) => {
                    setFilterForm(formData);
                    return updateFilteredHoldings();
                }}
        >
            {#if (details?.sectors)}
                <Filter
                        on:filterChanged={submitFilter}
                        bind:searchTerm={holdingsFilter.searchTerm}
                        bind:sectorName={holdingsFilter.sectorName}
                        sectors="{details.sectors || []}"
                ></Filter>
            {/if}
        </form>
        <div class="flex flex-col flex-grow">
            <div class="flex flex-col p-4">
                {#if (details?.information)}
                    <Information fundInformation="{details.information}"></Information>
                {/if}
            </div>
            <form
                    class="flex flex-col p-4"
                    method="POST"
                    action="?/filterHoldings"
                    bind:this={filterForm}
                    use:enhance={({formData}) => {
                    setFilterForm(formData);
                    return updateFilteredHoldings()
                }}
            >
                {#if (details?.sectorWeightings)}
                    <Weightings
                            on:sectorClicked={filterSector}
                            sectorWeightings="{details.sectorWeightings || []}"
                    ></Weightings>
                {/if}
            </form>
            <form
                    method="POST"
                    action="?/filterHoldings"
                    bind:this={searchForm}
                    use:enhance={({formData}) => {
                    setFilterForm(formData);
                    formData.set("holdingsLength", holdings.length.toString());
                    return updateNextPage()
                }}
            >
                <Holdings
                        on:endOfPageReached={submitNextPage}
                        holdings="{holdings || []}"
                ></Holdings>
            </form>
        </div>
    </div>
{:else}
    <h1>{error}</h1>
{/if}
