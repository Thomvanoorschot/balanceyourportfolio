<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let disabled: boolean | undefined = false;

	const dispatch = createEventDispatcher();
</script>

<button
	class="w-full px-10 py-4 font-medium text-primary transition-colors duration-150
        bg-quaternary border border-transparent rounded-lg focus:outline-none
        {disabled ? 'bg-opacity-80' : ' hover:opacity-90'}
        "
	on:click={() => dispatch('buttonClicked')}
	type="button"
	{disabled}
>
	<slot />
</button>
