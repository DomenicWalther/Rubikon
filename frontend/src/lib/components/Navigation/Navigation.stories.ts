import "../../../app.css"

import type { Meta, StoryObj } from '@storybook/svelte'

import Navigation from './Navigation.svelte'

const meta: Meta<typeof Navigation> = {
    component: Navigation,
};

export default meta;
type Story = StoryObj<typeof Navigation>;

export const Primary: Story = {
    render: () => ({
        Component: Navigation,
    }),

}