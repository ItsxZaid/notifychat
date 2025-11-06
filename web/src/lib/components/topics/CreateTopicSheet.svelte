<script lang="ts">
	import {
		Sheet,
		SheetContent,
		SheetDescription,
		SheetHeader,
		SheetTitle,
		SheetFooter,
		SheetClose
	} from '@/components/ui/sheet';
	import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert';
	import { Button } from '@/components/ui/button';
	import { Input } from '@/components/ui/input';
	import { Label } from '@/components/ui/label';
	import { Textarea } from '@/components/ui/textarea';
	import { AlertCircle, Loader2 } from '@lucide/svelte';
	import type { CreateTopicPayload } from '@/types';
	import { topicStore } from '@/stores/topics.svelte';

	let { open = $bindable(false) } = $props();

	let data = $state<CreateTopicPayload>({
		name: '',
		description: ''
	});
	let loading = $state(false);
	let error = $state<string | null>(null);

	async function handleSubmit() {
		loading = true;
		error = null;

		if (!data.name) {
			error = 'Topic name is required.';
			loading = false;
			return;
		}

		try {
			const payload: CreateTopicPayload = {
				name: data.name,
				description: data.description || null
			};

			await topicStore.createTopic(payload);

			await topicStore.load();
			open = false;
		} catch (err: any) {
			console.error('Failed to create topic:', err);
			error = err.response?.data?.error?.message || 'An unknown error occurred.';
		} finally {
			loading = false;
		}
	}

	$effect(() => {
		if (!open) {
			data.name = '';
			data.description = '';
			error = null;
			loading = false;
		}
	});
</script>

<Sheet bind:open>
	<SheetContent side="right" class="grid grid-rows-[auto_1fr_auto] gap-8 p-6 sm:max-w-lg">
		<SheetHeader class="pr-8">
			<SheetTitle>Create New Topic</SheetTitle>
			<SheetDescription>Topics group your notification channels.</SheetDescription>
		</SheetHeader>

		<form id="create-topic-form" class="space-y-6 overflow-y-auto pr-8" onsubmit={handleSubmit}>
			{#if error}
				<Alert variant="destructive">
					<AlertCircle class="size-4" />
					<AlertTitle>Error</AlertTitle>
					<AlertDescription>{error}</AlertDescription>
				</Alert>
			{/if}

			<div class="space-y-2">
				<Label for="name">Topic Name</Label>
				<Input
					id="name"
					bind:value={data.name}
					placeholder="e.g. 'New User Signup'"
					disabled={loading}
					required
				/>
			</div>
			<div class="space-y-2">
				<Label for="description">
					Description <span class="text-xs text-muted-foreground"> (Optional)</span>
				</Label>
				<Textarea
					id="description"
					bind:value={data.description}
					placeholder="What is this topic for?"
					class="min-h-[100px]"
					disabled={loading}
				/>
			</div>
		</form>

		<SheetFooter class="pr-8">
			<SheetClose>
				<Button variant="ghost" disabled={loading}>Cancel</Button>
			</SheetClose>
			<Button type="submit" form="create-topic-form" disabled={loading}>
				{#if loading}
					<Loader2 class="mr-2 size-4 animate-spin" />
				{/if}
				Create Topic
			</Button>
		</SheetFooter>
	</SheetContent>
</Sheet>
