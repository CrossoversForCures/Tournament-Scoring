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
	import { isAdmin } from '$lib/stores/admin';
	import { enhance } from '$app/forms';

	export let data: PageData;
</script>

{#if data.seeding == null}
	{#if $isAdmin}
		<form method="POST" action="?/start" use:enhance>
			<Heading tag="h5" class="font-heading ml-2" customSize="text-xl"
				>The elimination round for this division hasn't started yet.
				<button class="link text-theme hover:text-hover" type="submit">Start Elimination?</button>
			</Heading>
		</form>
	{:else}
		<Heading tag="h5" class="font-heading ml-2" customSize="text-xl"
			>The elimination round for this division hasn't started yet. Check back later!</Heading
		>
	{/if}
{:else}
	<Table divClass="ml-2 mr-2 font-default">
		<TableHead class="bg-theme text-white">
			<TableHeadCell>Seed</TableHeadCell>
			<TableHeadCell>Team</TableHeadCell>
			<TableHeadCell>Games Won</TableHeadCell>
			<TableHeadCell>Total Points</TableHeadCell>
		</TableHead>
		<TableBody>
			{#each data.seeding as team}
				<TableBodyRow color="custom">
					<TableBodyCell class="w-1/10 py-2"
						><div class="text-black">{team.seeding}</div></TableBodyCell
					>
					<TableBodyCell class="w-1/6 py-2"><div class="text-black">{team.name}</div></TableBodyCell
					>
					<TableBodyCell class="w-1/10 py-2"
						><div class="text-black">
							{#if team.poolsWon === undefined}
								0
							{:else}
								{team.poolsWon}{/if}
						</div></TableBodyCell
					>
					<TableBodyCell class="w-1/10 py-2"
						><div class="text-black">
							{#if team.totalPoints === undefined}
								0
							{:else}
								{team.totalPoints}{/if}
						</div></TableBodyCell
					>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
{/if}
