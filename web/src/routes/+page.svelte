<script>
	import { campaignStore } from '@/stores/campaigns.svelte';

	$effect(() => {
		campaignStore.load();
	});
	// Ass UI, Skill issue 🫩
	let data = $derived(campaignStore.state.data);
	let loading = $derived(campaignStore.state.loading);
	let error = $derived(campaignStore.state.error);
</script>

{#if data}
	<div class="grid grid-cols-6">
		{#each data as d, i (i)}
			<div class="hover:scale-150">
				<p class="text-white">{d.id}</p>
				<p>{d.name}</p>
				<p>{d.created_at}</p>
			</div>
		{/each}
	</div>
{:else if loading}
	<div>loading...</div>
{:else if error}
	<div>{JSON.stringify(error, null, 2)}</div>
{/if}
