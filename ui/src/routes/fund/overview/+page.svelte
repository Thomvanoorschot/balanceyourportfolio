<script lang="ts">
	import type { PageData } from './$types';
	import SearchBar from '$lib/search/SearchBar.svelte';
	import CheckButtonList from '$lib/filters/CheckButtonList.svelte';
	import List from '$lib/list/List.svelte';
	import DetailMenu from '$lib/menu/DetailMenu.svelte';
	import ListItem from '$lib/list/ListItem.svelte';
	import FundLineItem from '$lib/list/FundLineItem.svelte';
	import { enhance } from '$app/forms';
	import type { ActionResult } from '@sveltejs/kit';
	import { debounce } from '$lib/utils.ts';

	export let data: PageData;
	let error: string | undefined;
	let searchTerm: string = '';
	let providers: string[] = ['Vanguard'];
	let filterForm: HTMLFormElement;
	let resetSearch: boolean;

	$: ({ funds } = data);

	const filterFunds = debounce(async function () {
		resetSearch = true;
		filterForm.requestSubmit();
	}, 200);
	const updateNextPage = () => {
		return ({ result }: { result: ActionResult }) => {
			if (result.type === 'success' && result?.data?.funds && funds) {
				if (resetSearch) {
					funds = [...result?.data?.funds];
					resetSearch = false;
					return;
				}
				funds = [...funds, ...result?.data?.funds];
			} else if (result.type === 'failure') {
				error = result.data?.error;
			}
		};
	};
	const setFilterForm = (formData: FormData) => {
		resetSearch
			? formData.set('fundsLength', '0')
			: formData.set('fundsLength', funds!.length.toString());
		formData.set('searchTerm', searchTerm);
	};
</script>

{#if !error && funds}
	<div class="flex flex-col lg:flex-row w-full gap-5 p-2">
		<DetailMenu>
			<SearchBar
				placeholder="Fund name or ticker"
				on:inputChanged={filterFunds}
				bind:value={searchTerm}
			/>
			<CheckButtonList title="Providers" list={providers} on:checkButtonClicked={() => {}} />
		</DetailMenu>
		<div class="flex-grow">
			<form
				method="POST"
				action="?/filterFunds"
				bind:this={filterForm}
				use:enhance={({ formData }) => {
					setFilterForm(formData);
					return updateNextPage();
				}}
			>
				<List on:endOfPageReached={() => filterForm.requestSubmit()}>
					{#each funds as fund}
						<a href="/fund/{fund.id}" class="hover:opacity-90">
							<ListItem>
								<FundLineItem {fund} />
							</ListItem>
						</a>
					{/each}
				</List>
			</form>
		</div>
	</div>
{/if}
