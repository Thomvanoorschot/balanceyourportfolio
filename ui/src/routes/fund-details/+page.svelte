<script lang="ts">
import Filter from "$lib/fund-details/Filter.svelte";
import Information from "$lib/fund-details/Information.svelte";
import Weightings from "$lib/fund-details/Weightings.svelte";
import Holdings from "$lib/fund-details/Holdings.svelte";

</script>
<div id="fundDetails" class="flex flex-grow items-start justify-between w-full">
    <Filter></Filter>
    <div class="flex flex-col flex-grow">
        <div class="flex flex-col p-4">
            <Information></Information>
        </div>
        <div class="flex flex-col p-4">
            <Weightings></Weightings>
        </div>
        <ul class="bg-gray-300 p-4">
            <Holdings></Holdings>
        </ul>
    </div>
</div>