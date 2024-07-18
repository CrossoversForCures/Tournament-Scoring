<script lang="ts">
	import { onMount } from 'svelte';

	interface Team {
		team?: string;
		teamId: string;
		seeding?: number;
	}

	interface Node extends Team {
		left?: Node;
		right?: Node;
	}

	interface BracketData {
		event: string;
		root: Node;
	}

	let bracketData: BracketData | null = null;
	let debug: string = '';

	onMount(async () => {
		try {
			const response = await fetch('http://localhost:8000/api/3rd-4th-boys/bracket');
			bracketData = await response.json();
			debug = JSON.stringify(bracketData, null, 2);
		} catch (error) {
			console.error('Error fetching bracket data:', error);
			debug = 'Error fetching data: ' + error;
		}
	});

	function renderNode(node: Node | undefined, depth: number): string {
		if (!node) return '';

		const hasTeam = node.team !== undefined;
		const teamDisplay = hasTeam ? `${node.team} (${node.seeding})` : 'TBD';
		const isBye = !hasTeam && !node.left && !node.right;

		let html = `
			<div class="flex flex-col items-start ${getDepthClass(depth)}">
				<div class="border border-gray-300 p-2 m-1 min-w-[150px] ${isBye ? 'opacity-50' : ''}">
					${teamDisplay}
					<br>
					<small class="text-gray-500">ID: ${node.teamId}</small>
				</div>
				${node.left ? renderNode(node.left, depth + 1) : ''}
				${node.right ? renderNode(node.right, depth + 1) : ''}
			</div>
		`;

		return html;
	}

	function getDepthClass(depth: number): string {
		switch (depth) {
			case 0:
				return 'mt-0';
			case 1:
				return 'mt-12';
			case 2:
				return 'mt-24';
			case 3:
				return 'mt-48';
			case 4:
				return 'mt-96';
			default:
				return 'mt-0';
		}
	}

	function countNodes(node: Node | undefined): number {
		if (!node) return 0;
		return 1 + countNodes(node.left) + countNodes(node.right);
	}
</script>

<main class="p-4">
	<h1 class="mb-4 text-2xl font-bold">{bracketData ? bracketData.event : 'Loading...'}</h1>
	{#if bracketData}
		<div class="flex flex-row-reverse items-center justify-start overflow-x-auto p-4">
			{@html renderNode(bracketData.root, 0)}
		</div>
		<p class="mt-4">Total nodes: {countNodes(bracketData.root)}</p>
	{:else}
		<p>Loading bracket data...</p>
	{/if}
	<pre class="mt-4 overflow-x-auto bg-gray-100 p-4">
			{debug}
		</pre>
</main>
