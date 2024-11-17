import type { LayoutServerLoad } from "./$types";
import { makeFlasher } from "$lib/flasher";

export const load: LayoutServerLoad = ({ cookies }) => {
	const flasher = makeFlasher(cookies);
	const flash = flasher.get();

	return {
		flash,
	};
};
