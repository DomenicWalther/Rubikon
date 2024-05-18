import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

export const isSidebarOpen: Writable<boolean> = writable(false);
export const isTimerRunning: Writable<boolean> = writable(false);
export const currentGroupChatMessages: Writable<Array<string>> = writable([]);
export const userJoinedGroups: Writable<Array<string>> = writable([]);
export const userNewGroups: Writable<Array<string>> = writable([]);
