<script lang="ts">
	import type { Holding } from '$lib/holding.ts';
	import { browser } from '$app/environment';

	export let holding: Holding;
	const handleError = (ev: any) => ev.target.src = '/company-logos/UNKNOWN.png';
</script>

<div class="flex flex-1">
	<div class="flex flex-col rounded-xl w-10 h-10 bg-gray-300 justify-center items-center mr-4">
		{#if (browser)}
			{#if holding.ticker.endsWith('- Bond')}
				<img
					class="scale-75"
					src="/company-logos/{holding.ticker.replace(' - Bond', '')}.png"
					alt=""
					on:error={handleError}
				/>
			{:else}
				<img class="scale-75" src="/company-logos/{holding.ticker}.png" alt="" on:error={handleError} />
			{/if}
		{/if}
	</div>
	<div class="flex-1 pl-1 mr-16">
		<div class="font-medium">{holding.ticker}</div>
		<div class="text-sm">{holding.name}</div>
	</div>
	<div class="text-xs">{Math.round(holding.percentage * 100) / 100}%</div>
</div>
