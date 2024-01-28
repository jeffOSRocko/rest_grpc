<script lang="ts">
	export let title = '';
	let expanded = false;

	const uuid = '8f4b819f-ec97-4182-b043-f933109cd842';

	function toggleExpanded(event: Event) {
		// Check if the click event originated from the child component
		const target = event.target as Element;

		console.log('target: ' + target?.id);

		// Check if the click event originated from the child component
		if (target?.id !== uuid) {
			// Click originated from the child component, ignore it
			return;
		}

		expanded = !expanded;
		console.log('expanded: ' + expanded);
	}

	function handleKeydown(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			toggleExpanded(event);
		}
	}
</script>

<div on:click={toggleExpanded} on:keydown={handleKeydown} class="toggle-container">
	<div class="header">
		<span id={uuid} class="triangle {expanded ? 'triangle-expanded' : 'triangle-collapsed'}" />
		<span class="title-text">{title}</span>
	</div>
	<div class={expanded ? 'content content-expanded' : 'content'}>
		<slot />
		<!-- Slot for the actual control content -->
	</div>
</div>

<style>
	.toggle-container {
		cursor: pointer;
		/* Additional styles for the container */
	}
	.header {
		display: flex;
		align-items: center; /* Vertically center the child elements */
	}
	.triangle {
		display: inline-block;
		width: 0;
		height: 0;
		border-left: 10px solid transparent; /* Increase size */
		border-right: 10px solid transparent; /* Increase size */
		border-top: 10px solid black; /* Increase size */
		transition: transform 0.3s ease;
		cursor: pointer; /* Add cursor pointer for better UX */
		margin-right: 10px; /* Add space between the triangle and the title */
		margin-bottom: 20px; /* Increase space between triangles */
	}
	.triangle-expanded {
		transform: rotate(0deg);
	}
	.triangle-collapsed {
		transform: rotate(-90deg);
	}
	.title-text {
		margin-top: -20px; /* Slightly raise the text */
		/* Alternatively, you can adjust line-height instead of margin-top */
		/* line-height: 0.9; */
	}
	.content {
		margin: 10px 0;
		display: none;
	}
	.content-expanded {
		display: block;
	}
</style>
