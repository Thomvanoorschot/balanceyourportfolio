<script lang="ts">
    export let colorMap: Map<string, string>;
</script>
<div class="flex gap-2 flex-wrap">
    {#each [...colorMap] as [key, value]}
        <div class="flex items-center gap-2">
            <div class="rounded-full h-3 w-3" style="background-color: {value}"></div>
            <div class="text-xs">{key}</div>
        </div>
    {/each}
</div>