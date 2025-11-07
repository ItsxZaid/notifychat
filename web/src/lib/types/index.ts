// Topic Types
export interface Topic {
	id: string;
	name: string;
	description: string | null;
	created_at: Date;
}

export interface CreateTopicPayload {
	name: string;
	description?: string | null;
}

export interface UpdateTopicPayload {
	name?: string | null;
	description?: string | null;
}
