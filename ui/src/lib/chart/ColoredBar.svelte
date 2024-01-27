<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { onNavigate } from '$app/navigation';

	export let title: string;
	export let percentage: number;
	export let width: number;
	let count = 0;

	onNavigate(() => {
		count++;
	});
	const dispatch = createEventDispatcher();
</script>

{#key count}
	<div class="flex justify-start items-center text-primary">
		<div class="flex w-3/5 items-center lg:w-1/4">
			<div class="text-[0.5rem] lg:text-xs grow">{title}</div>
			<div class="text-[0.5rem] lg:text-xs pr-2">{Math.round(percentage * 100) / 100}%</div>
		</div>
		<div class="w-full flex">
			<div
				aria-hidden="true"
				class="flex w-full resizableElement hover:opacity-80 transition-opacity"
				on:click={() => dispatch('onBarClicked', title)}
			>
				<div class="flex h-full w-full rounded-md overflow-hidden" style=" width: {`${width}%`}">
					<slot />
				</div>
				<div class="bg-none flex-grow"></div>
			</div>
		</div>
	</div>
{/key}

<style>
    @keyframes changeWidth {
        0% {
            width: 0;
        }
        100% {
            width: 100%;
        }
    }

    .resizableElement {
        will-change: transform;
        animation: changeWidth 1s ease-in-out forwards;
    }
</style>
