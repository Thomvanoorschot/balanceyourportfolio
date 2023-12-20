<script lang="ts">
    export let colorMap: Map<string, {fundName: string, color: string}>
</script>
<div class="flex gap-2 flex-wrap">
    {#each [...colorMap] as [key, value]}
        <div class="flex items-center">
            <div class="aspect-square rounded-full h-3 w-3" style="background-color: {value.color}"></div>
            <div class="pl-2 text-xs">{value.fundName}</div>
        </div>
    {/each}
</div>