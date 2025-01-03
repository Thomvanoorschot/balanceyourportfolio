import { writable } from 'svelte/store';

export const createWritableStore = (key: string, startValue: any): any => {
	const { subscribe, set } = writable(startValue);

	return {
		subscribe,
		set,
		useLocalStorage: (): string => {
			const json = localStorage.getItem(key);
			if (json) {
				set(JSON.parse(json));
			}

			subscribe((current) => {
				localStorage.setItem(key, JSON.stringify(current));
			});
			return json || '';
		}
	};
};

export const theme = createWritableStore('theme', { mode: 'light' });
