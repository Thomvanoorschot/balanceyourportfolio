<script lang="ts">
	export let value: number;
	export let disabled: boolean = true;
</script>

<td class="px-4 py-3 text-sm relative">
	<input
		{disabled}
		bind:value
		type="text"
		class="w-20 rounded-lg text-center outline-none border-2 border-white p-3 bg-primary {disabled
			? ''
			: 'focus:border-2 focus:border-quaternary focus:outline-none bg-secondary'}"
	/>
</td>
