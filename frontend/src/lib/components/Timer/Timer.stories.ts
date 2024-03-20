
import "../../../app.css";
import type { Meta, StoryObj } from '@storybook/svelte';

import Timer from './Timer.svelte';

const meta: Meta<typeof Timer> = {
	component: Timer,
};

export default meta;
type Story = StoryObj<typeof Timer>;

export const Primary: Story = {
	render: () => ({
		Component: Timer,
	}),
};
