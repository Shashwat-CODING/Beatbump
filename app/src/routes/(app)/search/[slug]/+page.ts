import type { NextContinuationData } from "$lib/types";
import type { SearchFilter } from "$lib/types/api/search";
import type { MusicShelf } from "$lib/types/musicShelf";
import type { PageLoad } from "./$types";
import { SERVER_DOMAIN } from "../../../../env";

export interface SearchResponse {
	results?: MusicShelf[];
	continuation?: NextContinuationData;
	filter: SearchFilter;
	type?: "next" | undefined;
}
export const load: PageLoad = async ({
	url,
	params,
	fetch,
}): Promise<SearchResponse> => {
	const slug = params.slug;
	const filter = url.searchParams.get("filter") || "";
	const restricted = url.searchParams.get("restricted") || "";

	const apiUrl = `${SERVER_DOMAIN}/api/v1/search.json?q=${slug}${
		filter !== "" ? `&filter=${encodeURIComponent(filter)}` : ""
	}${restricted ? `&restricted=${restricted}` : ""}`;
	const response = await fetch(apiUrl);
	const data = (await response.json()) as SearchResponse;
	Object.assign(data, { filter });
	// if (response.ok) {
	return data;
	// }
};
