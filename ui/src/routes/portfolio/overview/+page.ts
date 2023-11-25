import type {PageLoad} from './$types';
import type {Portfolio} from "$lib/portfolio";

export const load = (async ({fetch, params, url}) => {
    const fetchPortfolios = async (): Promise<Portfolio[]> => {
        const portfolios = await fetch(`http://localhost:8080/api/v1/portfolio/`);
        return await portfolios.json();
    };
    return {
        portfolios: fetchPortfolios()
    };
}) satisfies PageLoad;