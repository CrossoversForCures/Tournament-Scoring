// src/lib/stores/tabStore.ts
import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import { page } from '$app/stores';
import { derived } from 'svelte/store';

// Define the possible tab values
type TabValue = 'teams' | 'pools' | 'seeding' | 'bracket' | 'results';

// Create a writable store for the active tab
const createActiveTabStore = () => {
    const { subscribe, set } = writable<TabValue>('teams');

    return {
        subscribe,
        set: (value: TabValue) => {
            set(value);
            if (browser) {
                localStorage.setItem('activeTab', value);
            }
        },
        init: () => {
            if (browser) {
                const storedTab = localStorage.getItem('activeTab') as TabValue | null;
                if (storedTab) {
                    set(storedTab);
                }
            }
        }
    };
};

export const activeTab = createActiveTabStore();

// Create a derived store that updates the active tab based on the current route
export const routeTab = derived(page, ($page) => {
    const path = $page.url.pathname;
    const tab = path.split('/').pop() as TabValue;
    if (['teams', 'pools', 'seeding', 'bracket', 'results'].includes(tab)) {
        activeTab.set(tab);
        return tab;
    }
    return 'teams'; // Default to 'teams' if the route doesn't match any tab
});