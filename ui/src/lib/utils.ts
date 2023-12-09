export const debounce = (callback: Function, wait = 300) => {
    let timeout: ReturnType<typeof setTimeout>;

    return (...args: any[]) => {
        clearTimeout(timeout);
        timeout = setTimeout(() => callback(...args), wait);
    };
};

export const EMPTY_UUID = "00000000-0000-0000-0000-000000000000"

export const colors = [
    "#cb8655",
    "#e09b80",
    "#7e8e65",
    "#deb3e0",
    "#b65149",
    "#feecd6",
    "#a39fe1",
    "#9bb8ed"
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