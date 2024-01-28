<!-- GetIDVal.svelte -->
<script lang="ts">
	// @ts-ignore
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	export const GetIdValId = 'getIdVal';
	export let responseData = writable('ham'); // Make responseData a writable store

	let id: string = '';

	async function handleGet() {
		try {
			const response = await fetch(`http://localhost:8080/getidval/${id}`);
			if (!response.ok) {
				responseData.set('');
				throw new Error(`HTTP error! status: ${response.status}`);
			}
			const foo = await response.text();
			responseData.set(foo);
			console.log('responseData: ' + $responseData);
		} catch (error: any) {
			console.error('Error fetching data:', error);
			responseData.set(error.message || 'unknown error');
		}
	}

	onMount(() => {
		responseData.set('Enter the ID and press Get!');
	});
</script>

<div class="container">
	<form on:submit={handleGet}>
		<div class="input-group">
			<label for="idInput">ID:</label>
			<input id="idInput" type="string" bind:value={id} />
			<button class={'get-button'} type="submit">Get</button>
		</div>
		<textarea class="json-textarea" readonly>{$responseData}</textarea>
	</form>
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		align-items: stretch;
		max-width: 100%;
	}
	.get-button {
		margin-left: 10px;
		align-self: right;
	}
	.input-group {
		display: flex;
		align-items: right;
		margin-bottom: 10px;
	}

	.input-group label {
		margin-right: 10px;
	}

	.json-textarea {
		width: 100%;
		resize: horizontal; /* Allows resizing both horizontally and vertically */
		border: 1px solid #ccc;
		border-radius: 4px;
		padding: 10px;
		font-family: monospace;
		background-color: #f8f8f8;
	}
</style>
