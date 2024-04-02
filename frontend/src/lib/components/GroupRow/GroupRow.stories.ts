import "../../../app.css"

import type { Meta, StoryObj } from '@storybook/svelte'

import GroupRow from './GroupRow.svelte'

const meta: Meta<typeof GroupRow> = {
    component: GroupRow,
};

export default meta;
type Story = StoryObj<typeof GroupRow>;

export const Primary: Story = {
    args: {
        group: {
  "id": 1,
  "CreatedAt": "2024-04-01T22:40:37.746917Z",
  "UpdatedAt": "2024-04-02T20:18:26.375402Z",
  "DeletedAt": null,
  "name": "20h in der Woche",
  "description": "Los gehts!",
  "owner_id": "user_2dy76EPkY5s86dm4FShjD86fnGu",
  "imageURL": "https://via.placeholder.com/150",
  "userCount": 2,
  "Users": null,
  "isPrivate": true,
  "is_member": true
}
    },
    render: (args) => ({
        Component: GroupRow,
        props: args
    }),

}