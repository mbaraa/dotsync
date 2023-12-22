/** @type {import('tailwindcss').Config}*/
const config = {
	content: ["./src/**/*.{html,js,svelte,ts}"],

	theme: {
		extend: {
			fontFamily: {
				Terminus: ["Terminus", "sans"]
			}
		}
	},

	plugins: []
};

module.exports = config;
