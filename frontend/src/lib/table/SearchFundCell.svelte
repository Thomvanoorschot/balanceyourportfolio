<script lang="ts">
	import { clickOutside } from '$lib/custom-svelte-typings';
	import Result from '$lib/search/Result.svelte';
	import { debounce } from '$lib/utils.ts';
	import { onMount, tick } from 'svelte';
	import { enhance } from '$app/forms';
	import type { ActionResult } from '@sveltejs/kit';
	import type { FilterFundsResponseEntry__Output } from '$lib/proto/proto/FilterFundsResponseEntry.ts';

	export let label: string;
	export let value: string;
	export let disabled: boolean = true;

	let inputElement: HTMLInputElement;
	$: {
		if (!disabled) {
			tick().then(() => inputElement.focus());
		}
	}
	let searchForm: HTMLFormElement;
	let showDropdown = false;
	let funds: FilterFundsResponseEntry__Output[] = [];

	onMount(() => {
		document.addEventListener('keydown', closeOnEscape);
		return () => document.removeEventListener('keydown', closeOnEscape);
	});
	const closeOnEscape = (e: KeyboardEvent) => {
		if (e.key === 'Escape') showDropdown = false;
	};
	const handleClickOutside = () => {
		showDropdown = false;
	};

	const search = debounce(async function () {
		if (label === '') {
			funds = [];
			return;
		}
		searchForm.requestSubmit();
	}, 200);
	const handleFundClicked = (f: FilterFundsResponseEntry__Output) => {
		showDropdown = false;
		label = f.name;
		value = f.id;
	};
	const updateFunds = () => {
		return ({ result }: { result: ActionResult }) => {
			if (result.type === 'success' && result?.data?.funds) {
				funds = result?.data?.funds;
			} else if (result.type === 'failure') {
				// error = result.data?.error
			}
		};
	};
</script>

<td class="px-4 py-3 text-sm relative">
	<form
		method="POST"
		action="?/searchFunds"
		bind:this={searchForm}
		use:enhance={({ formData }) => {
			formData.set('searchTerm', label);
			return updateFunds();
		}}
		class="w-full relative"
	>
		<input bind:value class="hidden" type="text" />
		<input
			{disabled}
			on:input={() => {
				if(!showDropdown){
					showDropdown = true
				}
				search()
			}}
			on:focus={() => (showDropdown = true)}
			on:click={() => (showDropdown = true)}
			bind:value={label}
			bind:this={inputElement}
			class="overflow-ellipsis w-full rounded-lg outline-none border-2 border-white p-3 bg-primary {disabled
				? ''
				: 'focus:border-2 focus:border-quaternary focus:outline-none bg-secondary'}"
			type="text"
		/>
	</form>
		{#if showDropdown}
			<ul
				use:clickOutside
				on:click_outside={handleClickOutside}
				class="absolute top-15 left-0 right-0 w-screen z-10 lg:w-full"
			>
				{#each funds as fund}
					<div
						tabindex="0"
						aria-label=""
						role="button"
						on:click={() => handleFundClicked(fund)}
						on:keydown={() => handleFundClicked(fund)}
					>
						<Result href="#" {fund} />
					</div>
				{/each}
			</ul>
		{/if}
</td>
