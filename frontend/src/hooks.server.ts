import type { Handle } from "@sveltejs/kit";
import { redirect } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
	const token = event.cookies.get("token");

	const isProtectedRoute = event.url.pathname === "/";
	if (isProtectedRoute && isTokenInvalid(token)) {
		throw redirect(302, "/login");
	}

	return await resolve(event);
};

function isTokenInvalid(token?: string): boolean {
	return !token || isTokenExpired(token);
}

function isTokenExpired(token: string): boolean {
	const payloadBase64 = token.split(".")[1];
	const payloadJson = atob(payloadBase64);
	const payload = JSON.parse(payloadJson);

	const currentTime = Math.floor(Date.now() / 1000);

	return payload.exp < currentTime;
}
