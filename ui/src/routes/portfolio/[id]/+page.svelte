<script lang="ts">
    import type {PageData} from './$types';
    import Weightings from "$lib/portfolios/Weightings.svelte";
    import Filter from "$lib/fund-details/Filter.svelte";
    import {page} from "$app/stores";
    import type {PortfolioHoldingsFilter, PortfolioSectorWeighting} from "$lib/portfolio";
    import FundColors from "$lib/portfolios/FundColors.svelte";
    import type {FundInformation__Output} from "$lib/proto/proto/FundInformation";
    import Holdings from "$lib/holding/Holdings.svelte";
    import type {Holding} from "$lib/holding";
    import type {ActionResult} from "@sveltejs/kit";
    import {enhance} from '$app/forms';
    import TableHeaderRow from "$lib/table/TableHeaderRow.svelte";
    import SearchFundCell from "$lib/table/SearchFundCell.svelte";
    import TableRow from "$lib/table/TableRow.svelte";
    import TableHeader from "$lib/table/TableHeader.svelte";
    import Table from "$lib/table/Table.svelte";
    import NumberCell from "$lib/table/NumberCell.svelte";
    import EditCell from "$lib/table/EditCell.svelte";
    import AddButton from "$lib/shared/AddButton.svelte";
    import TextCell from "$lib/table/TextCell.svelte";
    import FilterList from "$lib/filters/FilterList.svelte";
    import SearchBar from "$lib/search/SearchBar.svelte";
    import CheckButtonList from "$lib/filters/CheckButtonList.svelte";
    import {debounce} from "$lib/utils.ts";

    export let data: PageData;
    let colorMap: Map<string, {fundName: string, color: string}> | undefined
    let error: string | undefined
    let sectors: string[] | undefined
    let fundInformation: FundInformation__Output[] | undefined
    let portfolioFundSectorWeightings: PortfolioSectorWeighting[] | undefined
    let holdings: Holding[] = []
    let fundsForm: HTMLFormElement;
    $: ({sectors, fundInformation, portfolioFundSectorWeightings, colorMap, holdings} = data);

    const updateNextPage = () => {
        return ({result}: { result: ActionResult }) => {
            if (result.type === "success" && result?.data?.holdings && holdings) {
                holdings = [...holdings, ...result?.data?.holdings]
            } else if (result.type === "failure") {
                error = result.data?.error
            }
        };
    };

    let holdingsFilter: PortfolioHoldingsFilter = {
        portfolioId: "",
        sectorName: "Any sector",
        searchTerm: "",
        limit: 20,
        offset: 0,
    }
    function submitNextPage(): void {
        fundsForm.requestSubmit()
    }
    const filterHoldings = debounce(async function () {
        searchForm.requestSubmit();
    }, 200)
    function setFilterForm(formData: FormData) {
        formData.set("sectorName", holdingsFilter.sectorName);
        formData.set("searchTerm", holdingsFilter.searchTerm);
    }
</script>
{#if (!error && sectors && colorMap && portfolioFundSectorWeightings && holdings)}
    <div class="flex flex-grow items-start justify-between w-full">
        <FilterList>
            <SearchBar placeholder="Company name or ticker" on:inputChanged={filterHoldings}></SearchBar>
            <CheckButtonList
                    title="Sectors"
                    list="{['Technology', 'Consumer Discretionary', 'HealthCare', 'Financials']}"
            >
            </CheckButtonList>
        </FilterList>
        <div class="flex flex-col flex-grow">
            <div class="pl-4 sticky top-0 bg-white w-full z-50">
                <FundColors colorMap="{colorMap}"></FundColors>
            </div>
            <div class="flex flex-col p-4">
                <Weightings colorMap="{colorMap}"
                            sectorWeightings="{portfolioFundSectorWeightings}"></Weightings>
            </div>
            <form
                    method="POST"
                    action="?/filterHoldings"
                    bind:this={fundsForm}
                    use:enhance={({formData}) => {
                    setFilterForm(formData);
                    formData.set("holdingsLength", holdings.length.toString());
                    return updateNextPage()
                }}
            >
                <Holdings
                        on:endOfPageReached={submitNextPage}
                        holdings="{holdings}"
                        colorMap="{colorMap}"
                ></Holdings>
            </form>
        </div>
    </div>
{/if}
