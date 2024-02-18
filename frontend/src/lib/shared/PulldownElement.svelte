<script lang="ts">
	import MenuIcon from '$lib/icons/MenuIcon.svelte';

	export let text: string;
	export let selector: string;
	export let size: string = 'text-l';
	let expanded: boolean = false;
	let elem: HTMLDivElement;

	export function growDiv() {
		if (expanded) {
			elem.style.height = '0';
			expanded = false;
		} else {
			const wrapper = document.querySelector('#' + selector);
			elem.style.height = wrapper?.clientHeight + 'px';
			expanded = true;
		}
	}
</script>

<div class="flex items-center">
	<div>
		<MenuIcon fillColor="fill-primary" />
	</div>
	<h1 class="pl-2 {size}">{text}</h1>
</div>
<div
	bind:this={elem}
	class="h-0 transition-all ease-in-out duration-500 overflow-hidden"
>
	<div id="{selector}" class="flex flex-col gap-5">
		<slot></slot>
	</div>
</div>