<script lang="ts">
    import CheckButton from "$lib/filters/CheckButton.svelte";

    export let title: string
    export let list: string[]
</script>
<div>
    {title}
    <ul class="flex flex-wrap">
        {#each list as item, index}
            <li class="p-0.5">
                <CheckButton index="{index}" text="{item}" on:checkButtonClicked></CheckButton>
            </li>
        {/each}
    </ul>
</div>