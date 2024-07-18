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

{#each Object.keys(data.games) as round}
	<Heading tag="h3">Round {round}</Heading>
	<br />
	<Table hoverable={true} divClass="ml-2 mr-2" color="blue">
		<TableHead>
			<TableHeadCell>Court</TableHeadCell>
			<TableHeadCell>Team 1</TableHeadCell>
			<TableHeadCell>Team 2</TableHeadCell>
		</TableHead>
		<TableBody tableBodyClass="divide-y">
			{#each data.games[round] as game}
				<TableBodyRow color="custom">
					<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
						><div class="text-black">Court {game.court}</div></TableBodyCell
					>
					<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
						><div class="text-black">{game.team1Name}</div></TableBodyCell
					>
					<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
						><div class="text-black">{game.team2Name}</div></TableBodyCell
					>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
{/each}
