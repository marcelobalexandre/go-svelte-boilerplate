import type { Actions } from "./$types";
import { fail, redirect } from "@sveltejs/kit";
import { makeFlasher } from "$lib/flasher";

export const actions = {
	default: async ({ cookies, fetch, request }) => {
		const data = await request.formData();
		const username = data.get("username");
		const password = data.get("password");

		const response = await fetch("/api/signup", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({ username, password }),
		});

		if (response.ok) {
			const flasher = makeFlasher(cookies);
			flasher.success("Account created succesfully!");
			throw redirect(302, "/");
		} else if (response.status === 422) {
			const error: { message: string; details: Record<string, string> } =
				await response.json();
			return fail(422, { username, password, error });
		} else {
			return fail(response.status, {
				error: { message: "An error occurred. Please try again." },
			});
		}
	},
} satisfies Actions;
