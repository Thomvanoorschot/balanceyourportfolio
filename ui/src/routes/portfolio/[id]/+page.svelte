<script lang="ts">
    import type {PageData} from './$types';
    import type {PortfolioSectorWeighting} from "$lib/portfolio";
    import FundColors from "$lib/portfolios/FundColors.svelte";
    import type {FundInformation__Output} from "$lib/proto/proto/FundInformation";
    import type {Holding} from "$lib/holding";
    import type {ActionResult} from "@sveltejs/kit";
    import {enhance} from '$app/forms';
    import SearchBar from "$lib/search/SearchBar.svelte";
    import CheckButtonList from "$lib/filters/CheckButtonList.svelte";
    import {debounce} from "$lib/utils.ts";
    import List from "$lib/list/List.svelte";
    import ListItem from "$lib/list/ListItem.svelte";
    import HoldingLineItem from "$lib/list/HoldingLineItem.svelte";
    import ColoredHoldingsBar from "$lib/portfolios/ColoredHoldingsBar.svelte";
    import ColoredBarChart from "$lib/chart/ColoredBarChart.svelte";
    import ColoredBar from "$lib/chart/ColoredBar.svelte";
    import ColoredBarEntry from "$lib/chart/ColoredBarEntry.svelte";
    import DetailMenu from "$lib/menu/DetailMenu.svelte";

    export let data: PageData;
    let colorMap: Map<string, { fundName: string, color: string }> | undefined
    let error: string | undefined
    let sectors: string[] | undefined
    let fundInformation: FundInformation__Output[] | undefined
    let portfolioFundSectorWeightings: PortfolioSectorWeighting[] | undefined
    let holdings: Holding[] | undefined = []
    let fundsForm: HTMLFormElement;
    let searchTerm: string
    let resetSearch: boolean
    let selectedSectors: string[] = []
    $: ({sectors, fundInformation, portfolioFundSectorWeightings, colorMap, holdings} = data);

    const updateNextPage = () => {
        return ({result}: { result: ActionResult }) => {
            if (result.type === "success" && result?.data?.holdings && holdings) {
                if (resetSearch) {
                    holdings = [...result?.data?.holdings]
                    resetSearch = false
                    return
                }
                holdings = [...holdings, ...result?.data?.holdings]
            } else if (result.type === "failure") {
                error = result.data?.error
            }
        };
    };

    function submitNextPage(): void {
        fundsForm.requestSubmit()
    }

    const filterHoldings = debounce(async function () {
        resetSearch = true
        fundsForm.requestSubmit();
    }, 200)
    const setFilterForm = (formData: FormData) => {
        resetSearch ? formData.set("holdingsLength", "0") : formData.set("holdingsLength", holdings!.length.toString())
        formData.set("selectedSectors", JSON.stringify(selectedSectors));
        formData.set("searchTerm", searchTerm);
    }
    const updateSelectedSectors = (clickEvent: CustomEvent<string>) => {
        if (selectedSectors.some(x => x === clickEvent.detail)) {
            selectedSectors = selectedSectors.filter(x => x !== clickEvent.detail)
            resetSearch = true
            fundsForm.requestSubmit()
            return
        }
        selectedSectors.push(clickEvent.detail)
        resetSearch = true
        fundsForm.requestSubmit()
    }

</script>
{#if (!error && sectors && colorMap && portfolioFundSectorWeightings && holdings)}
    <div id="page" class="flex flex-grow items-start justify-between w-full gap-5 p-5">
        <div class="sticky top-0">
            <DetailMenu>
                <SearchBar placeholder="Company name or ticker"
                           on:inputChanged={filterHoldings}
                           bind:value={searchTerm}
                           inPrimary="{false}"
                ></SearchBar>
                <CheckButtonList
                        title="Sectors"
                        list="{sectors}"
                        on:checkButtonClicked={updateSelectedSectors}
                >
                </CheckButtonList>
                <FundColors colorMap="{colorMap}"></FundColors>
            </DetailMenu>
        </div>
        <div class="flex flex-col flex-grow gap-5">
            <ColoredBarChart>
                {#each portfolioFundSectorWeightings as pfsw}
                    <ColoredBar title="{pfsw.sectorName}" percentage="{pfsw.weighting.totalPercentage}">
                        {#each pfsw.weighting.fundSectorWeighting as fsw, fswIndex}
                            <ColoredBarEntry
                                    roundedLeft="{fswIndex === 0}"
                                    roundedRight="{fswIndex === pfsw.weighting.fundSectorWeighting.length - 1}"
                                    color="{colorMap.get(fsw.fundId)?.color || ''}"
                                    width="{Math.max(Math.round(fsw.percentage / portfolioFundSectorWeightings[0].weighting.totalPercentage * 100), 0.01)}"
                            ></ColoredBarEntry>
                        {/each}
                    </ColoredBar>
                {/each}
            </ColoredBarChart>
            <form
                    method="POST"
                    action="?/filterHoldings"
                    bind:this={fundsForm}
                    use:enhance={({formData}) => {
                    setFilterForm(formData);
                    return updateNextPage()
                }}
            >
                <List on:endOfPageReached={submitNextPage}>
                    {#each holdings as holding}
                        <ListItem>
                            <HoldingLineItem holding="{holding}"></HoldingLineItem>
                            <ColoredHoldingsBar holding="{holding}" colorMap="{colorMap}"></ColoredHoldingsBar>
                        </ListItem>
                    {/each}
                </List>
            </form>
        </div>
    </div>
{/if}
