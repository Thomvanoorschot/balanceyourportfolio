<script lang="ts">
    import type {FundInformation__Output} from "$lib/proto/proto/FundInformation";

    export let fundInformation: FundInformation__Output;

</script>
<div class="p-10 shadow-md rounded-xl bg-tertiary text-primary">
    <h1 class="text-3xl ">{fundInformation.name}</h1>
    <div>Outstanding shares: {fundInformation.outstandingShares}</div>
    <div>Effective date: {fundInformation.effectiveDate}</div>
</div>
