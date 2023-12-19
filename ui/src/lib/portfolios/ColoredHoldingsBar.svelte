<script lang="ts">
    import type {Holding} from "$lib/holding.ts";

    export let holding: Holding
    export let colorMap: Map<string, {
        fundName: string,
        color: string
    }> | undefined = undefined
</script>
{#if (holding.funds?.length > 0)}
    <div class="flex w-full h-2 mt-3">
        {#each holding.funds as fund}
            <div class="h-full" style=" background-color: {colorMap?.get(fund.fundId)?.color}; width: { `${Math.round(fund.ratiodPercentage / holding.percentage * 100)}%`}"></div>
        {/each}
    </div>
{/if}