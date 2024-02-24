<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { theme } from '$lib/stores/theme-store.ts';
	import SidebarElement from '$lib/layout/SidebarElement.svelte';
	import PortfolioIcon from '$lib/icons/PortfolioIcon.svelte';
	import ThemeSelector from '$lib/theme/ThemeSelector.svelte';
	import TopMenu from '$lib/menu/TopMenu.svelte';
	import FundSearchBar from '$lib/search/FundSearchBar.svelte';
	import { Toaster } from 'svelte-french-toast';
	import type { LayoutData } from '../../.svelte-kit/types/src/routes/$types';
	import FundIcon from '$lib/icons/FundIcon.svelte';
	import { afterNavigate } from '$app/navigation';
	import MenuIcon from '$lib/icons/MenuIcon.svelte';
	import CompareIcon from '$lib/icons/CompareIcon.svelte';

	export let data: LayoutData;
	let { profilePictureURL } = data;
	onMount(() => {
		// We load the in the <script> tag in load, but then also here onMount to setup stores
		if (!('theme' in localStorage)) {
			theme.useLocalStorage();
			theme.set({ ...$theme, mode: 'light' });
		} else {
			theme.useLocalStorage();
		}
	});
	let isSideMenuOpen: boolean;

	afterNavigate(() => {
		isSideMenuOpen = false;
	});
</script>

<Toaster
	position="top-right"
	toastOptions={{
		iconTheme: {
			primary: '#f582ae',
			secondary: '#001858'
		}
	}}
/>
<div class="flex h-screen bg-primary {isSideMenuOpen ? 'overflow-hidden' : ''}">
	<!-- Desktop sidebar -->
	<aside class="z-20 hidden w-64 overflow-y-auto md:block flex-shrink-0">
		<div class="py-4 text-gray-500">
			<a class="ml-6 text-lg font-bold text-tertiary" href="/portfolio/overview"> BalanceYourPortfolio </a>
			<ul class="mt-6">
				<SidebarElement href="/fund/overview" text="Funds">
					<FundIcon />
				</SidebarElement>
				<SidebarElement href="/fund/compare" text="Compare Funds">
					<CompareIcon />
				</SidebarElement>
				<SidebarElement href="/portfolio/overview" text="Portfolio">
					<PortfolioIcon />
				</SidebarElement>
			</ul>
		</div>
	</aside>
	<!-- Mobile sidebar -->
	{#if isSideMenuOpen}
		<aside
			class="absolute h-max inset-y-0 z-20 flex-shrink-0 w-64 mt-14 overflow-y-auto md:hidden bg-primary"
		>
			<div class="py-4 text-gray-500">
				<ul>
					<SidebarElement href="/fund/overview" text="Funds">
						<FundIcon />
					</SidebarElement>
					<SidebarElement href="/fund/compare" text="Compare Funds">
						<CompareIcon />
					</SidebarElement>
					<SidebarElement href="/portfolio/overview" text="Portfolio">
						<PortfolioIcon />
					</SidebarElement>
				</ul>
			</div>
		</aside>
	{/if}
	<div class="flex flex-col flex-1 w-full">
		<header class="py-2 bg-background-primary">
			<div class="flex items-center gap-2">
				<!-- Mobile hamburger -->
				<button
					class="pl-2 rounded-xl md:hidden focus:outline-none"
					aria-label="Menu"
					on:click={() => (isSideMenuOpen = !isSideMenuOpen)}
				>
					<MenuIcon></MenuIcon>
				</button>
				<!-- Search input -->
				<FundSearchBar theme="secondary" />
				<div class="flex gap-5 justify-end pr-5">
<!--					<ThemeSelector />-->
					<TopMenu {profilePictureURL} />
				</div>
			</div>
		</header>
		<main class="h-full w-full overflow-x-hidden bg-secondary">
			<slot />
		</main>
	</div>
</div>
