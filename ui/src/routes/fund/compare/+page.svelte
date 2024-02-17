<script lang="ts">
	import type { PageData } from './$types';
	import TertiaryContainer from '$lib/shared/TertiaryContainer.svelte';
	import FundSearchBar from '$lib/search/FundSearchBar.svelte';
	import type { FilterFundsResponseEntry__Output } from '$lib/proto/proto/FilterFundsResponseEntry.ts';
	import QuaternaryButton from '$lib/shared/QuaternaryButton.svelte';
	import { enhance } from '$app/forms';
	import type { ActionResult } from '@sveltejs/kit';
	import { page } from '$app/stores';
	import ListItem from '$lib/list/ListItem.svelte';
	import List from '$lib/list/List.svelte';
	import OverlappingHoldingLineItem from '$lib/list/OverlappingHoldingLineItem.svelte';
	import PulldownElement from '$lib/shared/PulldownElement.svelte';
	import NonOverlappingHoldingLineItem from '$lib/list/NonOverlappingHoldingLineItem.svelte';
	import ColoredBarEntry from '$lib/chart/ColoredBarEntry.svelte';
	import ColoredBar from '$lib/chart/ColoredBar.svelte';
	import ColoredBarChart from '$lib/chart/ColoredBarChart.svelte';
	import FundColors from '$lib/portfolios/FundColors.svelte';
	import * as d3 from 'd3';

	export let data: PageData;
	let compareForm: HTMLFormElement;
	let overlappingHoldingsElem: PulldownElement;
	let fundOneNonOverlappingHoldingsElem: PulldownElement;
	let fundTwoNonOverlappingHoldingsElem: PulldownElement;
	let sectorOverviewElem: PulldownElement;

	const updateFundOne = (clickEvent: CustomEvent<FilterFundsResponseEntry__Output>) => {
		$page.url.searchParams.set('fundOne', clickEvent.detail.id);
		history.replaceState(history.state, '', $page.url);
	};
	const updateFundTwo = (clickEvent: CustomEvent<FilterFundsResponseEntry__Output>) => {
		$page.url.searchParams.set('fundTwo', clickEvent.detail.id);
		history.replaceState(history.state, '', $page.url);
	};
	$: ({ comparison, colorMap } = data);

	const setResults = () => {
		return ({ result }: { result: ActionResult }) => {
			if (result.type === 'success' && result?.data?.comparison) {
				comparison = result.data.comparison;
				colorMap = result.data.colorMap;
				console.log(comparison);
			} else if (result.type === 'failure') {
				console.log(result);
				// error = result.data?.error
			}
		};
	};

</script>

<!-- Mobile -->
<form class="lg:hidden flex flex-col gap-2 p-4"
			bind:this={compareForm}
			method="POST"
			use:enhance={({formData}) => {
				 formData.set('fundOne', $page.url.searchParams.get("fundOne") || '');
				 formData.set('fundTwo', $page.url.searchParams.get("fundTwo") || '');
				 return setResults();
			}}
			action="?/compareFunds"
>
	<TertiaryContainer>
		<div class="flex flex-col gap-4">
			<FundSearchBar placeholder="Fund one"
										 hasClickListener="{true}"
										 on:fundClicked={updateFundOne}
										 value="{comparison?.fundOneName}"
			></FundSearchBar>
			<FundSearchBar placeholder="Fund two"
										 hasClickListener="{true}"
										 on:fundClicked={updateFundTwo}
										 value="{comparison?.fundTwoName}"
			></FundSearchBar>
			<QuaternaryButton on:buttonClicked={() => compareForm.requestSubmit()}>
				Compare
			</QuaternaryButton>
		</div>
	</TertiaryContainer>

	{#if (comparison && colorMap)}
		<TertiaryContainer>
			<div>{comparison.overlappingHoldingsCount}</div>
			<div class="text-xs">Number of overlapping holdings</div>
		</TertiaryContainer>
		<TertiaryContainer>
			<div>{Math.round(comparison.fundOneOverlappingCountPercentage * 100) / 100}%</div>
			<div class="text-xs">of {comparison.fundOneName} {comparison.fundOneHoldingCount} holdings also
				in {comparison.fundTwoName}</div>
		</TertiaryContainer>
		<TertiaryContainer>
			<div>{Math.round(comparison.fundTwoOverlappingCountPercentage * 100) / 100}%</div>
			<div class="text-xs">of {comparison.fundTwoName} {comparison.fundTwoHoldingCount} holdings also
				in {comparison.fundOneName}</div>
		</TertiaryContainer>
		<TertiaryContainer>
			<div>Both funds have a weighted overlap of</div>
			<div class="relative flex justify-center items-center">
				<div class="relative h-20 w-20 rounded-full opacity-50 bg-none">
					<div class="h-20 w-20 absolute left-6 rounded-full opacity-85 bg-secondary"></div>
					<div class="h-20 w-20 absolute right-6 rounded-full opacity-85 bg-quaternary"></div>
				</div>
				<div class="absolute text-xs">{Math.round(comparison.totalOverlappingPercentage * 100) / 100}%</div>
			</div>
		</TertiaryContainer>
		<TertiaryContainer on:containerClicked={() => sectorOverviewElem.growDiv()}>
			<PulldownElement
				bind:this={sectorOverviewElem}
				text="Sector overview"
				selector="sector-overview-mobile"
			>
				<div class="text-left">
					<ColoredBarChart>
						<FundColors textColor="text-primary" {colorMap} />
						{#each comparison.sectorWeightings as fws}
							<ColoredBar
								width={Math.round((fws.percentage / (comparison.sectorWeightings[0].percentage > comparison.sectorWeightings[1].percentage ? comparison.sectorWeightings[0].percentage : comparison.sectorWeightings[1].percentage)) * 100)}
								title={fws.sectorName}
								percentage={fws.percentage}
							>
								<ColoredBarEntry color="{colorMap.get(fws.fundId)?.color || ''}"
																 width={100} />
							</ColoredBar>
						{/each}
					</ColoredBarChart>

				</div>
			</PulldownElement>
		</TertiaryContainer>
		<TertiaryContainer on:containerClicked={() => overlappingHoldingsElem.growDiv()}>
			<PulldownElement
				bind:this={overlappingHoldingsElem}
				text="Top overlapping funds"
				selector="overlapping-funds-mobile"
			>
				<div class="pt-3">
					<List>
						{#each comparison.overlappingHoldings as overlappingHolding}
							<ListItem isTertiary="{false}">
								<OverlappingHoldingLineItem
									fundOneName="{comparison.fundOneName}"
									fundTwoName="{comparison.fundTwoName}"
									{overlappingHolding} />
							</ListItem>
						{/each}
					</List>
				</div>
			</PulldownElement>
		</TertiaryContainer>
		<TertiaryContainer on:containerClicked={() => fundOneNonOverlappingHoldingsElem.growDiv()}>
			<PulldownElement
				bind:this={fundOneNonOverlappingHoldingsElem}
				text="{comparison.fundOneName} Overweight Relative to {comparison.fundTwoName}"
				selector="f1-non-overlapping-funds-mobile"
				size="text-xs"
			>
				<div class="pt-3">
					<List>
						{#each comparison.fundOneNonOverlappingHoldings as nonOverlappingHolding}
							<ListItem isTertiary="{false}">
								<NonOverlappingHoldingLineItem
									{nonOverlappingHolding} />
							</ListItem>
						{/each}
					</List>
				</div>
			</PulldownElement>
		</TertiaryContainer>
		<TertiaryContainer on:containerClicked={() => fundTwoNonOverlappingHoldingsElem.growDiv()}>
			<PulldownElement
				bind:this={fundTwoNonOverlappingHoldingsElem}
				text="{comparison.fundTwoName} Overweight Relative to {comparison.fundOneName}"
				selector="f2-non-overlapping-funds-mobile"
				size="text-xs"
			>
				<div class="pt-3">
					<List>
						{#each comparison.fundTwoNonOverlappingHoldings as nonOverlappingHolding}
							<ListItem isTertiary="{false}">
								<NonOverlappingHoldingLineItem
									{nonOverlappingHolding} />
							</ListItem>
						{/each}
					</List>
				</div>
			</PulldownElement>
		</TertiaryContainer>
	{/if}
</form>

<!-- PC -->
{#if (!comparison)}
	<form class="hidden lg:visible lg:flex gap-2 p-4 justify-center items-center"
				bind:this={compareForm}
				method="POST"
				use:enhance={({formData}) => {
				 formData.set('fundOne', $page.url.searchParams.get("fundOne") || '');
				 formData.set('fundTwo', $page.url.searchParams.get("fundTwo") || '');
				 return setResults();
			}}
				action="?/compareFunds"
	>
		<div class="flex flex-col gap-2 flex-1">
			<TertiaryContainer>
				<div class="flex flex-col gap-4">
					<FundSearchBar placeholder="Fund one"
												 hasClickListener="{true}"
												 on:fundClicked={updateFundOne}
					></FundSearchBar>
					<FundSearchBar placeholder="Fund two"
												 hasClickListener="{true}"
												 on:fundClicked={updateFundTwo}
					></FundSearchBar>
					<QuaternaryButton on:buttonClicked={() => compareForm.requestSubmit()}>
						Compare
					</QuaternaryButton>
				</div>
			</TertiaryContainer>
		</div>
	</form>
{:else}

	<form class="hidden lg:visible lg:flex gap-2 p-4"
				bind:this={compareForm}
				method="POST"
				use:enhance={({formData}) => {
				 formData.set('fundOne', $page.url.searchParams.get("fundOne") || '');
				 formData.set('fundTwo', $page.url.searchParams.get("fundTwo") || '');
				 return setResults();
			}}
				action="?/compareFunds"
	>
		<div class="flex flex-col gap-2 flex-1">
			<TertiaryContainer>
				<div class="flex flex-col gap-4">
					<FundSearchBar placeholder="Fund one"
												 hasClickListener="{true}"
												 on:fundClicked={updateFundOne}
												 value="{comparison?.fundOneName}"
					></FundSearchBar>
					<FundSearchBar placeholder="Fund two"
												 hasClickListener="{true}"
												 on:fundClicked={updateFundTwo}
												 value="{comparison?.fundTwoName}"
					></FundSearchBar>
					<QuaternaryButton on:buttonClicked={() => compareForm.requestSubmit()}>
						Compare
					</QuaternaryButton>
				</div>
			</TertiaryContainer>
			{#if (comparison && colorMap)}
				<TertiaryContainer>
					<div>{comparison.overlappingHoldingsCount}</div>
					<div class="text-xs">Number of overlapping holdings</div>
				</TertiaryContainer>
				<TertiaryContainer>
					<div>{Math.round(comparison.fundOneOverlappingCountPercentage * 100) / 100}%</div>
					<div class="text-xs">of {comparison.fundOneName} {comparison.fundOneHoldingCount} holdings also
						in {comparison.fundTwoName}</div>
				</TertiaryContainer>
				<TertiaryContainer>
					<div>{Math.round(comparison.fundTwoOverlappingCountPercentage * 100) / 100}%</div>
					<div class="text-xs">of {comparison.fundTwoName} {comparison.fundTwoHoldingCount} holdings also
						in {comparison.fundOneName}</div>
				</TertiaryContainer>
				<TertiaryContainer>
					<div>Both funds have a weighted overlap of</div>
					<div class="relative flex justify-center items-center">
						<div class="relative h-20 w-20 rounded-full opacity-90 bg-none">
							<div class="h-20 w-20 absolute left-6 rounded-full opacity-85 bg-quaternary"></div>
							<div class="h-20 w-20 absolute right-6 rounded-full opacity-85 bg-[#008080]"></div>
						</div>
						<div class="absolute text-xs">{Math.round(comparison.totalOverlappingPercentage * 100) / 100}%</div>
					</div>
				</TertiaryContainer>
				<TertiaryContainer>
					<div class="mb-3">
						Top 20 overlapping funds
					</div>
					<List>
						{#each comparison.overlappingHoldings as overlappingHolding}
							<ListItem isTertiary="{false}">
								<OverlappingHoldingLineItem
									fundOneName="{comparison.fundOneName}"
									fundTwoName="{comparison.fundTwoName}"
									{overlappingHolding} />
							</ListItem>
						{/each}
					</List>
				</TertiaryContainer>
			{/if}
		</div>
		<div class="flex flex-col flex-1 gap-2">
			{#if (comparison && colorMap)}
				<ColoredBarChart>
					<FundColors textColor="text-primary" {colorMap} />
					{#each comparison.sectorWeightings as fws}
						<ColoredBar
							width={Math.round((fws.percentage / (comparison.sectorWeightings[0].percentage > comparison.sectorWeightings[1].percentage ? comparison.sectorWeightings[0].percentage : comparison.sectorWeightings[1].percentage)) * 100)}
							title={fws.sectorName}
							percentage={fws.percentage}
						>
							<ColoredBarEntry color="{colorMap.get(fws.fundId)?.color || ''}"
															 width={100} />
						</ColoredBar>
					{/each}
				</ColoredBarChart>
				<TertiaryContainer on:containerClicked={() => fundOneNonOverlappingHoldingsElem.growDiv()}>
					<div class="mb-3"><b>{comparison.fundOneName}</b> Overweight Relative to <b>{comparison.fundTwoName}</b></div>
					<List>
						{#each comparison.fundOneNonOverlappingHoldings as nonOverlappingHolding}
							<ListItem isTertiary="{false}">
								<NonOverlappingHoldingLineItem
									{nonOverlappingHolding} />
							</ListItem>
						{/each}
					</List>
				</TertiaryContainer>
				<TertiaryContainer on:containerClicked={() => fundTwoNonOverlappingHoldingsElem.growDiv()}>
					<div class="mb-3"><b>{comparison.fundTwoName}</b> Overweight Relative to <b>{comparison.fundOneName}</b></div>
					<List>
						{#each comparison.fundTwoNonOverlappingHoldings as nonOverlappingHolding}
							<ListItem isTertiary="{false}">
								<NonOverlappingHoldingLineItem
									{nonOverlappingHolding} />
							</ListItem>
						{/each}
					</List>
				</TertiaryContainer>
			{/if}
		</div>
	</form>
{/if}
