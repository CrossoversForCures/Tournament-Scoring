<script lang="ts">
	import { enhance } from '$app/forms';
	import { isAdmin } from '$lib/stores/admin';
	import { Heading } from 'flowbite-svelte';
	import BracketNode from './BracketNode.svelte';

	import type { PageData } from './$types';
	export let data: PageData;
</script>

{#if data.root == null}
	{#if $isAdmin}
		<form method="POST" action="?/start" use:enhance>
			<Heading tag="h5" class="font-heading ml-2" customSize="text-xl">
				The elimination round for this division hasn't started yet.
				<button class="link text-theme hover:text-hover" type="submit">Start Elimination?</button>
			</Heading>
		</form>
	{:else}
		<Heading tag="h5" class="font-heading ml-2" customSize="text-xl">
			The elimination round for this division hasn't started yet. Check back later!
		</Heading>
	{/if}
{:else}
	<div class="overflow-x-auto p-8">
		<div class="inline-block">
			<BracketNode node={data.root} />
		</div>
	</div>
{/if}
