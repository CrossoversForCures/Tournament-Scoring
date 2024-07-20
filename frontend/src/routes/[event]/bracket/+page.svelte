<!-- src/routes/+page.svelte -->
<script lang="ts">
	import { onMount } from 'svelte';

	interface Team {
		team?: string;
		teamId: string;
		seeding?: number;
	}

	interface BracketNode extends Team {
		left?: BracketNode;
		right?: BracketNode;
		court?: string;
	}

	interface BracketData {
		event: string;
		root: BracketNode;
	}
	import type { PageData } from './$types';
	export let data: PageData;

	function renderBracket(
		node: BracketNode | undefined,
		depth: number = 0,
		isSecondTeam: boolean = false
	): string {
		if (!node) return '';

		const team = node.team || 'TBD';
		const seed = node.seeding ? `(${node.seeding})` : '';
		const isBye = team.toUpperCase() === 'BYE';
		let html = `
      <div class="w-64 p-2 border border-black ${isSecondTeam ? (depth == data.rounds ? 'border-t-0' : '') : ''}" >
        <span class="font-bold ${isBye ? 'text-black opacity-50' : 'text-blue-600'}">${seed} ${team}</span>
      </div>
  `;

		if (node.left || node.right) {
			const spacingClass = isSecondTeam ? 'mt-16' : '';
			html = `
        <div class="flex ${spacingClass}">
          <div>
            ${renderBracket(node.left, depth + 1, false)}
            ${renderBracket(node.right, depth + 1, true)}
          </div>
          <div class="flex flex-col relative w-24">
            <div class="absolute left-[-1px] top-1/4 bottom-1/4 w-px bg-black"></div>
        	<div class="absolute left-0 right-0 top-1/2 h-px bg-black"></div>
			<div class="absolute top-1/2 ml-3 mt-1 text-xs font-bold text-red-500">${node.court === 'N/A' ? '' : node.court === undefined ? 'Court TBD' : `Court ${node.court}`}</div>
          </div>
          <div class="flex items-center">
            ${html}
          </div>
        </div>
      `;
		}

		return html;
	}
</script>

<div class="overflow-x-auto p-8">
	<div class="inline-block">
		{@html renderBracket(data.root)}
	</div>
</div>
