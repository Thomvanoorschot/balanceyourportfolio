<script lang="ts">
	import type { FilterFundsResponseEntry__Output } from '$lib/proto/proto/FilterFundsResponseEntry.ts';
	import { createEventDispatcher } from 'svelte';

	export let fund: FilterFundsResponseEntry__Output;
	export let href: string;
	const dispatch = createEventDispatcher();
</script>

<li>
	<a
		on:click={() => dispatch("fundClicked", fund)}
		{href}
		class="z-20 relative bg-primary border-2 border-tertiary rounded-xl flex items-center m-0.5 p-2 hover:drop-shadow-md hover:bg-quaternary"
	>
		<div>
			<div class="font-medium text-sm text-tertiary">{fund.name}</div>
			<div class="hidden lg:flex lg:visible">
				{#each fund.tickers || [] as ticker}
					<div class="font-light text-xs pr-2 text-tertiary">{ticker}</div>
				{/each}
			</div>
		</div>
	</a>
</li>
