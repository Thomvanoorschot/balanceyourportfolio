<script lang="ts">
    import {createEventDispatcher, onMount} from "svelte";

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
            <svg
                    class="w-4 h-4 fill-tertiary"
                    aria-hidden="true"
                    viewBox="0 0 20 20"
            >
                <path
                        fill-rule="evenodd"
                        d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                        clip-rule="evenodd"
                ></path>
            </svg>
        </div>
        <input
                name="searchTerm"
                class="w-full pt-2 pb-2 pl-8 pr-2 text-sm placeholder-gray-600
                border-2 rounded-xl focus:outline-none form-input text-tertiary
                {inPrimary ?
                    'bg-secondary border-primary focus:border-quaternary' :
                     'bg-primary border-secondary focus:border-quaternary'
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
