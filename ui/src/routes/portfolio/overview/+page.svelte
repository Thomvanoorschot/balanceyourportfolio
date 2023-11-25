<script lang="ts">
    import type {PageData} from './$types';
    import Portfolio from "$lib/portfolios/Portfolio.svelte";
    import {createPortfoliosStore} from "$lib/stores/portfolio-store";
    import {setContext} from "svelte";

    export let data: PageData;
    const {portfolios} = data;

    const portfoliosStore = createPortfoliosStore(portfolios)
    setContext("portfoliosStore", portfoliosStore)
</script>
<div id="portfolios" class="flex flex-grow flex-col w-full items-center justify-start">
    {#each $portfoliosStore.portfolios as portfolio}
        <Portfolio portfolio="{portfolio}"></Portfolio>
    {/each}
</div>