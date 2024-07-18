<script lang="ts">
	import {
		Heading,
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

<Table hoverable={true} divClass="ml-2 mr-2" color="blue">
	<TableHead>
		<TableHeadCell>Seed</TableHeadCell>
		<TableHeadCell>Team</TableHeadCell>
		<TableHeadCell>Games Won</TableHeadCell>
		<TableHeadCell>Total Points</TableHeadCell>
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each data.seeding as team}
			<TableBodyRow color="custom">
				<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
					><div class="text-black">{team.seeding}</div></TableBodyCell
				>
				<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
					><div class="text-black">{team.name}</div></TableBodyCell
				>
				<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
					><div class="text-black">{team.poolsWon}</div></TableBodyCell
				>
				<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
					><div class="text-black">{team.totalPoints}</div></TableBodyCell
				>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>
