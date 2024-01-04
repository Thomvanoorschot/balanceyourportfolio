<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let index: number;
	export let text: string;
	export let checked: boolean | undefined = false;

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
		class="select-none cursor-pointer rounded-lg border-2 border-tertiary bg-tertiary
            py-2 px-2.5 font-small text-xs text-primary transition-colors duration-200
            ease-in-out peer-checked:border-quaternary peer-hover:bg-quaternary
             peer-hover:border-quaternary peer-checked:bg-primary peer-checked:text-tertiary"
	>
		{text}
	</label>
</div>
