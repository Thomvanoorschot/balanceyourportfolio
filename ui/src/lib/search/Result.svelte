<script lang="ts">

    import type {Fund} from "$lib/fund";

    export let fund: Fund;
    export let href: string;
</script>
<li class="">
    <a href="{href}" class="cursor-pointer bg-gray-200 rounded-md flex flex-1 items-center p-3 m-1 transition duration-500 ease-in-out transform hover:-translate-y-1 hover:shadow-lg">
        <div class="flex-1 pl-1 mr-16">
            <div class="font-medium text-sm">{fund.name}</div>
            <div class="flex">
                {#each fund.tickers as ticker}
                    <div class="font-light text-xs pr-2">{ticker}</div>
                {/each}
            </div>
        </div>
    </a>
</li>
