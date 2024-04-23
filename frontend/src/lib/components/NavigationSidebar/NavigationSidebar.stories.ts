import "../../../app.css"

import type { Meta, StoryObj } from '@storybook/svelte'

import NavigationSidebar from './NavigationSidebar.svelte'

const meta: Meta<typeof NavigationSidebar> = {
    component: NavigationSidebar,
};

export default meta;
type Story = StoryObj<typeof NavigationSidebar>;

export const Primary: Story = {
    render: () => ({
        Component: NavigationSidebar,
    }),

}

