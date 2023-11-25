<script lang="ts">
    import ListItem from "$lib/portfolios/ListItem.svelte";
    import type {Portfolio, PortfolioListItem} from "$lib/portfolio";
    import {getContext} from "svelte";
    import type {PortfolioStore} from "$lib/stores/portfolio-store";
    import {EMPTY_UUID} from "$lib/utils";
    import CustomButton from "$lib/CustomButton.svelte";
    import {goto} from "$app/navigation";

    export let portfolio: Portfolio;
    const portfoliosStore = getContext<PortfolioStore>("portfoliosStore")
    const addNewRow = (clickedRow: PortfolioListItem) => {
        if (portfolio.items[portfolio.items.length - 1] === clickedRow) {
            portfoliosStore.addEmptyItem(portfolio)
        }
    }

</script>

<div class="flex relative flex-col m-20 border-2 border-slate-800 p-2 w-[50vw]">
    <input
            bind:value={portfolio.name}
            class="text-2xl" name="name" type="text" placeholder="Portfolio name"
    >
    {#each portfolio.items as item}
        <ListItem listItem="{item}" on:blurField={() => addNewRow(item)}></ListItem>
    {/each}
    <div class="flex pt-5 gap-2">
        {#if portfolio.id !== EMPTY_UUID }
            <CustomButton
                    buttonText="Details"
                    on:buttonClicked={() => goto(`/portfolio/${portfolio.id}`)}
            ></CustomButton>
        {/if}
        <CustomButton

                buttonText="{portfolio.id === EMPTY_UUID ? 'Create portfolio' : 'Update portfolio'}"
                on:buttonClicked={() => portfoliosStore.upsertPortfolio(portfolio)}
        ></CustomButton>
    </div>
</div>