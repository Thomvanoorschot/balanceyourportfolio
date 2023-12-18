<script lang="ts">
    import {createEventDispatcher, onMount} from "svelte";

    export let placeholder: string
    export let value: string | undefined = ""
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
<div class="flex justify-center w-full flex-1 lg:mr-32">
    <div
            class="relative w-full max-w-xl mr-6 focus-within:text-violet-500"
    >
        <div class="absolute inset-y-0 flex items-center pl-2">
            <svg
                    class="w-4 h-4"
                    aria-hidden="true"
                    fill="currentColor"
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
                class="w-full pt-2 pb-2 pl-8 pr-2 text-sm text-gray-700 placeholder-gray-600 bg-gray-100 border-0 rounded-md dark:placeholder-gray-500 dark:focus:shadow-outline-gray dark:focus:placeholder-gray-600 dark:bg-gray-700 dark:text-gray-200 focus:placeholder-gray-500 focus:bg-white focus:border-violet-300 focus:outline-none focus:shadow-outline-violet form-input"
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
