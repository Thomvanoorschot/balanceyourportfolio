<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import type { themeType } from '$lib/utils.ts';

	export let theme: themeType = 'primary';
	export let index: number;
	export let text: string;
	export let checked: boolean | undefined = false;
	export let identifier: string

	let labelElem: HTMLLabelElement
	onMount(() => {
		if (theme === 'primary') {
			labelElem.classList.add(
				'border-tertiary',
				'bg-primary',
				'text-tertiary',
				'peer-checked:bg-secondary',
				'peer-checked:border-quaternary',
				'lg:peer-hover:bg-quaternary',
				'lg:peer-hover:border-quaternary',
				'peer-checked:text-tertiary',
			)
		} else if (theme === 'tertiary') {
			labelElem.classList.add(
				'border-tertiary',
				'bg-tertiary',
				'text-primary',
				'peer-checked:bg-primary',
				'peer-checked:border-quaternary',
				'peer-hover:bg-quaternary',
				'peer-hover:border-quaternary',
				'peer-checked:text-tertiary',
			)
		}
	});
	const dispatch = createEventDispatcher();
</script>

<div class="flex">
	<input
		type="checkbox"
		id="check-button-{index}-{identifier}"
		class="peer hidden"
		bind:checked
		on:click={() => dispatch('checkButtonClicked', text)}
	/>
	<label
		bind:this={labelElem}
		for="check-button-{index}-{identifier}"
		class="select-none cursor-pointer rounded-lg border-2
            py-2 px-2.5 font-small text-xs  transition-colors duration-200 ease-in-out"
	>
		{text}
	</label>
</div>
