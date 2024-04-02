export interface GroupRowProps {

	id: number;
	CreatedAt: string; // ISO 8601 Date string
	UpdatedAt: string; // ISO 8601 Date string
	DeletedAt: string | null; // ISO 8601 Date string or null
	name: string;
	description: string;
	owner_id: string;
	imageURL: string;
	userCount: number;
	Users: null;
	isPrivate: boolean;
	is_member: boolean;
}
