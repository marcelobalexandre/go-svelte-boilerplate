import forms from "@tailwindcss/forms";
import flowbite from "flowbite/plugin";
import type { Config } from "tailwindcss";

export default {
	content: [
		"./src/**/*.{html,js,svelte,ts}",
		"./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}",
	],

	plugins: [forms, flowbite],

	theme: {
		extend: {},
	},
} satisfies Config;
