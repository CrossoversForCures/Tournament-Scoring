<script lang="ts">
	import {
		Heading,
		P,
		Secondary,
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import { goto } from '$app/navigation';
	import type { PageData } from './$types';
	export let data: PageData;
</script>

<div class="fixed left-0 top-0 h-48 w-full">
	<Heading tag="h1" class="mb-8 mt-8"
		>Welcome to E4E<Secondary class="ms-2">{data.year}</Secondary></Heading
	>
	<Table hoverable={true} divClass="w-1/2 ml-4" color="blue">
		<TableHead>
			<TableHeadCell>Events</TableHeadCell>
		</TableHead>
		<TableBody tableBodyClass="divide-y">
			{#each data.events as event}
				<TableBodyRow on:click={() => goto(`/${event._id}/teams`)}>
					<TableBodyCell>{event.name}</TableBodyCell>
				</TableBodyRow>
			{/each}
		</TableBody>
	</Table>
</div>
<div></div>
