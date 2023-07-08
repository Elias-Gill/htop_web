<script lang="ts">
	import type { Data } from '../clases.ts';
    import Nodo from './nodo.svelte';
	export let data: Data;
	export let color: number = 0;

	let colorsList: string[] = ['', '#eeeeee', '#faa61a', '#a84aaf', '#1e7251', '#b0917d', '#3f2f0c'];

    function transform(e: Data) {
        e.Childs.forEach((n) => {
            // anadir la data al array
            transform(n)
        });
    }
</script>

<export {color}>
	{#if data}
		<ul>
			<li>
                <Nodo proceso={data.Node} color={colorsList[color]} />
			</li>
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
