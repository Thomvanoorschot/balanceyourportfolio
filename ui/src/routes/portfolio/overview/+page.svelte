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
        <form class="flex relative flex-col m-20 w-[50vw] rounded-lg shadow-lg">
            <Table headers="{['Ticker or name', 'Amount', 'Edit', 'Delete']}"></Table>

        </form>
        {#each portfolios.entries as portfolio}
            <Portfolio portfolio="{portfolio}"></Portfolio>
        {/each}
    </div>
{:else}
    <h1>{error}</h1>
{/if}
