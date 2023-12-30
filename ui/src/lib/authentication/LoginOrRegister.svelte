<script lang="ts">

    import SecondaryButton from "$lib/shared/SecondaryButton.svelte";
    import TertiaryButton from "$lib/shared/TertiaryButton.svelte";
    import {goto} from "$app/navigation";
</script>
<div class="flex flex-col p-8 w-128 bg-primary gap-4 rounded-xl">
    <div class="text-center text-tertiary ">
        You need an account to create portfolio's
    </div>
    <SecondaryButton on:buttonClicked={() => goto("/api/auth/login")}>
        Login
    </SecondaryButton>
    <TertiaryButton on:buttonClicked={() => goto("/api/auth/register")}>
        Register
    </TertiaryButton>
</div>