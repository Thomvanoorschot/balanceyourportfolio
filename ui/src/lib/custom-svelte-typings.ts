declare namespace svelteHTML {
    interface HTMLAttributes<T> {
        'on:click_outside'?: (event: CustomEvent) => void;
    }
}
export const clickOutside = (node: HTMLElement | null) => {
    const handleClick = (event: MouseEvent) => {
        if (node && !node.contains(event.target as Node) && !event.defaultPrevented) {
            node.dispatchEvent(
                new CustomEvent('click_outside', { detail: node })
            );
        }
    };
    document.addEventListener('click', handleClick, true);
    return {
        destroy() {
            document.removeEventListener('click', handleClick, true);
        }
    };
}
