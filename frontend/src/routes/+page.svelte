<script lang="ts" context="module">
	// parseo de la api
	export interface Nodo {
		Pid: Int32Array;
		Name: string;
		MemUsage: Float32Array;
	}
	export interface Data {
		Node: Nodo;
		Childs: Data[];
	}
</script>

<script lang="ts">
	import Procesos from './procesos.svelte';
	import '@picocss/pico';

	let socket: WebSocket;
	let procecess: Data;

	function connect() {
		socket = new WebSocket('ws://localhost:8080');
		socket.onopen = () => {
			console.log('Connected to WebSocket.');
		};
		socket.onmessage = (event) => {
			procecess = JSON.parse(event.data);
			// console.log(procecess);
		};
		socket.onclose = (event) => {
			console.log(`Closed connection with code ${event.code}.`);
		};

		// traer data cada 400 milisegundos
		setInterval(() => {
			socket.send('Get: ' + Date());
		}, 500);
	}

	function disconnect() {
		socket.close();
	}
</script>

<div data-theme="dark" class="container">
	<nav class="grid">
		<ul>
			<li><button on:click={connect}>Connect</button></li>
			<li class="center"><button on:click={disconnect}>Disconnect</button></li>
		</ul>
	</nav>

	<h2>Procesos activos</h2>
	<div class="scrollable">
		<table class="table">
			<thead class="thead">
				<tr>
					<th class="th">Name</th>
					<th class="th">PID</th>
					<th class="th">MemUsage</th>
				</tr>
			</thead>
			<tbody class="tbody">
				<Procesos data={procecess} />
				<!-- <Datos data={procecess}/> -->
			</tbody>
		</table>
	</div>
</div>

<style>
	.scrollable {
		max-height: 65vh;
		overflow-y: scroll;
	}

	.scrollable:hover {
		-webkit-scrollbar: none;
		-ms-overflow-style: none;
		scrollbar-width: none;
	}
</style>
