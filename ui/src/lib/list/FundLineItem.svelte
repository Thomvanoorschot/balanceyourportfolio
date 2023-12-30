<script lang="ts">
    import type {FilterFundsResponseEntry__Output} from "$lib/proto/proto/FilterFundsResponseEntry.ts";
    import {currencySignMap, formatNumber} from "$lib/utils.ts";

    export let fund: FilterFundsResponseEntry__Output

</script>

<div class="flex flex-1" aria-hidden="true">
    <div class="flex flex-col rounded-xl w-10 h-10 bg-gray-300 justify-center items-center mr-4">
        <img class="scale-75" src="/provider-logos/vanguard.png" alt="">
    </div>
    <div class="flex-1 pl-1 mr-16">
        <div class="font-medium">{fund.name}</div>
        <div class="flex gap-5">
            {#each fund.tickers || [] as ticker}
                <div class="text-sm">{ticker}</div>
            {/each}
        </div>
    </div>
    <div class="w-20">{currencySignMap[fund.currency] || fund.currency} {formatNumber(fund.marketCap)}</div>
</div>
