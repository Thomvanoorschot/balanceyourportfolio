import type {PageLoad} from './$types';
import type {PortfolioDetails} from "$lib/portfolio";

export const load = (async ({fetch, params, url}) => {
    const fetchPortfolioDetails = async (): Promise<PortfolioDetails> => {
        const details = await fetch(`http://localhost:8080/api/v1/portfolio/${params.id}`);
        return await details.json();
    };
    return {
        details: fetchPortfolioDetails()
    };
}) satisfies PageLoad;