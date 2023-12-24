<script lang="ts">

    import type {SearchFundsEntry} from "$lib/proto/proto/SearchFundsEntry";

    export let fund: SearchFundsEntry;
    export let href: string;
</script>
<li>
    <a href="{href}" class="cursor-pointer bg-primary border-2 border-tertiary rounded-xl flex items-center m-0.5 p-3 hover:drop-shadow-md hover:bg-quaternary">
        <div>
            <div class="font-medium text-sm text-tertiary">{fund.name}</div>
            <div class="flex">
                {#each fund.tickers || [] as ticker}
                    <div class="font-light text-xs pr-2 text-tertiary">{ticker}</div>
                {/each}
            </div>
        </div>
    </a>
</li>
