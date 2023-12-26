<script lang="ts">
    import type {PageData} from './$types';
    import type {Holding} from "$lib/holding.ts";
    import type {ActionResult} from "@sveltejs/kit";
    import {enhance} from '$app/forms';
    import SearchBar from "$lib/search/SearchBar.svelte";
    import CheckButtonList from "$lib/filters/CheckButtonList.svelte";
    import {debounce} from "$lib/utils.ts";
    import List from "$lib/list/List.svelte";
    import ListItem from "$lib/list/ListItem.svelte";
    import HoldingLineItem from "$lib/list/HoldingLineItem.svelte";
    import type {FundSectorWeighting__Output} from "$lib/proto/proto/FundSectorWeighting.ts";
    import ColoredBarChart from "$lib/chart/ColoredBarChart.svelte";
    import ColoredBar from "$lib/chart/ColoredBar.svelte";
    import ColoredBarEntry from "$lib/chart/ColoredBarEntry.svelte";
    import DetailMenu from "$lib/menu/DetailMenu.svelte";
    import Information from "$lib/fund-details/Information.svelte";
    import type {FundInformation__Output} from "$lib/proto/proto/FundInformation.ts";
    import TertiaryButton from "$lib/shared/TertiaryButton.svelte";
    import Modal from "$lib/shared/Modal.svelte";
    import AddToPortfolioPopup from "$lib/portfolios/AddToPortfolioPopup.svelte";

    export let data: PageData;
    let error: string | undefined
    let sectors: string[] | undefined
    let fundSectorWeightings: FundSectorWeighting__Output[] | undefined
    let fundInformation: FundInformation__Output | null | undefined
    let holdings: Holding[] | undefined = []
    let fundsForm: HTMLFormElement;
    let searchTerm: string
    let resetSearch: boolean
    let selectedSectors: string[] = []
    let showModal = false
    $: ({sectors, fundInformation, fundSectorWeightings, holdings} = data);

    const updateNextPage = () => {
        return ({result}: {
            result: ActionResult
        }) => {
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
    const updateSelectedSectorsFromEvent = (clickEvent: CustomEvent<string>) => {
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
{#if (showModal)}
    <Modal bind:showModal>
        <AddToPortfolioPopup bind:showModal></AddToPortfolioPopup>
    </Modal>
{/if}
{#if (!error && sectors && fundSectorWeightings && holdings && fundInformation)}
    <div class="flex flex-grow items-start justify-between w-full gap-5 p-5">
        <DetailMenu>
            <TertiaryButton on:buttonClicked={() => showModal = !showModal}>Add to portfolio</TertiaryButton>
            <SearchBar
                    placeholder="Company name or ticker"
                    on:inputChanged={filterHoldings}
                    bind:value={searchTerm}
                    inPrimary="{false}"
            ></SearchBar>
            <CheckButtonList
                    title="Sectors"
                    list="{sectors}"
                    on:checkButtonClicked={updateSelectedSectorsFromEvent}
            >
            </CheckButtonList>
        </DetailMenu>
        <div class="flex flex-col flex-grow gap-5">
            <Information fundInformation="{fundInformation}"></Information>
            <ColoredBarChart>
                {#each fundSectorWeightings as fws, fswIndex}
                    <ColoredBar
                            title="{fws.sectorName}" percentage="{fws.percentage}"
                    >
                        <ColoredBarEntry
                                roundedLeft="{true}"
                                roundedRight="{true}"
                                color="#f582ae"
                                width="{Math.round(fws.percentage / fundSectorWeightings[0].percentage * 100)}"
                        ></ColoredBarEntry>
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
                        </ListItem>
                    {/each}
                </List>
            </form>
        </div>
    </div>
{/if}
