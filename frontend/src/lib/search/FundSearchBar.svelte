<script lang="ts">
	import Result from '$lib/search/Result.svelte';
	import SearchBar from '$lib/search/SearchBar.svelte';
	import { debounce, type themeType } from '$lib/utils.ts';
	import type { FilterFundsResponseEntry__Output } from '$lib/proto/proto/FilterFundsResponseEntry.ts';
	import { createEventDispatcher } from 'svelte';

	let funds: FilterFundsResponseEntry__Output[] = [];

	let searchForm: HTMLFormElement;
	export let value: string | undefined = '';
	export let theme: themeType | undefined = undefined;
	export let placeholder: string | undefined = '';
	export let hasClickListener: boolean = false;
	const search = debounce(async function() {
		const resp = await fetch('/api/search-funds', {
			method: 'POST',
			body: JSON.stringify({
				value
			})
		});
		if (resp.ok) {
			funds = (await resp.json()) as FilterFundsResponseEntry__Output[];
		} else if (!resp.ok) {
			// error = result.data?.error
		}
	}, 200);

	const dispatch = createEventDispatcher();
	const updateFundValue = (clickEvent: CustomEvent<FilterFundsResponseEntry__Output>) => {
		value = clickEvent.detail.name
		dispatch("fundClicked", clickEvent.detail)
	}
</script>

<div class="relative w-full">
	<SearchBar theme="{theme}" placeholder="{placeholder || 'Search for funds'}" on:inputChanged={search} bind:value>
		<ul class="absolute top-12 left-0 right-0 w-full">
			{#each funds || [] as fund}
				{#if (hasClickListener)}
					<Result href="#" on:fundClicked={updateFundValue} {fund} />
				{:else}
					<Result href="/fund/{fund.id}" on:click={() => console.log(fund.id)} {fund} />
				{/if}
			{/each}
		</ul>
	</SearchBar>
</div>
