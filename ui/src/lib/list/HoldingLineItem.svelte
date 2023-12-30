<script lang="ts">
    import type {Holding} from "$lib/holding.ts";
    import {onMount} from "svelte";

    export let holding: Holding
    let imgUrl: string | undefined = ""
    onMount(async () => {
            const imgResp = await fetch(`/company-logos/${holding.ticker}.png`, {
                method: 'HEAD',
            })
            if (imgResp.ok) {
                imgUrl = imgResp.url
            }
        }
    )
</script>

<div class="flex flex-1">
    <div class="flex flex-col rounded-xl w-10 h-10 bg-gray-300 justify-center items-center mr-4">
        {#if (imgUrl)}
            <img class="scale-75" src="{imgUrl}" alt="">
        {:else}
            <img src="/company-logos/UNKNOWN.png" alt="">
        {/if}
    </div>
    <div class="flex-1 pl-1 mr-16">
        <div class="font-medium">{holding.ticker}</div>
        <div class="text-sm">{holding.name}</div>
    </div>
    <div class="text-xs">{Math.round(holding.percentage * 100) / 100}%</div>
</div>
