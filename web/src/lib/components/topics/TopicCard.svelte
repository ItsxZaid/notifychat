<script lang="ts">
	import type { Topic } from '$lib/types';
	import { MoreHorizontal, Pencil, Trash2 } from '@lucide/svelte';
	import Button from '@/components/ui/button/button.svelte';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuItem,
		DropdownMenuSeparator,
		DropdownMenuTrigger
	} from '@/components/ui/dropdown-menu';
	import DeleteTopicDialog from './DeleteTopicDialog.svelte';
	import EditTopicDialog from './EditTopicDialog.svelte';

	let { topic, class: className = '' }: { topic: Topic; class?: string } = $props();

	let showDeleteDialog = $state(false);
	let showEditDialog = $state(false);

	function formatDate(date: Date) {
		return new Date(date).toLocaleDateString('en-US', {
			month: 'short',
			day: 'numeric',
			year: 'numeric'
		});
	}

	function handleDeleteClick() {
		showDeleteDialog = true;
	}

	function handleEditClick() {
		showEditDialog = true;
	}
</script>

<div
	class={['topic-card flex flex-col rounded-lg bg-card text-card-foreground shadow-sm', className]}
>
	<div class="flex items-start justify-between gap-4 p-6">
		<a href={`/topics/${topic.id}`} class="group grow">
			<h3 class="text-lg font-semibold tracking-tight group-hover:text-primary">
				{topic.name}
			</h3>
		</a>
		<DropdownMenu>
			<DropdownMenuTrigger>
				<div role="button" aria-label="Manage topic">
					<Button
						variant="ghost"
						size="icon-sm"
						class="-mt-1 -mr-2 text-muted-foreground"
						tabindex={-1}
					>
						<MoreHorizontal class="size-4" />
					</Button>
				</div>
			</DropdownMenuTrigger>
			<DropdownMenuContent side="bottom" align="end">
				<DropdownMenuItem onclick={handleEditClick}>
					<Pencil class="mr-2 size-4" />
					<span>Edit</span>
				</DropdownMenuItem>
				<DropdownMenuSeparator />
				<DropdownMenuItem
					class="text-destructive focus:bg-destructive/10 focus:text-destructive"
					onclick={handleDeleteClick}
				>
					<Trash2 class="mr-2 size-4" />
					<span>Delete</span>
				</DropdownMenuItem>
			</DropdownMenuContent>
		</DropdownMenu>
	</div>
	<div class="grow px-6 pt-0 pb-6">
		{#if topic.description}
			<p class="line-clamp-2 text-sm text-muted-foreground">
				{topic.description}
			</p>
		{:else}
			<p class="text-sm text-muted-foreground/60 italic">No description provided.</p>
		{/if}
	</div>
	<div class="rounded-b-lg border-t bg-muted/50 px-6 py-3">
		<p class="text-xs text-muted-foreground">
			Created: {formatDate(topic.created_at)}
		</p>
	</div>
</div>

<DeleteTopicDialog bind:open={showDeleteDialog} {topic} />
<EditTopicDialog bind:open={showEditDialog} {topic} />
