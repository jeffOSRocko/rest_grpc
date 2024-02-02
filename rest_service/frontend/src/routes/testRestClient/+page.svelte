<script lang="ts">
	import AddIdVal from '$lib/components/AddIDVal.svelte';
	import GetAllIdVals from '$lib/components/GetAllIDVals.svelte';
	import GetIdVal from '$lib/components/GetIDVal.svelte';
	import ModifyIdVal from '$lib/components/ModifyIDVal.svelte';
	import DeleteIdVal from '$lib/components/DeleteIDVal.svelte';
	import ToggleableControl from '$lib/components/ToggleableControl.svelte';
	import { onMount } from 'svelte';

	const apiUrl = import.meta.env.VITE_API_URL;

	let GetAllIdValsId: string = 'a';
	let GetIdValId: string = 'b';
	let ModifyIdValId: string = 'c';

	onMount(async () => {
		const GetAllIdValsIdModule = (await import('../../lib/components/GetAllIDVals.svelte')).default;
		GetAllIdValsIdModule.bind(GetAllIdValsId);

		const GetIdValIdModule = (await import('../../lib/components/GetIDVal.svelte')).default;
		GetIdValIdModule.bind(GetIdValId);

		const ModifyIdValIdModule = (await import('../../lib/components/ModifyIDVal.svelte')).default;
		ModifyIdValIdModule.bind(ModifyIdValId);
	});
</script>

<svelte:head>
	<title>Test Rest Service</title>
	<meta name="description" content="Tests the endpoints of the Rest service" />
</svelte:head>

<main>
	<div class="text-column">
		<h1>Test Rest Client</h1>

		<ToggleableControl title="Add ID Val">
			<AddIdVal />
		</ToggleableControl>
		<ToggleableControl title="Get All ID Vals">
			<GetAllIdVals />
		</ToggleableControl>
		<ToggleableControl title="Get ID Val">
			<GetIdVal />
		</ToggleableControl>
		<ToggleableControl title="Modify ID Val">
			<ModifyIdVal />
		</ToggleableControl>
		<ToggleableControl title="Delete ID Val">
			<DeleteIdVal />
		</ToggleableControl>
	</div>
</main>

<style>
</style>
