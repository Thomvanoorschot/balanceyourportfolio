<script lang="ts">
</script>

<div class="w-full mb-8 rounded-lg">
    <div class="w-full">
        <table class="w-full">
            <thead>
            <slot name="headerRow"/>

            </thead>
            <tbody
                    class="divide-y"
            >
            <slot/>
            </tbody>
        </table>
    </div>
</div>