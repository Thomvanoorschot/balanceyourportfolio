<script lang="ts">
    import {createEventDispatcher} from "svelte";
    const dispatch = createEventDispatcher()

    export let text: string = ""

</script>
<button
        class="w-full px-10 py-4 font-medium text-white transition-colors duration-150 bg-violet-600 border border-transparent rounded-lg active:bg-violet-600 hover:bg-violet-700 focus:outline-none focus:shadow-outline-violet"
        on:click={() => dispatch('buttonClicked')}
        type="button"
>
    {text}
</button>