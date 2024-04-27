import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';

export const isSidebarOpen : Writable<boolean> = writable(false);
export const isTimerRunning : Writable<boolean> = writable(false);