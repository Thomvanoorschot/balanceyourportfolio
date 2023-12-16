<script lang="ts">
    import {createEventDispatcher} from "svelte";
    const dispatch = createEventDispatcher()

    export let text: string = ""

</script>
<button
        class="w-full px-10 py-4 font-medium leading-5 text-gray-500 transition-colors duration-150 border-purple-600 bg-purple-50 border-2 rounded-lg active:bg-purple-600 hover:bg-purple-200 focus:outline-none focus:shadow-outline-purple"
        on:click={() => dispatch('buttonClicked')}
        type="button"
>
    {text}
</button>