<script lang="ts">
	import { createEventDispatcher, onDestroy, onMount } from 'svelte';

	let observer: IntersectionObserver;
	let root: HTMLElement;

	const dispatch = createEventDispatcher();

	function endOfPageReached() {
		dispatch('endOfPageReached');
	}

	onMount(() => {
		observer = new IntersectionObserver(async (entries) => {
			if (entries[0].isIntersecting) {
				endOfPageReached();
			}
		});
		observer.observe(root);
	});
	onDestroy(() => {
		if (observer) {
			observer.disconnect();
		}
	});
</script>

<div class="gap-2 flex flex-col">
	<slot />
	<div bind:this={root} />
</div>
