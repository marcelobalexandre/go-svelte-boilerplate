import { fail, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions = {
	default: async ({ fetch, request }) => {
		const data = await request.formData();
		const username = data.get("username");
		const password = data.get("password");

		const response = await fetch("/api/login", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({ username, password }),
		});

		if (response.ok) {
			throw redirect(302, "/dashboard");
		} else if (response.status === 401) {
			const { errorMessage }: { errorMessage: string } = await response.json();
			return fail(401, { username, password, errorMessage });
		} else {
			return fail(response.status, {
				errorMessage: "An error occurred. Please try again.",
			});
		}
	},
} satisfies Actions;
