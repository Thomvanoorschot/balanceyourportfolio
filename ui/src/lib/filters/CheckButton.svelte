<script lang="ts">
    import {createEventDispatcher} from "svelte";

    export let index: number
    export let text: string
    export let checked: boolean | undefined = false

    const dispatch = createEventDispatcher()
</script>
<div class="flex">
    <input type="checkbox" id="check-button-{index}" class="peer hidden" bind:checked={checked}
           on:click={() => dispatch("checkButtonClicked", text)}/>
    <label
            for="check-button-{index}"
            class="select-none cursor-pointer rounded-lg border-2 border-gray-200
   py-2 px-2.5 font-small text-xs text-gray-500 transition-colors duration-200 ease-in-out peer-checked:bg-violet-200 peer-checked:border-violet-600 peer-hover:bg-violet-300 peer-hover:border-violet-600 ">
        {text}
    </label>
</div>