<script lang="ts">

    import type {SearchFundsEntry} from "$lib/proto/proto/SearchFundsEntry";

    export let fund: SearchFundsEntry;
    export let href: string;
</script>
<li>
    <a href="{href}" class="cursor-pointer bg-violet-100 rounded-md flex items-center m-0.5 p-3 hover:drop-shadow-md hover:bg-violet-200">
        <div>
            <div class="font-medium text-sm">{fund.name}</div>
            <div class="flex">
                {#each fund.tickers || [] as ticker}
                    <div class="font-light text-xs pr-2 text-violet-900 ">{ticker}</div>
                {/each}
            </div>
        </div>
    </a>
</li>
