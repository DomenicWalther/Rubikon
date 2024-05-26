import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

interface GroupChatMessage {
	id: string;
	message: string;
	group_id: string;
	user_id: string;
}

interface Group {
	CreatedAt: string;
	DeletedAt: string | null;
	UpdatedAt: string;
	description: string;
	id: string;
	imageURL: string;
	isPrivate: boolean;
	is_member: boolean;
	name: string;
	owner_id: string,
	userCount: number;
}
interface Skin {
	activeArms: string | null;
	activeBGAnimations: string | null;
	activeEyes: string | null;
	activeHats: string | null;
	activeMouths: string | null;
}


export const isSidebarOpen: Writable<boolean> = writable(false);
export const isTimerRunning: Writable<boolean> = writable(false);
export const currentGroupChatMessages: Writable<Array<GroupChatMessage>> = writable([]);
export const currentGroups: Writable<Array<Group>> = writable([]);
export const activeSkin: Writable<Skin> = writable({} as Skin);
