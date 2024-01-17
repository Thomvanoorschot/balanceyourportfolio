<script lang="ts">
	export let color: string;
	export let width: number;
</script>

<div
	class="h-3 lg:h-4"
	style=" background-color: {color}; width: {`${width}%`}"
/>
