<script>
	import Header from './component/header.svelte'
	import LogViewer from './component/logviewer.svelte'
</script>

<svelte:head>
	<title>Draught Log</title>
</svelte:head>

<style>
	.brand-logo {
		margin-left: 10px;
	}
</style>

<Header />

<LogViewer />