<script lang="ts">
	import { goto } from '$app/navigation';
	import { Tabs, TabItem } from 'flowbite-svelte';
	import { activeTab, routeTab } from '$lib/stores/tabStore';
	import { onMount } from 'svelte';

	onMount(() => {
		activeTab.init();
	});

	function handleTabClick(tab: string) {
		goto(`./${tab}`);
		activeTab.set(tab as 'teams' | 'pools' | 'seeding' | 'bracket' | 'results');
	}
</script>

<Tabs
	class="font-heading ml-2"
	contentClass=""
	activeClasses="p-4 text-theme border-b-2 border-theme"
>
	<TabItem open={$routeTab === 'teams'} title="Teams" on:click={() => handleTabClick('teams')}
	></TabItem>
	<TabItem open={$routeTab === 'pools'} title="Pools" on:click={() => handleTabClick('pools')}
	></TabItem>
	<TabItem open={$routeTab === 'seeding'} title="Seeding" on:click={() => handleTabClick('seeding')}
	></TabItem>
	<TabItem open={$routeTab === 'bracket'} title="Bracket" on:click={() => handleTabClick('bracket')}
	></TabItem>
	<TabItem open={$routeTab === 'results'} title="Results" on:click={() => handleTabClick('results')}
	></TabItem>
</Tabs>

<slot />
