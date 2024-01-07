<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import { clickOutside } from '$lib/custom-svelte-typings';
	import type { themeType } from '$lib/utils.ts';

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

	let bgColor: themeType;
	let borderColor: themeType;
	let focusBorderColor: themeType;
	let focusBgColor: themeType;
	let placeholderColor: themeType;
	let textColor: themeType;

	onMount(() => {
		if (theme === 'primary') {
			borderColor = 'tertiary';
			focusBorderColor = 'quaternary';
			bgColor = 'primary';
			focusBgColor = 'secondary';
			textColor = 'tertiary';
			placeholderColor = 'tertiary';
		} else if (theme === 'secondary') {
			borderColor = 'primary';
			focusBorderColor = 'quaternary';
			bgColor = 'secondary';
			focusBgColor = 'secondary';
			textColor = 'quaternary';
			placeholderColor = 'quaternary';
		}
	});
</script>

<div
	use:clickOutside
	on:click_outside={() => (hasFocus = false)}
	class="flex justify-center w-full"
>
	<!--	<div class="w-full relative">-->
	<!--		<div class="absolute inset-y-0 flex items-center pl-2">-->
	<!--			<SearchIcon inPrimary="{inPrimary}"></SearchIcon>-->
	<!--		</div>-->
	<input
		class="w-full pt-2 pb-2 pl-8 pr-2 text-sm placeholder-gray-600 border-2 rounded-xl focus:outline-none
                bg-{bgColor} border-{borderColor} focus:border-{focusBorderColor} focus:bg-{focusBgColor} placeholder-{placeholderColor} text-{textColor}"
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

	<!--	</div>-->
	{#if hasFocus}
		<slot />
	{/if}
</div>
