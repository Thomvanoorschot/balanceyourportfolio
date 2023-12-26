<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import {onNavigate} from "$app/navigation";

    export let title: string
    export let percentage: number

    let count = 0;

    onNavigate(() => {
        count++
    })
    const dispatch = createEventDispatcher()
</script>
{#key count}
    <div class="flex justify-start items-center text-primary">
        <div class="w-1/4 text-xs">{title}</div>
        <div class="text-xs w-16">{Math.round(percentage * 100) / 100}%</div>
        <div class="w-full flex">
            <div aria-hidden="true" class="flex resizableElement hover:opacity-80 transition-opacity"
                 on:click={() => dispatch("onBarClicked", title)}>
                <slot></slot>
            </div>
        </div>
    </div>
{/key}

<style>
    @keyframes changeWidth {
        0% {
            width: 0
        }
        100% {
            width: 100%;
        }
    }

    .resizableElement {
        will-change: transform;
        animation: changeWidth 1s ease-in-out forwards;
    }
</style>