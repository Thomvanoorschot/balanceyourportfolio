<script lang="ts">

    import type {SearchFundsEntry} from "$lib/proto/proto/SearchFundsEntry";

    export let fund: SearchFundsEntry;
    export let href: string;
</script>
<li>
    <a href="{href}" class="cursor-pointer bg-gray-200 rounded-md flex items-center m-1 p-3 transition duration-500 ease-in-out transform hover:-translate-y-1 hover:shadow-lg">
        <div>
            <div class="font-medium text-sm">{fund.name}</div>
            <div class="flex">
                {#each fund.tickers || [] as ticker}
                    <div class="font-light text-xs pr-2">{ticker}</div>
                {/each}
            </div>
        </div>
    </a>
</li>
