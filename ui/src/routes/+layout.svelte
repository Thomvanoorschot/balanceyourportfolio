<script lang="ts">
    import "../app.css";
    import {onMount} from 'svelte';
    import {theme} from "$lib/stores/theme-store.ts";
    import SidebarElement from "$lib/layout/SidebarElement.svelte";
    import PortfolioIcon from "$lib/icons/PortfolioIcon.svelte";
    import ThemeSelector from "$lib/theme/ThemeSelector.svelte";
    import TopMenu from "$lib/menu/TopMenu.svelte";
    import FundSearchBar from "$lib/search/FundSearchBar.svelte";
    import {Toaster} from "svelte-french-toast";
    import type {LayoutData} from "../../.svelte-kit/types/src/routes/$types";
    import FundIcon from "$lib/icons/FundIcon.svelte";
    import {afterNavigate} from "$app/navigation";

    export let data: LayoutData;
    let {profilePictureURL} = data
    onMount(() => {
        // We load the in the <script> tag in load, but then also here onMount to setup stores
        if (!('theme' in localStorage)) {
            theme.useLocalStorage();
            theme.set({...$theme, mode: 'light'});
        } else {
            theme.useLocalStorage();
        }
    });
    let isSideMenuOpen: boolean

    afterNavigate(() => {
        document.getElementById('page')?.scroll(0, 0);
    });
</script>


<Toaster position="top-right" toastOptions="{{iconTheme: {
		primary: '#f582ae',
		secondary: '#001858',
	}}}"/>
<div
        class="flex h-screen bg-primary {isSideMenuOpen ? 'overflow-hidden' : ''}"
>
    <!-- Desktop sidebar -->
    <aside
            class="z-20 hidden w-64 overflow-y-auto md:block flex-shrink-0"
    >
        <div class="py-4 text-gray-500">
            <a
                    class="ml-6 text-lg font-bold text-tertiary"
                    href="/portfolio/overview"
            >
                ETF Insight
            </a>
            <ul class="mt-6">
                <SidebarElement href="/fund/overview" text="Funds">
                    <FundIcon></FundIcon>
                </SidebarElement>
                <SidebarElement href="/portfolio/overview" text="Portfolio">
                    <PortfolioIcon></PortfolioIcon>
                </SidebarElement>
            </ul>
        </div>
    </aside>
    <!-- Mobile sidebar -->
    {#if (isSideMenuOpen)}
        <aside
                class="fixed inset-y-0 z-20 flex-shrink-0 w-64 mt-16 overflow-y-auto md:hidden "
        >
            <div class="py-4 text-gray-500">
                <a
                        class="ml-6 text-lg font-bold text-gray-800"
                        href="/search"
                >
                    ETF Insight
                </a>
                <ul class="mt-6">
                    <SidebarElement href="/fund/overview" text="Funds">
                        <FundIcon></FundIcon>
                    </SidebarElement>
                    <SidebarElement href="/portfolio/overview" text="Portfolio">
                        <PortfolioIcon></PortfolioIcon>
                    </SidebarElement>
                </ul>
            </div>
        </aside>
    {/if}
    <div class="flex flex-col flex-1 w-full">
        <header class="z-10 py-2 bg-background-primary">
            <div
                    class="flex items-center justify-between h-full text-tertiary"
            >
                <div class="flex-1">

                <!-- Mobile hamburger -->
                <button
                        class="p-1 mr-5 -ml-1 rounded-xl md:hidden focus:outline-none"
                        aria-label="Menu"
                        on:click={() => isSideMenuOpen = !isSideMenuOpen}
                >
                    <svg
                            class="w-6 h-6 fill-tertiary"
                            aria-hidden="true"
                            viewBox="0 0 20 20"
                    >
                        <path
                                fill-rule="evenodd"
                                d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z"
                                clip-rule="evenodd"
                        ></path>
                    </svg>
                </button>
                <!-- Search input -->
                <FundSearchBar></FundSearchBar>
                </div>
                <div class="flex gap-5 justify-end pr-5">
                    <ThemeSelector></ThemeSelector>
                    <TopMenu profilePictureURL="{profilePictureURL}"></TopMenu>
                </div>
            </div>
        </header>
        <main class="h-full overflow-y-auto bg-secondary">
            <slot></slot>
        </main>
        <footer class="bg-primary text-center">
            <div class="p-0.5 text-center text-neutral-800">
                © 2023 Copyright:
                <a class="text-neutral-800" href="https://tailwind-elements.com/">EtfInsight</a>
            </div>
        </footer>
    </div>
</div>