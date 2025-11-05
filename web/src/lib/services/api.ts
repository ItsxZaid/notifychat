import type { Campaign } from '@/types';
import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL;

const apiClient = axios.create({
	baseURL: API_URL,
	timeout: 15000,
	headers: {
		'Content-Type': 'application/json'
	}
});

export const campaignApi = {
	getAll: async (): Promise<Campaign[]> => {
		const response = await apiClient.get('/campaigns');
		return response.data.data;
	}
};
