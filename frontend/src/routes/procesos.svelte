<script lang="ts">
	import type { Data } from './+page.svelte';
	import Nodo from './nodo.svelte';
	export let data: Data;
	export let color: number = 0;

	let colorsList: string[] = ['', '#eeeeee', '#faa61a', '#a84aaf', '#1e7251', '#b0917d', '#3f2f0c'];
	console.log(color);
</script>

<export {color}>
	{#if data}
		<ul>
			<li>
				{#if data.Node}
					<Nodo proceso={data.Node} color={colorsList[color]} />
				{/if}
			</li>
			{#each data.Childs as process}
				{#if process.Childs}
					<svelte:self data={process} color={color + 1} />
				{/if}
			{/each}
		</ul>
	{/if}
</export>

<style>
	ul {
		padding: 0.2em 0 0 0.5em;
		margin: 0 0 0 0.5em;
		/* border-left: 1px solid #424242; */
	}

	li {
        list-style-type: none !important;
		padding: 0.2em 0;
	}
</style>
