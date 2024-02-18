<script lang="ts">
	import { browser } from '$app/environment';
	import type { OverlappingHolding__Output } from '$lib/proto/proto/OverlappingHolding.ts';
	import type { NonOverlappingHolding__Output } from '$lib/proto/proto/NonOverlappingHolding.ts';

	export let nonOverlappingHolding: NonOverlappingHolding__Output;
	const handleError = (ev: any) => ev.target.src = '/company-logos/UNKNOWN.png';
</script>

<div class="flex flex-col gap-2">
	<div class="flex items-center">
		<div class="flex flex-col rounded-xl w-10 h-10 bg-gray-300 justify-center items-center mr-4">
			{#if (browser)}
				<img class="scale-75" src="/company-logos/{nonOverlappingHolding.holdingTicker}.png" alt=""
						 on:error={handleError} />
			{/if}
		</div>
		<div class="flex-1 pl-1 mr-16">
			<div class="font-medium">{nonOverlappingHolding.holdingTicker}</div>
			<div class="text-sm">{nonOverlappingHolding.holdingName}</div>
		</div>
		<div class="text-xs">{Math.round(nonOverlappingHolding.nonOverlappingPercentage * 100) / 100}%</div>
	</div>
</div>
