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
	import { Input } from '@/components/ui/input';
	import { Label } from '@/components/ui/label';
	import { Textarea } from '@/components/ui/textarea';
	import type { Topic } from '$lib/types';
	import { topicStore } from '@/stores/topics.svelte';

	let {
		open = $bindable(false),
		topic
	}: {
		open?: boolean;
		topic: Topic;
	} = $props();

	let name = $state(topic.name);
	let description = $state(topic.description || '');
	let isSaving = $state(false);
	let nameError = $state('');

	$effect(() => {
		if (open) {
			name = topic.name;
			description = topic.description || '';
			nameError = '';
		}
	});

	function validateName(): boolean {
		if (!name.trim()) {
			nameError = 'Topic name is required';
			return false;
		}
		if (name.trim().length < 2) {
			nameError = 'Topic name must be at least 2 characters';
			return false;
		}
		nameError = '';
		return true;
	}

	async function handleSave() {
		if (!validateName()) return;

		isSaving = true;
		try {
			await topicStore.updateTopic(topic.id, {
				name: name,
				description: description
			});
			open = false;
		} catch (error) {
			console.error('Failed to update topic:', error);
			nameError = 'Failed to save changes. Please try again.';
		} finally {
			isSaving = false;
		}
	}

	function handleCancel() {
		if (!isSaving) {
			open = false;
		}
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter' && (e.metaKey || e.ctrlKey)) {
			handleSave();
		}
	}
</script>

<Dialog bind:open>
	<DialogContent class="sm:max-w-lg">
		<DialogHeader>
			<DialogTitle class="text-lg font-semibold">Edit Topic</DialogTitle>
			<DialogDescription class="text-sm text-muted-foreground">
				Make changes to your topic. Click save when you're done.
			</DialogDescription>
		</DialogHeader>

		<div class="space-y-4 py-4">
			<div class="space-y-2">
				<Label for="topic-name" class="text-sm font-medium">
					Name <span class="text-destructive">*</span>
				</Label>
				<Input
					id="topic-name"
					bind:value={name}
					placeholder="Enter topic name"
					disabled={isSaving}
					class={nameError ? 'border-destructive focus-visible:ring-destructive' : ''}
					onkeydown={handleKeydown}
					oninput={() => nameError && validateName()}
				/>
				{#if nameError}
					<p class="text-xs text-destructive">{nameError}</p>
				{/if}
			</div>

			<div class="space-y-2">
				<Label for="topic-description" class="text-sm font-medium">Description</Label>
				<Textarea
					id="topic-description"
					bind:value={description}
					placeholder="Enter a brief description (optional)"
					disabled={isSaving}
					rows={3}
					class="resize-none"
				/>
			</div>
		</div>

		<DialogFooter class="flex-col-reverse gap-2 sm:flex-row sm:justify-end sm:gap-2">
			<Button variant="outline" onclick={handleCancel} disabled={isSaving} class="w-full sm:w-auto">
				Cancel
			</Button>
			<Button onclick={handleSave} disabled={isSaving || !name.trim()} class="w-full sm:w-auto">
				{isSaving ? 'Saving...' : 'Save Changes'}
			</Button>
		</DialogFooter>
	</DialogContent>
</Dialog>
