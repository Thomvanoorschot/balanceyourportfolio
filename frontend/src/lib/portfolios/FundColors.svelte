<script lang="ts">
	export let colorMap: Map<string, { fundName: string; color: string }>;
	export let textColor: string | undefined = '';
</script>

<div class="flex gap-2 flex-wrap">
	{#each [...colorMap] as [key, value]}
		<div class="flex items-center">
			<div class="aspect-square rounded-full h-3 w-3" style="background-color: {value.color}" />
			<div class="pl-2 text-xs {textColor}">{value.fundName}</div>
		</div>
	{/each}
</div>
