<script lang="ts">
    import {createEventDispatcher} from "svelte";
    const dispatch = createEventDispatcher()
</script>
<button
        class="w-full px-10 py-4 font-medium leading-5 text-gray-500 transition-colors duration-150 border-primary bg-secondary border-2 rounded-lg active:bg-primary hover:bg-secondaryhighlighted focus:outline-none focus:shadow-outline-primary"
        on:click={() => dispatch('buttonClicked')}
        type="button"
>
    <slot></slot>
</button>