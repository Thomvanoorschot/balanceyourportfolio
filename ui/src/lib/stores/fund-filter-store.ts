import {type Updater, type Writable, writable, get} from "svelte/store";
import type {FundHolding, FundHoldingsFilter} from "$lib/fund";

type WritableHoldingsStore = {
    holdings: FundHolding[];
    filter: FundHoldingsFilter;
};


export interface HoldingsStore extends Writable<WritableHoldingsStore> {
    filter: (f: Partial<FundHoldingsFilter>) => Promise<void>;
    nextPage: () => Promise<void>;
}

export function createHoldingsStore(fundId: string) :HoldingsStore{
    const store = writable<WritableHoldingsStore>({
        holdings: [],
        filter: {
            fundId: fundId,
            sectorName: "Any sector",
            searchTerm: "",
            limit: 20,
            offset: 0,
        },
    });

    async function nextPage(): Promise<void> {
        const writableHoldingsStore = get(store)
        const fhFilter = writableHoldingsStore.filter
        fhFilter.offset = writableHoldingsStore.holdings.length
        const holdingsResult = await fetch(`http://localhost:8080/api/v1/fund/holdings/filter`, {
            method: "POST",
            body: JSON.stringify(fhFilter)
        });
        const newHoldings: FundHolding[] = await holdingsResult.json();
        store.update((store) => ({
            ...store,
            filter: fhFilter,
            holdings: [...store.holdings, ...newHoldings]
        }))
    }
    async function filter(f: Partial<FundHoldingsFilter>): Promise<void> {
        const writableHoldingsStore = get(store)
        const fhFilter = {
            ...writableHoldingsStore.filter,
            ...f,
            limit: 20,
            offset: 0
        }
        const holdingsResult = await fetch(`http://localhost:8080/api/v1/fund/holdings/filter`, {
            method: "POST",
            body: JSON.stringify(fhFilter)
        });
        const newHoldings: FundHolding[] = await holdingsResult.json();
        store.update((store) => ({
            ...store,
            filter: fhFilter,
            holdings: [...newHoldings]
        }))
    }

    return {
        ...store,
        filter,
        nextPage,
    }
}