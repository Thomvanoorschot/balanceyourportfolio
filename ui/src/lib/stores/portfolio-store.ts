import {type Writable, writable} from "svelte/store";
import type {Portfolio, PortfolioListItem} from "$lib/portfolio";
import {EMPTY_UUID} from "$lib/utils";

type WritablePortfolioStore = {
    portfolios: Portfolio[];
};

export interface PortfolioStore extends Writable<WritablePortfolioStore> {
    upsertPortfolio: (p: Portfolio) => Promise<void>;
    addEmptyItem: (p: Portfolio) =>  void;
}

function emptyPortfolio(): Portfolio {
    return {
        id: EMPTY_UUID,
        name: "",
        items: [emptyItem()],
    }
}
function emptyItem():PortfolioListItem{
    return {
        id: EMPTY_UUID,
        name: "",
        fundId: EMPTY_UUID,
        amount: undefined,
    }
}

export function createPortfoliosStore(portfolios: Portfolio[]): PortfolioStore {
    portfolios.push(emptyPortfolio())
    portfolios.forEach(p => p.items.push(emptyItem()))
    const store = writable<WritablePortfolioStore>({
        portfolios: portfolios,
    });

    async function upsertPortfolio(p: Portfolio): Promise<void> {
        const res = await fetch(
            `http://localhost:8080/api/v1/portfolio/`,
            {
                method: "PUT",
                body: JSON.stringify(p)
            },
        );
        const upsertedPortfolio: Portfolio = await res.json()
        store.update(store => {
            for (let portfolio of store.portfolios) {
                if (portfolio.id === upsertedPortfolio.id){
                    portfolio = upsertedPortfolio
                    return store
                }
            }
            store.portfolios = [...store.portfolios, upsertedPortfolio]
            return store
        })

    }
    function addEmptyItem(p :Portfolio):void {
        p.items.push(emptyItem())
        store.update(value => value)
    }

    return {
        ...store,
        upsertPortfolio,
        addEmptyItem,
    }
}