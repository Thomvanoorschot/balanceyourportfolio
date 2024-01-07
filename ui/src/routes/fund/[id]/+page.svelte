<script lang="ts">
	import type { PageData } from './$types';
	import type { Holding } from '$lib/holding.ts';
	import type { ActionResult } from '@sveltejs/kit';
	import { enhance } from '$app/forms';
	import SearchBar from '$lib/search/SearchBar.svelte';
	import CheckButtonList from '$lib/filters/CheckButtonList.svelte';
	import { debounce } from '$lib/utils.ts';
	import List from '$lib/list/List.svelte';
	import ListItem from '$lib/list/ListItem.svelte';
	import HoldingLineItem from '$lib/list/HoldingLineItem.svelte';
	import type { FundSectorWeighting__Output } from '$lib/proto/proto/FundSectorWeighting.ts';
	import ColoredBarChart from '$lib/chart/ColoredBarChart.svelte';
	import ColoredBar from '$lib/chart/ColoredBar.svelte';
	import ColoredBarEntry from '$lib/chart/ColoredBarEntry.svelte';
	import type { FundInformation__Output } from '$lib/proto/proto/FundInformation.ts';
	import Modal from '$lib/shared/Modal.svelte';
	import AddToPortfolioPopup from '$lib/portfolios/AddToPortfolioPopup.svelte';
	import LoginOrRegister from '$lib/authentication/LoginOrRegister.svelte';
	import TertiaryContainer from '$lib/shared/TertiaryContainer.svelte';
	import MenuIcon from '$lib/icons/MenuIcon.svelte';
	import PrimaryButton from '$lib/shared/PrimaryButton.svelte';

	export let data: PageData;
	let error: string | undefined;
	let sectors: string[] | undefined;
	let fundSectorWeightings: FundSectorWeighting__Output[] | undefined;
	let fundInformation: FundInformation__Output | null | undefined;
	let holdings: Holding[] | undefined = [];
	let fundsForm: HTMLFormElement;
	let searchTerm: string;
	let resetSearch: boolean;
	let selectedSectors: string[] = [];
	let showAddToPortfolioModal: boolean = false;
	let showLoginOrRegisterModal: boolean = false;
	let isAuthenticated: boolean = false;
	$: ({ sectors, fundInformation, fundSectorWeightings, holdings, isAuthenticated } = data);

	const updateNextPage = () => {
		return ({ result }: { result: ActionResult }) => {
			if (result.type === 'success' && result?.data?.holdings && holdings) {
				if (resetSearch) {
					holdings = [...result?.data?.holdings];
					resetSearch = false;
					return;
				}
				holdings = [...holdings, ...result?.data?.holdings];
			} else if (result.type === 'failure') {
				error = result.data?.error;
			}
		};
	};

	function submitNextPage(): void {
		fundsForm.requestSubmit();
	}

	const filterHoldings = debounce(async function() {
		resetSearch = true;
		fundsForm.requestSubmit();
	}, 200);
	const setFilterForm = (formData: FormData) => {
		resetSearch
			? formData.set('holdingsLength', '0')
			: formData.set('holdingsLength', holdings!.length.toString());
		formData.set('selectedSectors', JSON.stringify(selectedSectors));
		formData.set('searchTerm', searchTerm);
	};
	const updateSelectedSectorsFromEvent = (clickEvent: CustomEvent<string>) => {
		if (selectedSectors.some((x) => x === clickEvent.detail)) {
			selectedSectors = selectedSectors.filter((x) => x !== clickEvent.detail);
			resetSearch = true;
			fundsForm.requestSubmit();
			return;
		}
		selectedSectors.push(clickEvent.detail);
		resetSearch = true;
		fundsForm.requestSubmit();
	};
	const handleAddToPortfolioClicked = () => {
		if (isAuthenticated) {
			showAddToPortfolioModal = true;
			return;
		}
		showLoginOrRegisterModal = true;
	};
	let expanded: boolean = false;
	let elem: HTMLDivElement;

	function growDiv() {
		if (expanded) {
			elem.style.height = '0';
			expanded = false;
		} else {
			const wrapper = document.querySelector('.measuringWrapper');
			elem.style.height = wrapper?.clientHeight + 'px';
			expanded = true;
		}
	}
</script>

{#if showAddToPortfolioModal}
	<Modal bind:showModal={showAddToPortfolioModal}>
		<AddToPortfolioPopup bind:showModal={showAddToPortfolioModal} />
	</Modal>
{/if}
<Modal bind:showModal={showLoginOrRegisterModal}>
	<LoginOrRegister />
</Modal>
{#if !error && sectors && fundSectorWeightings && holdings && fundInformation}
	<div class="flex flex-col gap-2 p-2">
		<ColoredBarChart>
			{#each fundSectorWeightings as fws, fswIndex}
				<ColoredBar title={fws.sectorName} percentage={fws.percentage}>
					<ColoredBarEntry
						roundedLeft={true}
						roundedRight={true}
						color="#f582ae"
						width={Math.round((fws.percentage / fundSectorWeightings[0].percentage) * 100)}
					/>
				</ColoredBar>
			{/each}
		</ColoredBarChart>
		<TertiaryContainer on:containerClicked={() => growDiv()}>
			<div>
				<div>
					<div class="flex items-center">
						<MenuIcon fillColor="fill-primary"></MenuIcon>
						<h1 class="text-l">{fundInformation.name}</h1>
					</div>
					<div
						bind:this={elem}
						class="h-0 transition-all ease-in-out duration-1000 overflow-hidden">
						<div class=" flex flex-col gap-5 measuringWrapper">
							<SearchBar
								placeholder="Company name or ticker"
								on:inputChanged={filterHoldings}
								bind:value={searchTerm}
								theme="primary"
							/>
							<CheckButtonList
								theme="primary"
								title="Sectors"
								list={sectors}
								on:checkButtonClicked={updateSelectedSectorsFromEvent}
							/>
							<PrimaryButton on:buttonClicked={handleAddToPortfolioClicked}>Add to portfolio</PrimaryButton>
						</div>
					</div>
				</div>
			</div>

		</TertiaryContainer>
		<div class="flex flex-col flex-grow gap-5">
			<form
				method="POST"
				action="?/filterHoldings"
				bind:this={fundsForm}
				use:enhance={({ formData }) => {
					setFilterForm(formData);
					return updateNextPage();
				}}
			>
				<List on:endOfPageReached={submitNextPage}>
					{#each holdings as holding}
						<ListItem>
							<HoldingLineItem {holding} />
						</ListItem>
					{/each}
				</List>
			</form>
		</div>
	</div>
{/if}
