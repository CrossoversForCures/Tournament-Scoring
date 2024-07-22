<script lang="ts">
	import {
		Heading,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Button
	} from 'flowbite-svelte';

	import type { PageData } from './$types';
	import { isAdmin } from '$lib/stores/admin';
	export let data: PageData;
</script>

{#if data.seeding == null}
	{#if isAdmin}
		<form action="?/start" class="justify-center">
			<Button type="submit" class="bg-theme hover:bg-hover">Start Elimination Round</Button>
		</form>
	{:else}
		<Heading tag="h5" class="font-heading ml-2" customSize="text-xl"
			>The elimination round for this division hasn't started yet. Check back later!</Heading
		>
	{/if}
{:else}
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
{/if}
