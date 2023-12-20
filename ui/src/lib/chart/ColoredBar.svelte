<script lang="ts">
    export let title: string
    export let percentage: number
</script>
<div class="flex flex-1 justify-start items-center">
    <div class="w-1/4 text-xs">{title}</div>
    <div class="text-xs w-16">{Math.round(percentage * 100) / 100}%</div>
    <div class="flex w-full">
        <slot></slot>
    </div>
</div>