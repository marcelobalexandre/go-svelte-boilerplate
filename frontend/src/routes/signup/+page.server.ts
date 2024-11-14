import type { Actions } from "./$types";
import { fail, redirect } from "@sveltejs/kit";

export const actions = {
	default: async ({ fetch, request }) => {
		const data = await request.formData();
		const username = data.get("username");
		const password = data.get("password");

		const response = await fetch("/api/signup", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({ username, password }),
		});

		if (response.ok) {
			throw redirect(302, "/");
		} else if (response.status === 422) {
			const { errors }: { errors: Record<string, string> } =
				await response.json();
			return fail(422, { username, password, errors });
		} else {
			return fail(response.status, {
				error: "An error occurred. Please try again.",
			});
		}
	},
} satisfies Actions;
