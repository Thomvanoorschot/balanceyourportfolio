<script lang="ts">
    import {createEventDispatcher} from "svelte";
    const dispatch = createEventDispatcher()
</script>
<button
        class="w-full px-10 py-4 font-medium text-white transition-colors duration-150 bg-primary border border-transparent rounded-lg hover:bg-primaryhighlighted focus:outline-none"
        on:click={() => dispatch('buttonClicked')}
        type="button"
>
    <slot></slot>
</button>