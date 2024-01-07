<script lang="ts">
	import CheckButton from '$lib/filters/CheckButton.svelte';
	import type { themeType } from '$lib/utils.ts';

	export let title: string;
	export let list: string[];
	export let theme: themeType = 'primary';

</script>

<div class="text-center">
	{title}
	<ul class="flex flex-wrap">
		{#each list as item, index}
			<li class="p-0.5">
				<CheckButton theme="{theme}" {index} text={item} on:checkButtonClicked />
			</li>
		{/each}
	</ul>
</div>
