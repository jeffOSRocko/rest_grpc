<script lang="ts">
	import { writable } from 'svelte/store';

	let key: string = '';
	let val: string = '';
	const responseText = writable<string>('');

	async function addIdVal() {
		try {
			const response = await fetch('http://localhost:8080/addidval', {
				method: 'PUT',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ key, val })
			});
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}
			const responseData = await response.text();
			responseText.set(responseData);
		} catch (err: any) {
			console.error('Error:', err);
			responseText.set(err.message);
		}
	}

	$: isButtonDisabled = !(key && val);
</script>

<form on:submit={addIdVal}>
	<input type="text" bind:value={key} placeholder="Key" />
	<input type="text" bind:value={val} placeholder="Value" />
	<button type="submit" disabled={isButtonDisabled}>Add</button>

	<textarea class="resizeable-textarea" readonly>{$responseText}</textarea>
</form>

<style>
	.resizeable-textarea {
		width: 100%;
		min-height: 100px;
		resize: horizontal;
		border: 1px solid #ccc;
		border-radius: 4px;
		padding: 10px;
		font-family: monospace;
		background-color: #f8f8f8;
	}
</style>
