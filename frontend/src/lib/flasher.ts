import type { Cookies } from "@sveltejs/kit";

export function makeFlasher(cookies: Cookies) {
	const get = function () {
		const flash = cookies.get("flash");
		if (!flash) {
			return { color: undefined, message: undefined };
		}

		cookies.set(
			"flash",
			JSON.stringify({ color: undefined, message: undefined }),
			{ path: "/" },
		);

		return JSON.parse(flash) as { color: "green" | "red"; message: string };
	};

	const success = function (message: string) {
		cookies.set("flash", JSON.stringify({ color: "green", message }), {
			path: "/",
		});
	};

	const alert = function (message: string) {
		cookies.set("flash", JSON.stringify({ color: "red", message }), {
			path: "/",
		});
	};

	return {
		get,
		success,
		alert,
	};
}
