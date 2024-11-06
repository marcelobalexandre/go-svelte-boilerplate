import type { Actions } from "./$types";

export const actions = {
	default: async ({ request }) => {
		const data = await request.formData();
		const username = data.get("username");
		const password = data.get("password");

		// TODO: Sign Up.

		return { username, password };
	},
} satisfies Actions;
