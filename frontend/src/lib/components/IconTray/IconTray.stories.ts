import "../../../app.css"

import type { Meta, StoryObj } from '@storybook/svelte'

import IconTray from './IconTray.svelte'

const meta: Meta<typeof IconTray> = {
    component: IconTray,
};

export default meta;
type Story = StoryObj<typeof IconTray>;

export const Primary: Story = {
    render: () => ({
        Component: IconTray,
    }),

}