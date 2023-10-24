import type {PageLoad} from './$types';
import type {FundDetails} from "$lib/fund";

export const load = (async ({fetch, params, url}) => {
    const fundId = url.searchParams.get("fundId")
    const fetchDetails = async (): Promise<FundDetails> => {
        const fundDetails = await fetch(`http://localhost:8080/api/v1/fund/${fundId}/details`);
        return await fundDetails.json();
    };
    return {
        details: fetchDetails()
    };
}) satisfies PageLoad;