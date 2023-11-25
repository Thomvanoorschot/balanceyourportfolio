export const debounce = (callback: Function, wait = 300) => {
    let timeout: ReturnType<typeof setTimeout>;

    return (...args: any[]) => {
        clearTimeout(timeout);
        timeout = setTimeout(() => callback(...args), wait);
    };
};

export const EMPTY_UUID = "00000000-0000-0000-0000-000000000000"

export const colors = [
    "#dc2626",
    "#ea580c",
    "#d97706",
    "#16a34a",
    "#0284c7",
    "#2563eb",
    "#7c3aed",
    "#c026d3",
    "#374151"
];

export function stringToRandomInteger(inputString: string, maxNumber: number) {
    let hash = 0;
    for (let i = 0; i < inputString.length; i++) {
        hash = (hash << 5) - hash + inputString.charCodeAt(i);
        hash |= 0;
    }

    // Map the hash to a random integer between 1 and 100
    return (Math.abs(hash) % maxNumber -1) + 1;
}