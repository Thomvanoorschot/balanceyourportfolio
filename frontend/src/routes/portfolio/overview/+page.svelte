<script lang="ts">
	import type { PageData } from './$types';
	import Portfolio from '$lib/portfolios/Portfolio.svelte';
	import type { PortfoliosResponse__Output } from '$lib/proto/proto/PortfoliosResponse';

	export let data: PageData;
	let portfolios: PortfoliosResponse__Output | undefined;
	let error: string | undefined;
	let isAuthenticated: boolean = false;
	$: {
		portfolios = data?.portfolios;
		error = '';
		isAuthenticated = data.isAuthenticated;
	}
</script>

{#if !error && portfolios}
	<div class="flex flex-col w-full items-center justify-start pt-5 gap-5">
		{#each portfolios.entries as portfolio}
			<Portfolio {isAuthenticated} {portfolio} />
		{/each}
		<Portfolio
			{isAuthenticated}
			portfolio={{
				entries: [{ name: '', id: '', fundId: '', amount: 0 }],
				name: `Portfolio ${portfolios.entries.length + 1}`,
				id: ''
			}}
		/>
	</div>
{:else}
	<h1>{error}</h1>
{/if}
