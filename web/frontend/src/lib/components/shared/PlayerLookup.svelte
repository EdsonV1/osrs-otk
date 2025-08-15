<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { PlayerStats } from '$lib/types';

	const dispatch = createEventDispatcher<{
		statsLoaded: PlayerStats;
		error: string;
	}>();

	export let username: string = '';
	export let loading: boolean = false;
	export let placeholder: string = 'Enter OSRS username...';
	export let buttonText: string = 'Lookup Player';

	let error: string = '';

	async function lookupPlayer() {
		if (!username.trim()) {
			error = 'Please enter a username';
			return;
		}

		loading = true;
		error = '';

		try {
			const response = await fetch(`http://localhost:8080/api/player-stats/${encodeURIComponent(username.trim())}`);
			const data = await response.json();

			if (!data.success) {
				error = data.error || 'Failed to fetch player stats';
				dispatch('error', error);
				return;
			}

			dispatch('statsLoaded', data.data);
		} catch (err) {
			error = 'Network error - please try again';
			dispatch('error', error);
		} finally {
			loading = false;
		}
	}

	function handleKeyPress(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			lookupPlayer();
		}
	}
</script>

<div class="player-lookup">
	<div class="flex gap-2">
		<input
			type="text"
			bind:value={username}
			{placeholder}
			on:keypress={handleKeyPress}
			disabled={loading}
			class="block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150 flex-1"
			maxlength="12"
		/>
		<button 
			on:click={lookupPlayer} 
			disabled={loading || !username.trim()}
			class="px-4 py-2 bg-theme-accent text-white rounded-md hover:bg-theme-accent-hover disabled:opacity-50 disabled:cursor-not-allowed whitespace-nowrap font-medium transition-colors duration-150"
		>
			{loading ? 'Loading...' : buttonText}
		</button>
	</div>
	
	{#if error}
		<div class="text-red-400 text-xs mt-1 p-2 bg-red-900/20 border border-red-800/50 rounded">
			{error}
		</div>
	{/if}
</div>