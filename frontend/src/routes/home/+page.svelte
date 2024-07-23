<script lang="ts">
	import {
		Heading,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		Tabs,
		TabItem,
		Modal,
		Input,
		Label,
		Helper,
		Button
	} from 'flowbite-svelte';
	import { goto } from '$app/navigation';
	import type { PageData } from './$types';
	import type { ActionData } from './$types';
	import { enhance } from '$app/forms';

	export let formModal = false;

	export let form: ActionData;
	export let data: PageData;
</script>

<Tabs
	class="font-heading ml-2"
	contentClass=""
	activeClasses="p-4 text-theme border-b-2 border-theme"
>
	<TabItem open title="Events" on:click={() => goto(`./home`)}></TabItem>
	<TabItem title="Admin" on:click={() => (formModal = true)}></TabItem>
</Tabs>
<br />
<Heading tag="h3" class="font-heading ml-2" customSize="text-xl">Event Schedule</Heading>
<br />
<Table hoverable={true} divClass="ml-2 mr-2 !bg-theme font-default" color="custom">
	<TableHead>
		<TableHeadCell class="text-white">Time</TableHeadCell>
		<TableHeadCell class="text-white">Division</TableHeadCell>
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each data.events as event}
			<TableBodyRow color="default" on:click={() => goto(`/${event.slug}/teams`)}>
				<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
					><div class="text-black">{event.time}</div></TableBodyCell
				>
				<TableBodyCell tdClass="px-6 py-2 whitespace-nowrap font-medium"
					><div class="text-black">{event.name}</div></TableBodyCell
				>
			</TableBodyRow>
		{/each}
	</TableBody>
</Table>

<Modal bind:open={formModal} size="xs" autoclose={false} class="w-full">
	<form class="flex flex-col space-y-6" method="POST" use:enhance>
		<Label class="active:border-theme space-y-2">
			<span>Username</span>
			<Input type="text" name="username" required />
		</Label>
		<Label class="space-y-2">
			<span>Your password</span>
			<Input type="password" name="password" required />
		</Label>

		{#if form?.incorrect}
			<Helper class="mt-2" color="red">Incorrect Credentials.</Helper>
		{/if}

		<Button type="submit" class="w-full1 bg-theme hover:bg-hover">Login</Button>
	</form>
</Modal>
