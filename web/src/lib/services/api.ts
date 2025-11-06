import type { CreateTopicPayload, Topic } from '@/types';
import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL;

const apiClient = axios.create({
	baseURL: API_URL,
	timeout: 15000,
	headers: {
		'Content-Type': 'application/json'
	}
});

export const topicApi = {
	getAll: async (): Promise<Topic[]> => {
		const response = await apiClient.get('/topics');
		return response.data.data;
	},

	create: async (payload: CreateTopicPayload): Promise<Topic> => {
		const response = await apiClient.post('/topics', {
			name: payload.name,
			description: payload.description
		});

		return response.data.data;
	},

	delete: async (id: string): Promise<void> => {
		await apiClient.delete(`/topics/${id}`);
		return;
	}
};
