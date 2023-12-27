<script lang="ts">
    import {onMount} from "svelte";
    import type {Portfolio__Output} from "$lib/proto/proto/Portfolio.ts";
    import {page} from "$app/stores";
    import type {PortfolioListItem__Output} from "$lib/proto/proto/PortfolioListItem.ts";
    import TertiaryButton from "$lib/shared/TertiaryButton.svelte";
    import toast from "svelte-french-toast";
    import type {PatchValueRequest} from "../../routes/api/portfolio/+server.ts";
    import Dropdown from "$lib/dropdown/Dropdown.svelte";
    import Input from "$lib/input/Input.svelte";

    export let showModal: boolean

    let selectedPortfolioId: string
    let existingPortfolioFund: PortfolioListItem__Output | undefined
    let newPortfolioFundAmount: number | undefined = 0
    let portfolios: Portfolio__Output[] | undefined = []
    onMount(async () => {
            const resp = await fetch("/api/portfolio", {
                method: 'GET',
            })
            if (resp.ok) {
                portfolios = await resp.json() as Portfolio__Output[]
            } else if (!resp.ok) {
                // error = result.data?.error
            }
        }
    )
    const setValues = (portfolio: Portfolio__Output | undefined) => {
        if (!portfolio) {
            return
        }
        if (portfolio.entries.some((x: PortfolioListItem__Output) => x.fundId === $page.params.id)) {
            existingPortfolioFund = portfolio.entries.find((x: PortfolioListItem__Output) => x.fundId === $page.params.id)
        } else {
            existingPortfolioFund = undefined
        }
    }
    const portfolioChanged = () => {
        setValues(portfolios?.find(x => x.id === selectedPortfolioId))
    }
    const updateAmount = async () => {
        const req: PatchValueRequest = {
            amount: existingPortfolioFund?.amount || newPortfolioFundAmount || 0,
            fundId: $page.params.id,
            portfolioId: selectedPortfolioId,
        }
        const resp = await fetch("/api/portfolio", {
            method: 'PATCH',
            body: JSON.stringify(req)
        })
        if (resp.ok) {
            toast.success("Test");
        } else if (!resp.ok) {
            toast.error(resp.statusText);
        }
        showModal = false
    }
</script>


{#if portfolios}
    <div class="flex flex-col p-8 w-128 bg-primary gap-4 rounded-xl">
        <Dropdown
                bind:value={selectedPortfolioId}
                entries="{portfolios.map(x => ({label: x.name, value: x.id}))}"
                on:optionChanged={portfolioChanged}
        ></Dropdown>
        {#if (existingPortfolioFund)}
            <Input
                    placeholder="Amount"
                    bind:value={existingPortfolioFund.amount}
                    type="number"
            ></Input>
            <TertiaryButton
                    disabled="{!selectedPortfolioId}"
                    on:buttonClicked={updateAmount}
            >
                Update amount
            </TertiaryButton>
        {:else}
            <Input
                    placeholder="Amount"
                    bind:value={newPortfolioFundAmount}
                    type="number"
            ></Input>
            <TertiaryButton
                    disabled="{!selectedPortfolioId}"
                    on:buttonClicked={updateAmount}
            >
                Add to portfolio
            </TertiaryButton>
        {/if}
    </div>
{/if}
