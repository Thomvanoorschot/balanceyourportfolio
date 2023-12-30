<script lang="ts">
    import type {PageData} from './$types';
    import Portfolio from "$lib/portfolios/Portfolio.svelte";
    import type {PortfoliosResponse__Output} from "$lib/proto/proto/PortfoliosResponse";

    export let data: PageData;
    let portfolios: PortfoliosResponse__Output | undefined
    let error: string | undefined
    let isAuthenticated: boolean = false
    $:{
        portfolios = data?.portfolios
        error = ""
        isAuthenticated = data.isAuthenticated
    }
</script>
{#if (!error && portfolios)}
    <div id="page" class="flex flex-grow flex-col w-full items-center justify-start">
        {#each portfolios.entries as portfolio}
            <Portfolio isAuthenticated="{isAuthenticated}" portfolio="{portfolio}"></Portfolio>
        {/each}
        <Portfolio
                isAuthenticated="{isAuthenticated}"
                portfolio="{{entries:[{name:'',id:'', fundId: '', amount: 0}], name: `Portfolio ${portfolios.entries.length + 1}`, id: ''}}"
        ></Portfolio>
    </div>
{:else}
    <h1>{error}</h1>
{/if}
