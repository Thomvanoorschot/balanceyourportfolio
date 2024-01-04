<script lang="ts">
	import type { FilterFundsResponseEntry__Output } from '$lib/proto/proto/FilterFundsResponseEntry.ts';
	import { currencySignMap, formatNumber } from '$lib/utils.ts';

	export let fund: FilterFundsResponseEntry__Output;
</script>

<div class="flex items-center gap-2" aria-hidden="true">
	<div class="flex flex-col rounded-xl w-10 h-10 bg-gray-300">
		<img class="scale-75" src="/provider-logos/vanguard.png" alt="" />
	</div>
	<div class="flex-1">
		<div class="font-medium text-sm lg:text-lg">{fund.name}</div>
		<div class="flex-wrap hidden lg:flex lg:visible">
			{#each fund.tickers || [] as ticker}
				<div class="pr-3 text-xs lg:text-sm">{ticker}</div>
			{/each}
		</div>
	</div>
	<div>{currencySignMap[fund.currency] || fund.currency} {formatNumber(fund.marketCap)}</div>
</div>
