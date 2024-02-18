<script lang="ts">
	export let value: string;
	export let disabled: boolean = true;
</script>

<td class="px-4 py-3 text-sm relative">
	<input
		{disabled}
		bind:value
		type="text"
		class="w-20 text-center text-card-text outline-none border-2 border-white p-3 {disabled
			? ''
			: 'focus:border-2 focus:border-primary focus:outline-none bg-card-secondary rounded-lg'}"
	/>
</td>
