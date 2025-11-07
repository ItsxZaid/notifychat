import type { CreateTopicPayload, Topic, UpdateTopicPayload } from '@/types';
import { API } from '..';

interface StateProps {
	data: Topic[];
	loading: boolean;
	error: string | null;
}

let topics = $state<StateProps>({
	data: [],
	loading: true,
	error: null
});

export const topicStore = {
	get state() {
		return topics;
	},

	load: async () => {
		try {
			topics.loading = true;
			topics.data = await API.topic.getAll();
		} catch {
			topics.error = 'Failed to load topics';
		} finally {
			topics.loading = false;
		}
	},

	createTopic: async (payload: CreateTopicPayload) => {
		try {
			topics.loading = true;
			const newTopic = await API.topic.create(payload);
			topics.data.push(newTopic);
		} catch {
			topics.error = 'Failed to delete topic';
		} finally {
			topics.loading = false;
		}
	},

	updateTopic: async (id: string, payload: UpdateTopicPayload) => {
		try {
			topics.loading = true;

			const updatedTopic = await API.topic.update(id, payload);

			const index = topics.data.findIndex((t) => t.id === updatedTopic.id);
			if (index !== -1) {
				topics.data[index] = { ...topics.data[index], ...updatedTopic };
			} else {
				topics.data.push(updatedTopic);
			}
		} catch {
			topics.error = 'Failed to update topic';
		} finally {
			topics.loading = false;
		}
	},

	deleteTopic: async (id: string) => {
		try {
			topics.loading = true;
			await API.topic.delete(id);
			topics.data = topics.data.filter((t) => t.id !== id);
		} catch {
			topics.error = 'Failed to delete topic';
		} finally {
			topics.loading = false;
		}
	}
};
