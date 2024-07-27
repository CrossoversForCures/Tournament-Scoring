<script lang="ts">
	import { Dropdown, DropdownItem, DropdownDivider } from 'flowbite-svelte';
	import { EditOutline } from 'flowbite-svelte-icons';

	export let node: BracketNode | undefined;
	export let depth: number = 0;
	export let isSecondTeam: boolean = false;

	function submitForm(teamId: string | undefined) {
		const form = document.querySelector('form');
		if (form instanceof HTMLFormElement) {
			const input = form.querySelector('input[name="teamId"]');
			if (input instanceof HTMLInputElement) {
				input.value = teamId ?? '';
				form.submit();
			} else {
				console.error("Couldn't find input element with name 'teamId'");
			}
		} else {
			console.error("Couldn't find form element");
		}
	}

	interface BracketNode {
		team?: string;
		teamId: string;
		seeding?: number;
		left?: BracketNode;
		right?: BracketNode;
		court?: string;
	}
</script>

{#if node}
	{#if node.left && node.right}
		<div class="flex {isSecondTeam ? 'mt-16' : ''}">
			<div>
				<svelte:self node={node.left} depth={depth + 1} isSecondTeam={false} />
				<svelte:self node={node.right} depth={depth + 1} isSecondTeam={true} />
			</div>
			<div class="relative flex w-24 flex-col">
				<!-- Vertical Line-->
				<div class="absolute bottom-1/4 left-[-1px] top-1/4 w-px bg-black"></div>
				<!-- Horizontal Line-->
				<div class="absolute left-0 right-0 top-1/2 h-px bg-black"></div>
				<!-- Court-->
				<div
					class="font-default absolute top-1/2 ml-3 mt-1 text-xs font-bold {node.court?.includes(
						'-'
					)
						? 'text-heading'
						: 'text-red-500'}"
				>
					<!-- svelte-ignore empty-block -->
					{#if node.court === 'N/A' || node.left.team === undefined || node.right.team === undefined}{:else if node.court === undefined}
						Court TBD
					{:else}
						{#if node.court.includes('-')}
							Court {node.court.substring(1)}
						{:else}
							Court {node.court}
						{/if}
					{/if}
				</div>
			</div>
			<div class="flex items-center">
				<div class="w-48 border border-black p-2">
					<span
						class="font-default font-bold {node.team === 'BYE' || node.team === undefined
							? 'text-heading opacity-50'
							: 'text-theme'}"
					>
						{node.seeding ? `(${node.seeding})` : ''}
						{node.team || 'TBD'}
					</span>
					{#if node.court != 'N/A' && node.court != undefined && node.team === undefined}
						<button>
							<EditOutline class="text-theme ml-2 h-7 w-7 content-center" />
						</button>
						<Dropdown>
							<form method="POST" action="?/update">
								<input type="hidden" name="teamId" value="" />
								<DropdownItem on:click={() => submitForm(node.left?.teamId)}
									>({node.left.seeding}) {node.left.team}</DropdownItem
								>
								<DropdownDivider class="bg-heading" />
								<DropdownItem on:click={() => submitForm(node.right?.teamId)}
									>({node.right.seeding}) {node.right.team}</DropdownItem
								>
							</form>
						</Dropdown>
					{/if}
				</div>
			</div>
		</div>
	{:else}
		<!-- First round -->
		<div class="w-48 border border-black p-2 {isSecondTeam && depth != 0 ? 'border-t-0' : ''}">
			<span
				class="font-default font-bold {node.team?.toUpperCase() === 'BYE'
					? 'text-black opacity-50'
					: 'text-theme'}"
			>
				{node.seeding ? `(${node.seeding})` : ''}
				{node.team || 'TBD'}
			</span>
		</div>
	{/if}
{/if}
