<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import type { LabelValue } from '$lib/utils.ts';
	import { clickOutside } from '$lib/custom-svelte-typings';

	export let value: any;
	export let entries: LabelValue[] = [];

	let isActive: boolean = false;
	let label: string | undefined = '';

	onMount(() => {
		document.addEventListener('keydown', closeOnEscape);
		return () => document.removeEventListener('keydown', closeOnEscape);
	});
	const closeOnEscape = (e: KeyboardEvent) => {
		if (!isActive) {
			return;
		}
		if (e.key === 'Escape') {
			isActive = false;
			e.stopImmediatePropagation();
			e.preventDefault();
			return;
		}
	};

	const dispatch = createEventDispatcher();
</script>

<div class="relative w-full">
	<div>
		<button
			on:click={() => {
				isActive = !isActive;
			}}
			type="button"
			class="flex items-center outline-none bg-secondary w-full text-center rounded-xl border-2 border-primary {isActive
				? 'border-quaternary'
				: ''} pt-2 pb-2"
		>
			<div class="flex-1 {value ? '' : 'text-opacity-10'}">
				{label || 'Portfolio'}
			</div>
			<svg class="h-5 w-5 fill-tertiary mr-2" viewBox="0 0 20 20" aria-hidden="true">
				<path
					fill-rule="evenodd"
					d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
					clip-rule="evenodd"
				/>
			</svg>
		</button>
	</div>
	{#if isActive}
		<div class="absolute w-full" use:clickOutside on:click_outside={() => (isActive = !isActive)}>
			{#each entries as entry}
				<div
					on:click={() => {
						value = entry.value;
						label = entry.label;
						isActive = false;
						dispatch('optionChanged');
					}}
					aria-hidden="true"
					class="cursor-pointer bg-primary border-2 border-tertiary rounded-xl flex items-center
                             justify-center m-0.5 p-3 hover:drop-shadow-md hover:bg-quaternary"
				>
					{entry.label}
				</div>
			{/each}
		</div>
	{/if}
</div>
