<script lang="ts">
	import type { FundInformation__Output } from '$lib/proto/proto/FundInformation';

	export let fundInformation: FundInformation__Output;
</script>

<div class="p-2 text-center shadow-md rounded-xl bg-tertiary text-primary">
	<h1 class="text-xl">{fundInformation.name}</h1>
</div>
