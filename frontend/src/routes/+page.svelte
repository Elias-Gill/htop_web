<script lang="ts">
    import Procesos from "./procesos.svelte";

	// parseo de la api
	interface Nodo {
		Pid: Int32Array;
		Name: string;
		MemUsage: Float32Array;
	}
	interface Data {
		Node: Nodo;
		Childs: Nodo[];
	}

	let socket: WebSocket;
	let procecess: Data;

	function connect() {
		socket = new WebSocket('ws://localhost:8080');
		socket.onopen = () => {
			console.log('Connected to WebSocket.');
		};
		socket.onmessage = (event) => {
			procecess = JSON.parse(event.data);
			console.log(procecess);
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

<button on:click={connect}>Connect</button>
<button on:click={disconnect}>Disconnect</button>

<h1>Procesos activos</h1>
<Procesos>

</Procesos>
