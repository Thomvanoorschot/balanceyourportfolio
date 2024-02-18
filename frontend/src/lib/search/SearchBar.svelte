<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import { clickOutside } from '$lib/custom-svelte-typings';
	import type { themeType } from '$lib/utils.ts';
	import SearchIcon from '$lib/icons/SearchIcon.svelte';
	import CrossIcon from '$lib/icons/CrossIcon.svelte';

	export let placeholder: string;
	export let value: string | undefined = '';
	export let theme: themeType = 'primary';

	const dispatch = createEventDispatcher();

	function inputChanged() {
		if (!hasFocus) {
			hasFocus = true;
		}
		dispatch('inputChanged');
	}

	let hasFocus: boolean;
	onMount(() => {
		document.addEventListener('keydown', closeOnEscape);
		return () => document.removeEventListener('keydown', closeOnEscape);
	});
	const closeOnEscape = (e: KeyboardEvent) => {
		if (e.key === 'Escape' && hasFocus) hasFocus = false;
	};

	let inputElem: HTMLInputElement;
	onMount(() => {
		if (theme === 'primary') {
			inputElem.classList.add(
				'bg-primary',
				'border-tertiary',
				'focus:border-quaternary',
				'focus:bg-secondary',
				'placeholder-tertiary',
				'text-tertiary'
			);
		} else if (theme === 'secondary') {
			inputElem.classList.add(
				'bg-secondary',
				'border-primary',
				'focus:border-quaternary',
				'focus:bg-secondary',
				'placeholder-tertiary',
				'text-tertiary'
			);
		}
	});
</script>

<div
	use:clickOutside
	on:click_outside={() => {
		setTimeout(() => hasFocus = false, 500)
	}}
	class="flex justify-center w-full relative"
>
	<div class="absolute top-3 left-0 flex items-center pl-2">
		<SearchIcon inPrimary="{true}"></SearchIcon>
	</div>
	<input
		bind:this={inputElem}
		class="w-full pt-2 pb-2 pl-8 pr-2 text-sm placeholder-gray-600 border-2 rounded-xl focus:outline-none "
		type="text"
		{placeholder}
		aria-label="Search"
		bind:value
		on:input={inputChanged}
		on:click={(e) => {
			hasFocus = true;
			e.stopImmediatePropagation();
			e.preventDefault();
		}}
		on:blur={() =>
			setTimeout(() => {
				hasFocus = false;
			}, 200)}
	/>
	{#if (value)}
		<div aria-hidden="true" class="absolute top-3.5 right-3 flex items-center pl-2"
				 on:click={(e) => {
					 hasFocus = false;
					 value = ''
					 inputChanged()
					 e.stopImmediatePropagation();
			     e.preventDefault();
				  }}
		>
			<CrossIcon></CrossIcon>
		</div>
	{/if}
</div>
{#if hasFocus}
	<slot />
{/if}
