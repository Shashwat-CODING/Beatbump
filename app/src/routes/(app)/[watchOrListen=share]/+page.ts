import { redirect } from "@sveltejs/kit";
import { SERVER_DOMAIN } from "../../../env";
export const load = async ({ url, fetch }) => {
	const id =
		url.searchParams.get("id") ??
		url.searchParams.get("v") ??
		url.searchParams.get("videoId");
	const playlist = url.searchParams.get("list") || undefined;

	if (!id) {
		throw redirect(301, "/trending");
	}

	const [data, list] = await Promise.all([
		fetch(
			`${SERVER_DOMAIN}/api/v1/player.json?videoId=${id ? id : ""}${
				playlist ? `&playlistId=${playlist}` : ""
			}`,
		).then((res) => res.json()),
		fetch(
			`${SERVER_DOMAIN}/api/v1/next.json?videoId=${id ? id : ""}${
				playlist ? `&playlistId=${playlist}` : ""
			}`,
		).then((res) => res.json()),
	]);

	const {
		videoDetails: {
			title = "",
			videoId = "",
			thumbnail: { thumbnails = [] } = {},
		} = {},
	} = data;

	return {
		title,
		thumbnails,
		videoId,
		playlist,
		related: list,
		data,
	};
};
