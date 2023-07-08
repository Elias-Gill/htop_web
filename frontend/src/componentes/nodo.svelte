<script lang="ts">
	import type { Nodo_api } from '../clases.ts';

	export let color: string;
	export let proceso: Nodo_api;
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
