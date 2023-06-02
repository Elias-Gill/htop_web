<script lang="ts">
	interface Nodo {
		pid: number;
	}
	interface Data {
		nodo: Nodo;
		hijos: Nodo[];
	}

	let socket: WebSocket;
	let message: Data;

	function connect() {
		socket = new WebSocket('ws://localhost:8080');

		socket.onopen = () => {
			console.log('Connected to WebSocket.');
		};

		socket.onmessage = (event) => {
			message = JSON.parse(event.data);
			console.log(message);
		};

		socket.onclose = (event) => {
			console.log(`Closed connection with code ${event.code}.`);
		};
	}

	function send(message: string) {
		socket.send(message);
	}

	function disconnect() {
		socket.close();
	}
</script>

<button on:click={connect}>Connect</button>
<button on:click={() => send('Hello, server!')}>Send</button>
<button on:click={disconnect}>Disconnect</button>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
