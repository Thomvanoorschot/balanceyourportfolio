<script lang="ts">
    export let text: string
</script>
<li class="relative px-6 py-3">
              <span
                      class="absolute inset-y-0 left-0 w-1 bg-primary rounded-tr-lg rounded-br-lg"
                      aria-hidden="true"
              ></span>
    <a
            class="inline-flex items-center w-full text-sm font-semibold text-gray-800 transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200 dark:text-gray-100"
            href="/portfolio/overview"
    >
        <slot></slot>
        <span class="ml-4">{text}</span>
    </a>
</li>