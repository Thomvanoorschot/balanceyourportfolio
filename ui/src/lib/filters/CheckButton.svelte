<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import type { themeType } from '$lib/utils.ts';

	export let theme: themeType = 'primary';
	export let index: number;
	export let text: string;
	export let checked: boolean | undefined = false;

	let borderColor: themeType;
	let bgColor: themeType;
	let hoverColor: themeType;
	let borderCheckedColor: themeType;
	let textColor: themeType;
	let textCheckedColor: themeType;
	let bgCheckedColor: themeType;

	onMount(() => {
		if (theme === 'primary') {
			borderColor = 'tertiary';
			bgColor = 'primary';
			bgCheckedColor = 'secondary';
			hoverColor = 'quaternary';
			borderCheckedColor = 'quaternary';
			textColor = 'tertiary';
			textCheckedColor = 'tertiary';
		} else if (theme === 'tertiary') {
			borderColor = 'tertiary';
			bgColor = 'tertiary';
			bgCheckedColor = 'primary';
			hoverColor = 'quaternary';
			borderCheckedColor = 'quaternary';
			textColor = 'primary';
			textCheckedColor = 'tertiary';
		}
	});
	const dispatch = createEventDispatcher();
</script>

<div class="flex">
	<input
		type="checkbox"
		id="check-button-{index}"
		class="peer hidden"
		bind:checked
		on:click={() => dispatch('checkButtonClicked', text)}
	/>
	<label
		for="check-button-{index}"
		class="select-none cursor-pointer rounded-lg border-2
            py-2 px-2.5 font-small text-xs  transition-colors duration-200 ease-in-out
            border-{borderColor} bg-{bgColor} text-{textColor}  peer-checked:bg-{bgCheckedColor} peer-checked:border-{borderCheckedColor} peer-hover:bg-{hoverColor}
             peer-hover:border-{hoverColor}  peer-checked:text-{textCheckedColor}"
	>
		{text}
	</label>
</div>
