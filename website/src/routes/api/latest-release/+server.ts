import type { RequestEvent, RequestHandler } from "@sveltejs/kit";
import { GITHUB_TOKEN } from "$env/static/private";

export const GET: RequestHandler = async (_event: RequestEvent) => {
	const tag = await fetch("https://api.github.com/repos/mbaraa/dotsync/releases", {
		method: "GET",
		headers: {
			Authorization: `Bearer ${GITHUB_TOKEN}`,
			"X-GitHub-Api-Version": "2022-11-28",
			Accept: "application/vnd.github+json"
		}
	})
		.then((resp) => resp.json())
		.then((resp) => {
			if (resp && resp.length > 0 && resp[0]["tag_name"]) {
				return resp[0]["tag_name"];
			}
			return "latest";
		})
		.catch(() => {
			return "latest";
		});

	return new Response(JSON.stringify({ tag }), { status: 200 });
};
