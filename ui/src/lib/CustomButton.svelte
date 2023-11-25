<script>
import { createEventDispatcher } from 'svelte'
const dispatch = createEventDispatcher()

export let buttonText = ""
function buttonClicked() {
    dispatch('buttonClicked')
}
</script>
<div
        tabindex="0"
        aria-label=""
        role="button"
        on:keydown={buttonClicked}
        on:click={buttonClicked}
        class="cursor-pointer border-solid border-2 rounded border-slate-800 hover:bg-slate-100 w-full text-center p-1"
>
    {buttonText}
</div>