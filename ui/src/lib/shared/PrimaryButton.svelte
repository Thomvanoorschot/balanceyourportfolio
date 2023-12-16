<script lang="ts">
    import {createEventDispatcher} from "svelte";
    const dispatch = createEventDispatcher()

    export let text: string = ""

</script>
<button
        class="w-full px-10 py-4 font-medium text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple"
        on:click={() => dispatch('buttonClicked')}
        type="button"
>
    {text}
</button>