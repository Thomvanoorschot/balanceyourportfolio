<script lang="ts">
    export let type: "number" | "text" = "text"
    export let value: any | undefined
    export let placeholder: string | undefined = ""

    function typeAction(node: HTMLInputElement) {
        node.type = type;
    }

</script>

<input
        {placeholder}
        class="rounded-xl bg-secondary w-full outline-none text-center border-2 border-primary focus:border-quaternary pt-2 pb-2"
        use:typeAction
        bind:value={value}
>