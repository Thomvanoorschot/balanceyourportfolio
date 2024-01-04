<script lang="ts">
	import type { Portfolio__Output } from '$lib/proto/proto/Portfolio';
	import { enhance } from '$app/forms';
	import type { ActionResult } from '@sveltejs/kit';
	import Table from '$lib/table/Table.svelte';
	import TableRow from '$lib/table/TableRow.svelte';
	import EditCell from '$lib/table/EditCell.svelte';
	import SearchFundCell from '$lib/table/SearchFundCell.svelte';
	import NumberCell from '$lib/table/NumberCell.svelte';
	import { goto } from '$app/navigation';
	import PrimaryButton from '$lib/shared/PrimaryButton.svelte';
	import SecondaryButton from '$lib/shared/SecondaryButton.svelte';
	import TableHeaderRow from '$lib/table/TableHeaderRow.svelte';
	import TableHeader from '$lib/table/TableHeader.svelte';
	import AddButton from '$lib/shared/AddButton.svelte';
	import toast from 'svelte-french-toast';
	import Modal from '$lib/shared/Modal.svelte';
	import LoginOrRegister from '$lib/authentication/LoginOrRegister.svelte';

	export let isAuthenticated: boolean = false;
	export let portfolio: Portfolio__Output;
	let portfolioForm: HTMLFormElement;
	let disabledList: boolean[] = new Array(portfolio.entries.length).fill(true);
	let showModal: boolean = false;

	const addNewRow = () => {
		portfolio.entries.push({
			amount: 0,
			fundId: '',
			name: '',
			id: ''
		});
		disabledList.push(false);
		portfolio = portfolio;
	};
	const upsertPortfolio = (message: string) => {
		return ({ result }: { result: ActionResult }) => {
			if (result.type === 'success' && result.data) {
				portfolio = result.data.portfolio;
				toast.success(message);
			} else {
				toast.error('Encountered an error');
			}
		};
	};
	const removeFund = (index: number) => {
		portfolio.entries.splice(index, 1);
		disabledList.splice(index, 1);
		portfolio = portfolio;
	};

	const handleButtonClicked = () => {
		if (isAuthenticated) {
			portfolioForm.requestSubmit();
			return;
		}
		showModal = true;
	};
</script>

<Modal bind:showModal>
	<LoginOrRegister />
</Modal>
<form
	class="flex relative flex-col m-20 w-[50vw] rounded-lg bg-tertiary"
	bind:this={portfolioForm}
	method="POST"
	use:enhance={({ formData }) => {
		formData.set('portfolio', JSON.stringify(portfolio));
		return upsertPortfolio(portfolio.id ? 'Updated portfolio' : 'Created portfolio');
	}}
	action="?/upsertPortfolio"
>
	<input
		bind:value={portfolio.name}
		type="text"
		class="text-center rounded-t-lg bg-tertiary text-primary outline-none"
	/>
	<Table>
		<TableHeaderRow slot="headerRow">
			<TableHeader>Ticker or name</TableHeader>
			<TableHeader>Amount</TableHeader>
			<TableHeader>
				<AddButton on:buttonClicked={addNewRow} />
			</TableHeader>
		</TableHeaderRow>
		{#each portfolio.entries as row, index}
			<TableRow>
				<SearchFundCell
					disabled={disabledList[index]}
					bind:label={row.name}
					bind:value={row.fundId}
				/>
				<NumberCell disabled={disabledList[index]} bind:value={row.amount} />
				<EditCell
					on:editClicked={() => (disabledList[index] = !disabledList[index])}
					on:deleteClicked={() => removeFund(index)}
				/>
			</TableRow>
		{/each}
	</Table>
	<div class="flex">
		{#if portfolio.id}
			<div class="w-full p-2">
				<PrimaryButton on:buttonClicked={() => goto(`/portfolio/${portfolio.id}`)}>
					Details
				</PrimaryButton>
			</div>
		{/if}
		<div class="w-full p-2">
			<SecondaryButton on:buttonClicked={handleButtonClicked}>
				{portfolio.id ? 'Update portfolio' : 'Create portfolio'}
			</SecondaryButton>
		</div>
	</div>
</form>
