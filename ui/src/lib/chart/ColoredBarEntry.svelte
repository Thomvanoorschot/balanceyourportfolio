<script lang="ts">
    export let color: string
    export let width: number
    export let roundedLeft: boolean | undefined = false
    export let roundedRight: boolean | undefined = false
</script>

<div class="h-4 bg-blue-300 {roundedLeft ? 'rounded-l-md' : ''} {roundedRight ? 'rounded-r-md' : ''}" style=" background-color: {color}; width: { `${width}%`}"></div>