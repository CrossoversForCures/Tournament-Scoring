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

<Table hoverable={true} divClass="ml-2 mr-2 font-default">
	<TableHead class="bg-theme text-white">
		<TableHeadCell>Name</TableHeadCell>
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each data.teams as team}
			<TableBodyRow color="custom">
				<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
					><div class="text-black">{team.name}</div></TableBodyCell
				>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
