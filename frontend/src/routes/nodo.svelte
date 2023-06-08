<script lang="ts">
	import type { Nodo } from './+page.svelte';

	export let color: string;
	console.log(color);
	export let proceso: Nodo;
</script>

<export>
	<tr style="--color:{color}">
		<td class="td"> {proceso.Name}</td>
		<td class="td">{proceso.Pid}</td>
		<td class="td">{proceso.MemUsage}</td>
	</tr>
</export>

<style>
	span {
		padding: 0 0 0 1.5em;
		color: var(--color);
		background-size: 1em 1em;
	}
</style>
