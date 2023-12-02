<script lang="ts">
    import type {FundInformation__Output} from "$lib/proto/proto/FundInformation";

    export let fundInformation: FundInformation__Output;

</script>
<h1 class="text-3xl ">{fundInformation.name}</h1>
<div>Outstanding shares: {fundInformation.outstandingShares}</div>
<div>Effective date: {fundInformation.effectiveDate}</div>