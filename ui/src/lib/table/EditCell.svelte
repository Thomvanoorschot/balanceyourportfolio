<script lang="ts">
    import {createEventDispatcher} from "svelte";
    import EditIcon from "$lib/icons/EditIcon.svelte";
    import DeleteIcon from "$lib/icons/DeleteIcon.svelte";

    export let disabled: boolean = true

    const dispatch = createEventDispatcher()

    function editClicked() {
        disabled = !disabled
        dispatch('editClicked')
    }

    function deleteClicked() {
        dispatch('deleteClicked')
    }
</script>
<td class="px-4 py-3">
    <div class="flex items-center space-x-4 text-sm">
        <button
                class="flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-primary rounded-lg dark:text-gray-400 focus:outline-none focus:shadow-outline-gray"
                aria-label="Edit"
                on:click={editClicked}
                type="button"
        >
            <EditIcon></EditIcon>
        </button>
        <button
                class="flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-primary rounded-lg dark:text-gray-400 focus:outline-none focus:shadow-outline-gray"
                aria-label="Delete"
                type="button"
                on:click={deleteClicked}
        >
            <DeleteIcon></DeleteIcon>
        </button>
    </div>
</td>