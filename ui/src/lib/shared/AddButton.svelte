<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import AddIcon from '$lib/icons/AddIcon.svelte';

	const dispatch = createEventDispatcher();

	function buttonClicked() {
		dispatch('buttonClicked');
	}
</script>

<button
	class="px-2 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-primary border border-transparent rounded-full active:bg-primary hover:bg-quaternary focus:outline-none focus:shadow-outline-violet"
	type="button"
	on:click={buttonClicked}
>
	<AddIcon />
</button>
