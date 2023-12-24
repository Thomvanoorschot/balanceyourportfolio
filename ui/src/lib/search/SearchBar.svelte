<script lang="ts">
    import {createEventDispatcher, onMount} from "svelte";
    import SearchIcon from "$lib/icons/SearchIcon.svelte";

    export let placeholder: string
    export let value: string | undefined = ""
    export let inPrimary: boolean | undefined = true
    const dispatch = createEventDispatcher()

    function inputChanged() {
        dispatch('inputChanged')
    }

    let hasFocus: boolean
    onMount(() => {
        document.addEventListener('keydown', closeOnEscape);
        return () => document.removeEventListener('keydown', closeOnEscape);
    });
    const closeOnEscape = (e: KeyboardEvent) => {
        if (e.key === 'Escape' && hasFocus)
            hasFocus = false
    }
</script>
<div class="flex justify-center w-full flex-1">
    <div
            class="relative w-full max-w-xl focus-within:text-tertiary"
    >
        <div class="absolute inset-y-0 flex items-center pl-2">
            <SearchIcon inPrimary="{inPrimary}"></SearchIcon>
        </div>
        <input
                name="searchTerm"
                class="w-full pt-2 pb-2 pl-8 pr-2 text-sm placeholder-gray-600
                border-2 rounded-xl focus:outline-none form-input
                {inPrimary ?
                    'bg-secondary border-primary focus:border-quaternary' :
                     'bg-tertiary border-secondary focus:border-quaternary placeholder-primary text-primary'
                }"
                type="text"
                placeholder={placeholder}
                aria-label="Search"
                bind:value={value}
                on:input={inputChanged}
                on:click={() => hasFocus = true}
                on:blur={() => setTimeout(() => {hasFocus = false}, 200)}
        />
        {#if (hasFocus)}
            <slot></slot>
        {/if}
    </div>
</div>
