<!-- GetAllIdVals.svelte -->
<script lang="ts">
	import { onMount } from 'svelte';

	// @ts-ignore
	import { writable } from 'svelte/store';

	export const GetAllIdValsId: string = 'getAllIdVals';
	export let ShowTitle = writable<boolean>(false);
	export let responseData = writable<any>({}); // Make responseData a writable store

	async function handleGet() {
		try {
			const response = await fetch(`http://localhost:8080/getallidvals`);
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}
			responseData.set(await response.json());
		} catch (error: any) {
			console.error('Error fetching data:', error);
			responseData.set({ error: error.message || 'unknown error' });
		}
	}

	onMount(() => {
		handleGet();
	});
</script>

<div class="container">
	<div class="header">
		{#if $ShowTitle}
			<span>Get All ID Vals</span>
		{/if}
		<button class="refresh-button" on:click={handleGet}>Refresh</button>
	</div>
	<textarea class="json-textarea" readonly>{JSON.stringify($responseData, null, 2)}</textarea>
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		align-items: stretch;
		max-width: 100%;
	}

	.header {
		display: flex;
		justify-content: space-between; /* Aligns label to the left and button to the right */
		margin-bottom: 10px; /* Spacing between header and textarea */
	}

	.refresh-button {
		margin-left: 10px;
		align-self: right;
	}

	.json-textarea {
		width: 100%;
		min-height: 100px;
		resize: both;
		border: 1px solid #ccc;
		border-radius: 4px;
		margin-bottom: 10px; /* Spacing between textarea and button */
		padding: 10px;
		font-family: monospace;
		background-color: #f8f8f8;
	}
</style>
