<script lang="ts">



import ListItem from "$lib/portfolios/ListItem.svelte";
</script>

<div class="flex relative flex-col m-20 border-2 border-slate-800 p-2 w-[50vw]">
        <input class="text-2xl" name="name" type="text" placeholder="Portfolio name">
        <ListItem></ListItem>
        <ListItem></ListItem>
        <ListItem></ListItem>
        <div class="cursor-pointer flex justify-center pl-40 pr-40 pt-1 pb-1 mt-3 border-solid border-2 rounded  border-slate-800 p-2 hover:bg-slate-100">
            Create portfolio
        </div>
</div>