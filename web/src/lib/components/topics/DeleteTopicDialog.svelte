<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	import {
		Dialog,
		DialogContent,
		DialogDescription,
		DialogFooter,
		DialogHeader,
		DialogTitle
	} from '@/components/ui/dialog';
	import type { Topic } from '$lib/types';
	import { topicStore } from '@/stores/topics.svelte';

	let {
		open = $bindable(false),
		topic
	}: {
		open?: boolean;
		topic: Topic;
	} = $props();

	let isDeleting = $state(false);

	async function handleDelete() {
		isDeleting = true;
		try {
			await topicStore.deleteTopic(topic.id);
			open = false;
		} catch (error) {
			console.error('Failed to delete topic:', error);
		} finally {
			isDeleting = false;
		}
	}

	function handleCancel() {
		if (!isDeleting) {
			open = false;
		}
	}
</script>

<Dialog bind:open>
	<DialogContent class="sm:max-w-md">
		<DialogHeader>
			<DialogTitle class="text-lg font-semibold">Delete Topic</DialogTitle>
			<DialogDescription class="text-sm text-muted-foreground">
				Are you sure you want to delete "{topic.name}"? This action cannot be undone.
			</DialogDescription>
		</DialogHeader>

		<DialogFooter class="flex-col-reverse gap-2 sm:flex-row sm:justify-end sm:gap-2">
			<Button
				variant="outline"
				onclick={handleCancel}
				disabled={isDeleting}
				class="w-full sm:w-auto"
			>
				Cancel
			</Button>
			<Button
				variant="destructive"
				onclick={handleDelete}
				disabled={isDeleting}
				class="w-full sm:w-auto"
			>
				{isDeleting ? 'Deleting...' : 'Delete Topic'}
			</Button>
		</DialogFooter>
	</DialogContent>
</Dialog>
