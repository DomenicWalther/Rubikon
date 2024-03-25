import "../../../app.css"

import type { Meta, StoryObj } from '@storybook/svelte'

import PastSevenDays from './PastSevenDays.svelte'

const meta: Meta<typeof PastSevenDays> = {
    component: PastSevenDays,
};

export default meta;
type Story = StoryObj<typeof PastSevenDays>;

export const Primary: Story = {
    render: () => ({
        Component: PastSevenDays,
    }),

}