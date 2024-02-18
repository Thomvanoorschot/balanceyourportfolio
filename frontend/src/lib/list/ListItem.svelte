<script lang="ts">
	export let isTertiary: boolean = true
</script>
<div class="{isTertiary ? 'bg-tertiary' : 'bg-primary'} rounded-xl items-center p-4 {isTertiary ? 'text-primary' : 'text-tertiary'}">
	<slot />
</div>
