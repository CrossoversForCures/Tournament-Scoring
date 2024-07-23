<script lang="ts">
	import { goto } from '$app/navigation';
	import { Tabs, TabItem } from 'flowbite-svelte';
</script>

<Tabs
	class="font-heading ml-2"
	contentClass=""
	activeClasses="p-4 text-theme border-b-2 border-theme"
>
	<TabItem open title="Teams" on:click={() => goto('./teams')}></TabItem>
	<TabItem title="Pools" on:click={() => goto('./pools')}></TabItem>
	<TabItem title="Seeding" on:click={() => goto('./seeding')}></TabItem>
	<TabItem title="Bracket" on:click={() => goto('./bracket')}></TabItem>
	<TabItem title="Results" on:click={() => goto('./results')}></TabItem></Tabs
>
<br />
<slot />
