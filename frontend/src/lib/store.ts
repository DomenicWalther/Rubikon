import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

interface GroupChatMessage {
	id: string;
	message: string;
	group_id: string;
	user_id: string;
}


export const isSidebarOpen: Writable<boolean> = writable(false);
export const isTimerRunning: Writable<boolean> = writable(false);
export const currentGroupChatMessages: Writable<Array<GroupChatMessage>> = writable([]);
export const userJoinedGroups: Writable<Array<string>> = writable([]);
export const userNewGroups: Writable<Array<string>> = writable([]);
