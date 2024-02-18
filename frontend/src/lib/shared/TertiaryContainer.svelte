<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
</script>

<div aria-hidden="true" class="p-2 text-center shadow-md rounded-xl bg-tertiary text-primary"
		 on:click={() =>dispatch("containerClicked")}>
	<slot></slot>
</div>