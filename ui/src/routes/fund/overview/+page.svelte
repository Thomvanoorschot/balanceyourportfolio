<script lang="ts">
    import type {PageData} from './$types';
    import SearchBar from "$lib/search/SearchBar.svelte";
    import CheckButtonList from "$lib/filters/CheckButtonList.svelte";
    import List from "$lib/list/List.svelte";
    import DetailMenu from "$lib/menu/DetailMenu.svelte";
    import ListItem from "$lib/list/ListItem.svelte";
    import FundLineItem from "$lib/list/FundLineItem.svelte";

    export let data: PageData;
    let error: string | undefined
    let searchTerm: string | undefined
    let providers: string[] = ["Vanguard"]
    $: ({funds} = data);

    function submitNextPage(): void {
    }
</script>

{#if (!error && funds)}
    <div class="flex flex-grow items-start justify-between w-full gap-5 p-5">
        <DetailMenu>
            <SearchBar
                    placeholder="Fund name or ticker"
                    on:inputChanged={() => {}}
                    bind:value={searchTerm}
                    inPrimary="{false}"
            ></SearchBar>
            <CheckButtonList
                    title="Providers"
                    list="{providers}"
                    on:checkButtonClicked={() => {}}
            >
            </CheckButtonList>
        </DetailMenu>
        <div class="flex flex-col flex-grow gap-5">
            <List on:endOfPageReached={submitNextPage}>
                {#each funds as fund}
                    <a href="/fund/{fund.id}" class="hover:opacity-90">
                        <ListItem>
                            <FundLineItem fund="{fund}"></FundLineItem>
                        </ListItem>
                    </a>
                {/each}
            </List>
        </div>
    </div>
{/if}
