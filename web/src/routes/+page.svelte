<script lang="ts">
	import CreateTopicSheet from '@/components/topics/CreateTopicSheet.svelte';
	import TopicCard from '@/components/topics/TopicCard.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import { Separator } from '@/components/ui/separator';
	import { topicStore } from '@/stores/topics.svelte';
	import { CirclePlus } from '@lucide/svelte';

	// States
	let showCreateSheet = $state(false);

	$effect(() => {
		topicStore.load();
	});

	let topics = $derived(topicStore.state.data);
	let loading = $derived(topicStore.state.loading);
	let error = $derived(topicStore.state.error);
</script>

<div class="min-h-screen bg-gradient-to-br from-background via-background to-muted/20">
	<div class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
		<header class="mb-8">
			<div class="flex flex-col gap-6 sm:flex-row sm:items-start sm:justify-between">
				<div class="space-y-2">
					<h1
						class="bg-gradient-to-r from-foreground to-foreground/70 bg-clip-text text-4xl font-bold tracking-tight text-transparent"
					>
						Topics
					</h1>
					<p class="text-base text-muted-foreground">
						Manage your notification topics and their channels.
					</p>
				</div>
				<Button
					type="button"
					size="lg"
					class="group inline-flex items-center gap-2 rounded-xl px-6 py-3 text-sm font-semibold shadow-md shadow-primary/25 transition-all hover:shadow-lg hover:shadow-primary/30"
					onclick={() => (showCreateSheet = true)}
				>
					<CirclePlus class="h-5 w-5 transition-transform group-hover:rotate-90" />
					Create Topic
				</Button>
			</div>
		</header>

		<CreateTopicSheet bind:open={showCreateSheet} />

		<Separator class="mb-8" />

		<main>
			{#if loading}
				<div class="flex min-h-[400px] items-center justify-center">
					<div class="text-center">
						<div
							class="mx-auto mb-4 h-12 w-12 animate-spin rounded-full border-4 border-primary/20 border-t-primary"
						></div>
						<p class="text-lg font-medium text-muted-foreground">Loading your topics...</p>
					</div>
				</div>
			{:else if error}
				<div class="rounded-2xl border border-destructive/50 bg-destructive/5 p-8 text-center">
					<div
						class="mx-auto mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-destructive/10"
					>
						<span class="text-2xl">⚠️</span>
					</div>
					<h3 class="mb-2 text-xl font-semibold text-destructive">Oops! Something went wrong</h3>
					<p class="text-muted-foreground">Error: {error}</p>
				</div>
			{:else if topics && topics.length > 0}
				<div id="topic-grid" class="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
					{#each topics as topic (topic.id)}
						<TopicCard {topic} class="border border-gray-200" />
					{/each}
				</div>
			{:else}
				<div class="flex min-h-[400px] items-center justify-center">
					<div class="max-w-md text-center">
						<div
							class="mx-auto mb-6 flex h-24 w-24 items-center justify-center rounded-full bg-muted/50"
						>
							<CirclePlus class="h-12 w-12 text-muted-foreground/50" />
						</div>
						<h3 class="mb-3 text-2xl font-bold text-foreground">No topics yet</h3>
						<p class="mb-6 text-muted-foreground">
							Get started by creating your first topic to relay notifications.
						</p>
						<Button
							class="inline-flex items-center gap-2 rounded-xl bg-primary px-6 py-3 font-semibold text-primary-foreground shadow-lg shadow-primary/25 transition-all hover:scale-105 hover:bg-primary/90"
							onclick={() => (showCreateSheet = true)}
						>
							<CirclePlus class="h-5 w-5" />
							Create Your First Topic
						</Button>
					</div>
				</div>
			{/if}
		</main>
	</div>
</div>
