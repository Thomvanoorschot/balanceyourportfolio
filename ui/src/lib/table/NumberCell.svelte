<script lang="ts">
    export let value: number
    export let disabled: boolean = true
</script>


<td class="px-4 py-3 text-sm relative">
    <input
            disabled="{disabled}"
            bind:value="{value}"
            type="text"
            class="w-20 text-center outline-none border-2 border-white p-3 {disabled ? '' : 'focus:border-2 focus:border-violet-600 focus:outline-none bg-violet-50 rounded-lg'}"
    >
</td>
