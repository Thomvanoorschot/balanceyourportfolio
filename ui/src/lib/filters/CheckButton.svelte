<script lang="ts">
    export let index: number
    export let text: string
    export let checked:boolean | undefined = false
</script>
<div class="flex">
    <input type="checkbox" id="check-button-{index}" class="peer hidden"  bind:checked={checked} />
    <label for="check-button-{index}" class="select-none cursor-pointer rounded-lg border-2 border-gray-200
   py-2 px-3 font-small text-xs text-gray-500 transition-colors duration-200 ease-in-out peer-checked:bg-purple-200 peer-checked:border-purple-600 peer-hover:bg-purple-300 peer-hover:border-purple-600 ">
        {text}
    </label>
</div>