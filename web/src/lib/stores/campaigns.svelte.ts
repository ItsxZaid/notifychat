import type { Campaign } from '@/types';
import { API } from '..';

interface StateProps {
	data: Campaign[];
	loading: boolean;
	error: string | null;
}

let campaigns = $state<StateProps>({
	data: [],
	loading: true,
	error: null
});

export const campaignStore = {
	get state() {
		return campaigns;
	},

	load: async () => {
		try {
			campaigns.loading = true;
			campaigns.data = await API.campaign.getAll();
		} catch {
			campaigns.error = 'Failed to load campaigns';
		} finally {
			campaigns.loading = false;
		}
	}
};
