<script lang="ts">
    import type {FundInformation} from "$lib/fund";

    export let fundInformation: FundInformation;

</script>
<h1 class="text-3xl ">{fundInformation.name}</h1>
<div>Outstanding shares: {fundInformation.outstandingShares}</div>
<div>Effective date: {fundInformation.effectiveDate}</div>