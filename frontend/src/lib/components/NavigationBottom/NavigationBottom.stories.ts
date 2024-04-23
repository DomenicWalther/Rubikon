import "../../../app.css"

import type { Meta, StoryObj } from '@storybook/svelte'

import NavigationBottom from './NavigationBottom.svelte'

const meta: Meta<typeof NavigationBottom> = {
    component: NavigationBottom,
};

export default meta;
type Story = StoryObj<typeof NavigationBottom>;

export const Primary: Story = {
    render: () => ({
        Component: NavigationBottom,
    }),

}