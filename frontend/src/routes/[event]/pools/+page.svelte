<script lang="ts">
	import {
		Heading,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Input,
		Button,
		Modal,
		Label
	} from 'flowbite-svelte';

	import { writable, type Writable } from 'svelte/store';
	import { EditOutline } from 'flowbite-svelte-icons';
	import { enhance } from '$app/forms';
	import type { PageData } from './$types';
	import { isAdmin } from '$lib/stores/admin';

	export let data: PageData;
	export let formModal = false;

	const editingGame: Writable<string | null> = writable(null);

	function openModal(gameId: string) {
		editingGame.set(gameId);
		formModal = true;
	}

	function closeModal() {
		editingGame.set(null);
		formModal = false;
	}
</script>

{#if data.games == null}
	{#if $isAdmin}
		<form method="POST" action="?/start" use:enhance>
			<Heading tag="h5" class="font-heading ml-2" customSize="text-xl"
				>This division hasn't started yet.
				<button class="link text-theme hover:text-hover" type="submit">Start Pools?</button>
			</Heading>
		</form>
	{:else}
		<Heading tag="h5" class="font-heading ml-2" customSize="text-xl"
			>This division hasn't started yet. Check back later!</Heading
		>
	{/if}
{:else}
	{#each Object.keys(data.games) as round}
		<Heading tag="h5" class="font-heading ml-2" customSize="text-xl">Round {round}</Heading>
		<Table divClass="ml-2 mr-2 font-default">
			<TableHead class="bg-theme text-white">
				<TableHeadCell>Court</TableHeadCell>
				<TableHeadCell>Team 1</TableHeadCell>
				<TableHeadCell>Score</TableHeadCell>
				<TableHeadCell>Team 2</TableHeadCell>
				<TableHeadCell>Score</TableHeadCell>
				{#if $isAdmin}
					<TableHeadCell>Update</TableHeadCell>
				{/if}
			</TableHead>
			<TableBody>
				{#each data.games[round] as game}
					<TableBodyRow color="default">
						<TableBodyCell class="py-2">
							<div class="text-black">Court {game.court}</div></TableBodyCell
						>
						<TableBodyCell class="py-2"
							><div class="text-black">{game.team1Name}</div></TableBodyCell
						>
						<TableBodyCell class="py-2">
							<div class="text-black">
								{game.team1Score === undefined ? '' : game.team1Score}
							</div>
						</TableBodyCell>
						<TableBodyCell class="py-2"
							><div class="text-black">{game.team2Name}</div></TableBodyCell
						>
						<TableBodyCell class="py-2"
							><div class="text-black">
								{game.team2Score === undefined ? '' : game.team2Score}
							</div></TableBodyCell
						>
						{#if $isAdmin}
							<TableBodyCell class="py-0">
								<button
									on:click={() => {
										openModal(game._id);
									}}
								>
									<EditOutline class="text-theme h-7 w-7 content-center " />
								</button>
							</TableBodyCell>
						{/if}

						<Modal
							bind:open={formModal}
							size="xs"
							autoclose={false}
							class="w-full"
							backdropClass="fixed inset-0 z-40 bg-gray-900 bg-opacity-50"
						>
							<form
								class="flex flex-col space-y-6"
								method="POST"
								action="?/update"
								use:enhance
								on:submit={closeModal}
							>
								<input type="hidden" name="gameId" value={$editingGame} />
								<Label class="active:border-theme space-y-2">
									<span>Team 1 Score</span>
									<Input type="number" name="team1Score" required />
								</Label>
								<Label class="space-y-2">
									<span>Team 2 Score</span>
									<Input type="number" name="team2Score" required />
								</Label>
								<Button type="submit" class="w-full1 bg-theme hover:bg-hover">Confirm</Button>
							</form>
						</Modal>
					</TableBodyRow>
				{/each}
			</TableBody>
		</Table>
		<br />
	{/each}
{/if}
