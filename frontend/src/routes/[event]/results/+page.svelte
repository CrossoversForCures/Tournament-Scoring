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

{#if data.results == null}
	<Heading tag="h5" class="font-heading ml-2" customSize="text-xl"
		>This division hasn't finished yet. Check back later!</Heading
	>
{:else}
	<Table divClass="ml-2 mr-2 font-default">
		<TableHead class="bg-theme text-white">
			<TableHeadCell>Rank</TableHeadCell>
			<TableHeadCell>Name</TableHeadCell>
		</TableHead>
		<TableBody tableBodyClass="divide-y">
			{#each data.results as team}
				<TableBodyRow
					class={team.rank == 1
						? 'bg-gold opacity-90'
						: team.rank == 2
							? 'bg-silver opacity-90'
							: team.rank == 3
								? 'bg-bronze opacity-90'
								: ''}
				>
					<TableBodyCell class="py-2"
						><div
							class="text-lg {team.rank === 2 || team.rank === 3 ? 'text-white' : 'text-black'}"
						>
							{#if team.rank === 3}
								{team.rank}T
							{:else}
								{team.rank}
							{/if}
						</div></TableBodyCell
					>
					<TableBodyCell class="py-2"
						><div
							class="text-lg {team.rank === 2 || team.rank === 3 ? 'text-white' : 'text-black'}"
						>
							{team.name}
						</div></TableBodyCell
					>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
{/if}
