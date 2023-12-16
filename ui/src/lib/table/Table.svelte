<script lang="ts">
</script>

<div class="w-full mb-8 rounded-lg shadow-xs">
    <div class="w-full">
        <table class="w-full">
            <thead>
            <slot name="headerRow"/>

            </thead>
            <tbody
                    class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800"
            >
            <slot/>
            </tbody>
        </table>
    </div>
</div>