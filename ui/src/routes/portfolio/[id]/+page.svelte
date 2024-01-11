<script lang="ts">
	import type { PageData } from './$types';
	import type { PortfolioSectorWeighting } from '$lib/portfolio';
	import FundColors from '$lib/portfolios/FundColors.svelte';
	import type { Holding } from '$lib/holding';
	import type { ActionResult } from '@sveltejs/kit';
	import { enhance } from '$app/forms';
	import SearchBar from '$lib/search/SearchBar.svelte';
	import CheckButtonList from '$lib/filters/CheckButtonList.svelte';
	import { debounce } from '$lib/utils.ts';
	import List from '$lib/list/List.svelte';
	import ListItem from '$lib/list/ListItem.svelte';
	import HoldingLineItem from '$lib/list/HoldingLineItem.svelte';
	import ColoredHoldingsBar from '$lib/portfolios/ColoredHoldingsBar.svelte';
	import ColoredBarChart from '$lib/chart/ColoredBarChart.svelte';
	import ColoredBar from '$lib/chart/ColoredBar.svelte';
	import ColoredBarEntry from '$lib/chart/ColoredBarEntry.svelte';
	import DetailMenu from '$lib/menu/DetailMenu.svelte';
	import TertiaryContainer from '$lib/shared/TertiaryContainer.svelte';
	import MenuIcon from '$lib/icons/MenuIcon.svelte';

	export let data: PageData;
	let colorMap: Map<string, { fundName: string; color: string }> | undefined;
	let error: string | undefined;
	let sectors: string[] | undefined;
	let portfolioFundSectorWeightings: PortfolioSectorWeighting[] | undefined;
	let holdings: Holding[] | undefined = [];
	let fundsForm: HTMLFormElement;
	let searchTerm: string;
	let resetSearch: boolean;
	let selectedSectors: string[] = [];
	$: ({ sectors, portfolioFundSectorWeightings, colorMap, holdings } = data);

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
	const updateSelectedSectors = (clickEvent: CustomEvent<string>) => {
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

{#if !error && sectors && colorMap && portfolioFundSectorWeightings && holdings}
	<!-- Mobile -->
	<div class="lg:hidden flex flex-col gap-2 p-2">
		<div class="sticky top-2 z-10">
			<TertiaryContainer>
				<FundColors {colorMap} />
			</TertiaryContainer>
		</div>
		<ColoredBarChart>
			{#each portfolioFundSectorWeightings as pfsw}
				<ColoredBar title={pfsw.sectorName} percentage={pfsw.weighting.totalPercentage}>
					{#each pfsw.weighting.fundSectorWeighting as fsw, fswIndex}
						<ColoredBarEntry
							color={colorMap.get(fsw.fundId)?.color || ''}
							width={
									Math.round(
										(fsw.percentage / portfolioFundSectorWeightings[0].weighting.totalPercentage) *
											100
									)}
						/>
					{/each}
				</ColoredBar>
			{/each}
		</ColoredBarChart>
		<TertiaryContainer on:containerClicked={() => growDiv()}>
			<div>
				<div>
					<div class="flex items-center">
						<MenuIcon fillColor="fill-primary"></MenuIcon>
						<h1 class="align-middle">Filters</h1>
					</div>
					<div
						bind:this={elem}
						class="h-0 transition-all ease-in-out duration-1000 overflow-hidden">
						<div class="flex flex-col gap-5 measuringWrapper">
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
								on:checkButtonClicked={updateSelectedSectors}
							/>
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
							<ColoredHoldingsBar {holding} {colorMap} />
						</ListItem>
					{/each}
				</List>
			</form>
		</div>
	</div>

	<!-- PC -->
	<div class="hidden w-full gap-2 lg:flex lg:visible p-5">
		<div class="sticky top-0">
			<DetailMenu>
				<SearchBar
					placeholder="Company name or ticker"
					on:inputChanged={filterHoldings}
					bind:value={searchTerm}
					theme="primary"
				/>
				<CheckButtonList
					title="Sectors"
					list={sectors}
					on:checkButtonClicked={updateSelectedSectors}
				/>
				<FundColors {colorMap} />
			</DetailMenu>
		</div>
		<div class="flex flex-col flex-grow gap-5">
			<ColoredBarChart>
				{#each portfolioFundSectorWeightings as pfsw}
					<ColoredBar
						width="{Math.round((pfsw.weighting.totalPercentage / portfolioFundSectorWeightings[0].weighting.totalPercentage) * 100)}"

						title={pfsw.sectorName}
						percentage={pfsw.weighting.totalPercentage}
					>
						{#each pfsw.weighting.fundSectorWeighting as fsw, fswIndex}
							<ColoredBarEntry
								color={colorMap.get(fsw.fundId)?.color || ''}
								width={
									Math.round(
										(fsw.percentage / portfolioFundSectorWeightings[0].weighting.totalPercentage) *
											100
									)
								}
							/>
						{/each}
					</ColoredBar>
				{/each}
			</ColoredBarChart>
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
							<ColoredHoldingsBar {holding} {colorMap} />
						</ListItem>
					{/each}
				</List>
			</form>
		</div>
	</div>
{/if}
