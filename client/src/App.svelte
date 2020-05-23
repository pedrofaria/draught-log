<script>
	import { Router } from 'svelte-router-spa'
	import { routes } from './routes'
</script>

<svelte:head>
	<title>Draught Log</title>
</svelte:head>

<Router {routes} />