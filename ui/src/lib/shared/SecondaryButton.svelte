<script lang="ts">
    import {createEventDispatcher} from "svelte";
    const dispatch = createEventDispatcher()

    export let text: string = ""

</script>
<button
        class="w-full px-10 py-4 font-medium leading-5 text-gray-500 transition-colors duration-150 border-violet-600 bg-violet-50 border-2 rounded-lg active:bg-violet-600 hover:bg-violet-200 focus:outline-none focus:shadow-outline-violet"
        on:click={() => dispatch('buttonClicked')}
        type="button"
>
    {text}
</button>