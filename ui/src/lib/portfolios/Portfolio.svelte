<script lang="ts">
    import ListItem from "$lib/portfolios/ListItem.svelte";
    import {EMPTY_UUID} from "$lib/utils";
    import CustomButton from "$lib/CustomButton.svelte";
    import {goto} from "$app/navigation";
    import type {Portfolio__Output} from "$lib/proto/proto/Portfolio";
    import type {PortfolioListItem__Output} from "$lib/proto/proto/PortfolioListItem";
    import {enhance} from '$app/forms';
    import type {ActionResult} from "@sveltejs/kit";

    let testForm: HTMLFormElement;


    export let portfolio: Portfolio__Output;
    // const portfoliosStore = getContext<PortfolioStore>("portfoliosStore")
    const addNewRow = (clickedRow: PortfolioListItem__Output) => {
        // if (portfolio.items[portfolio.items.length - 1] === clickedRow) {
        //     portfoliosStore.addEmptyItem(portfolio)
        // }
    }
    const updatePortfolio = () => {
        return ({result}: {
            result: ActionResult
        }) => {
            if (result.type === "success" && result.data) {
                portfolio = result.data.portfolio
            }
        };
    };
</script>

<form
        bind:this={testForm}
        method="POST"
        use:enhance={({formData}) => {
                     formData.set("portfolio", JSON.stringify(portfolio))
                         updatePortfolio()
                     return async ({ update, result }) => {
                     };
                }}
        action="?/upsertPortfolio"
        class="flex relative flex-col m-20 border-2 border-slate-800 p-2 w-[50vw]">
    <input
            bind:value={portfolio.name}
            class="text-2xl" name="name" type="text" placeholder="Portfolio name"
    >
    {#each portfolio.entries as entry}
        <ListItem listItem="{entry}" on:blurField={() => addNewRow(entry)}></ListItem>
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
                on:buttonClicked={() => {
                    testForm.requestSubmit()
                    "portfoliosStore.upsertPortfolio(portfolio)"
                }
                }
        ></CustomButton>
    </div>
</form>