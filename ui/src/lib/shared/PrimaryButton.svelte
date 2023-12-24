<script lang="ts">
    import {createEventDispatcher} from "svelte";
    const dispatch = createEventDispatcher()
</script>
<button
        class="w-full px-10 py-4 font-medium text-button-text transition-colors duration-150 bg-primary border border-transparent rounded-lg hover:bg-quaternary focus:outline-none"
        on:click={() => dispatch('buttonClicked')}
        type="button"
>
    <slot></slot>
</button>