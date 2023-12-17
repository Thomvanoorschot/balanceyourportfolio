<script lang="ts">
    import type {PageData} from './$types';
    import Portfolio from "$lib/portfolios/Portfolio.svelte";
    import type {PortfoliosResponse__Output} from "$lib/proto/proto/PortfoliosResponse";
    import Table from "$lib/table/Table.svelte";

    export let data: PageData;
    let portfolios: PortfoliosResponse__Output | undefined
    let error: string | undefined
    $:{
        portfolios = data?.portfolios
        error = ""
    }
    // const portfoliosStore = createPortfoliosStore(portfolios)
    // setContext("portfoliosStore", portfoliosStore)
</script>
{#if (!error && portfolios)}
    <div id="portfolios" class="flex flex-grow flex-col w-full items-center justify-start">
        {#each portfolios.entries as portfolio}
            <Portfolio portfolio="{portfolio}"></Portfolio>
        {/each}
        <Portfolio portfolio="{{entries:[], name: '', id: ''}}"></Portfolio>
    </div>
{:else}
    <h1>{error}</h1>
{/if}
