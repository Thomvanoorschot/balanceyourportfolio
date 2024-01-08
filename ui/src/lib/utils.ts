export const debounce = (callback: Function, wait = 300) => {
	let timeout: ReturnType<typeof setTimeout>;

	return (...args: any[]) => {
		clearTimeout(timeout);
		timeout = setTimeout(() => callback(...args), wait);
	};
};

export interface LabelValue {
	label: string;
	value: any;
}

export const EMPTY_UUID = '00000000-0000-0000-0000-000000000000';

export const colors = [
	'#cb8655',
	'#e09b80',
	'#7e8e65',
	'#deb3e0',
	'#b65149',
	'#a39fe1',
	'#9bb8ed',
	'#800020',
	'#777F8C',
	'#FC764A',
	'#008080',
	'#C6989F',
	'#40E0D0'
];

interface Map {
	[key: string]: string | undefined;
}

export const currencySignMap: Map = {
	EUR: '€',
	USD: '$',
	JPY: '¥',
	GBP: '£',
	CHF: '₣',
	MXN: '$',
	AUD: '$'
};

export const formatNumber = (n: number): string => {
	return new Intl.NumberFormat('en-US', {
		notation: 'compact'
	}).format(n);
};

export function stringToRandomInteger(inputString: string, maxNumber: number) {
	let hash = 0;
	for (let i = 0; i < inputString.length; i++) {
		hash = (hash << 5) - hash + inputString.charCodeAt(i);
		hash |= 0;
	}

	// Map the hash to a random integer between 1 and 100
	return (Math.abs(hash) % maxNumber) - 1 + 1;
}

export type themeType = 'primary' | 'secondary' | 'tertiary' | 'quaternary'

export function randomId(length:number) {
	let result = '';
	const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
	const charactersLength = characters.length;
	let counter = 0;
	while (counter < length) {
		result += characters.charAt(Math.floor(Math.random() * charactersLength));
		counter += 1;
	}
	return result;
}