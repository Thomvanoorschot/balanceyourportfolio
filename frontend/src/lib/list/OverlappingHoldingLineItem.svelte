<script lang="ts">
	import { browser } from '$app/environment';
	import type { OverlappingHolding__Output } from '$lib/proto/proto/OverlappingHolding.ts';

	export let fundOneName: string;
	export let fundTwoName: string;
	export let overlappingHolding: OverlappingHolding__Output;
	const handleError = (ev: any) => ev.target.src = '/company-logos/UNKNOWN.png';
</script>

<div class="flex flex-col gap-2">
	<div class="flex items-center">
		<div class="flex flex-col rounded-xl w-10 h-10 bg-gray-300 justify-center items-center mr-4">
			{#if (browser)}
				<img class="scale-75" src="/company-logos/{overlappingHolding.holdingTicker}.png" alt=""
						 on:error={handleError} />
			{/if}
		</div>
		<div class="flex-1 pl-1 mr-16">
			<div class="font-medium">{overlappingHolding.holdingTicker}</div>
			<div class="text-sm">{overlappingHolding.holdingName}</div>
		</div>
		<div class="text-xs">{Math.round(overlappingHolding.overlappingPercentage * 100) / 100}%</div>
	</div>
	<div class="flex justify-between items-center text-start text-xs">
		<div>
			Weight in <b>{fundOneName}</b> is
		</div>
		<div class="ml-5">
			{Math.round(overlappingHolding.fundOnePercentage * 100) / 100}%
		</div>
	</div>
	<div class="flex justify-between items-center text-start text-xs">
		<div>
			Weight in <b>{fundTwoName}</b> is
		</div>
		<div class="ml-5">
			{Math.round(overlappingHolding.fundTwoPercentage * 100) / 100}%
		</div>
	</div>
</div>
