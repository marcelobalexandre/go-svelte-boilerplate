import { fail, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions = {
	default: async ({ cookies, fetch, request }) => {
		const data = await request.formData();
		const username = data.get("username");
		const password = data.get("password");

		const response = await fetch("/api/login", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({ username, password }),
		});

		if (response.ok) {
			const { token }: { token: string } = await response.json();
			cookies.set("token", token, { path: "/" });

			throw redirect(302, "/");
		} else if (response.status === 401) {
			const error: { message: string } = await response.json();
			return fail(401, { username, password, error });
		} else {
			return fail(response.status, {
				error: { message: "An error occurred. Please try again." },
			});
		}
	},
} satisfies Actions;
