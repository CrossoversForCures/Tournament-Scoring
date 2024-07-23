<script lang="ts">
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

	import type { PageData } from './$types';
	export let data: PageData;
</script>

<Table divClass="ml-2 mr-2 font-default">
	<TableHead class="bg-theme text-white">
		<TableHeadCell>Name</TableHeadCell>
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each data.teams as team}
			<TableBodyRow color="default">
				<TableBodyCell class="py-2"><div class="text-black">{team.name}</div></TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
