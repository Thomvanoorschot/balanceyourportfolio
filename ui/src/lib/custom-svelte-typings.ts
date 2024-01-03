export const clickOutside = (node: HTMLElement | null) => {
    const handleClick = (event: MouseEvent) => {
        if (node && !node.contains(event.target as Node) && !event.defaultPrevented) {
            node.dispatchEvent(
                new CustomEvent('click_outside', { detail: node })
            );
        }
    };
    const handleTouch = (event: TouchEvent) => {
        if (node && !node.contains(event.target as Node) && !event.defaultPrevented) {
            node.dispatchEvent(
                new CustomEvent('click_outside', { detail: node })
            );
        }
    };
    document.addEventListener('click', handleClick, true);
    document.addEventListener('touchstart', handleTouch, true);
    return {
        destroy() {
            document.removeEventListener('click', handleClick, true);
        }
    };
}
